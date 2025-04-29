package handle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-blog/internal/EinoCompile"
	"gin-blog/internal/utils"
	"io"
	"log"
	"net/http"

	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Chat struct{}
type ChatReq struct {
	ChatString string `json:"chat_string" binding:"required"`
}

type ChatResp struct {
	Code     int    `json:"code"`
	Cost     string `json:"cost"`
	Data     string `json:"data"`
	DebugUrl string `json:"debug_url"`
	Msg      string `json:"msg"`
	Token    int    `json:"token"`
}

func (*Chat) Send(c *gin.Context) {
	id := c.Query("user_id")
	message := c.Query("message")
	if id == "" || message == "" {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  "missing id or message parameter",
		})
		return
	}

	zap.L().Sugar().Infof("[Chat] Starting chat with ID: %s, Message: %s\n", id, message)

	// 设置SSE响应头
	c.Writer.Header().Set("Content-Type", sse.ContentType)
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Transfer-Encoding", "chunked")
	c.Writer.WriteHeader(http.StatusOK)

	// 获取消息流
	sr, err := RunAgent(c, id, message)
	if err != nil {
		log.Printf("[Chat] Error running agent: %v\n", err)
		// 发送错误消息
		err = sse.Encode(c.Writer, sse.Event{
			Event: "error",
			Data:  err.Error(),
		})
		if err != nil {
			log.Printf("[Chat] Error encoding SSE event: %v\n", err)
		}
		return
	}

	// 刷新缓冲区，确保头信息被发送
	c.Writer.Flush()

	// 发送初始连接成功消息
	err = sse.Encode(c.Writer, sse.Event{
		Event: "connected",
		Data:  "SSE连接已建立",
	})
	if err != nil {
		log.Printf("[Chat] Error sending connected event: %v\n", err)
		return
	}
	c.Writer.Flush()

	// 处理消息流
	// 用于跟踪token使用情况
	tokenCount := 0

outer:
	for {
		select {
		case <-c.Done():
			log.Printf("[Chat] Context done for chat ID: %s\n", id)
			return
		default:
			msg, err := sr.Recv()
			if errors.Is(err, io.EOF) {
				log.Printf("[Chat] EOF received for chat ID: %s\n", id)
				// 发送结束消息
				_ = sse.Encode(c.Writer, sse.Event{
					Event: "done",
					Data:  "聊天完成",
				})
				c.Writer.Flush()

				// 发送token使用信息
				tokenInfo := map[string]int{"token_count": tokenCount}
				tokenData, _ := json.Marshal(tokenInfo)
				_ = sse.Encode(c.Writer, sse.Event{
					Event: "token_info",
					Data:  string(tokenData),
				})
				c.Writer.Flush()
				break outer
			}
			if err != nil {
				log.Printf("[Chat] Error receiving message: %v\n", err)
				// 发送错误消息
				_ = sse.Encode(c.Writer, sse.Event{
					Event: "error",
					Data:  err.Error(),
				})
				c.Writer.Flush()
				break outer
			}

			// 估算token数量（简单实现，每个字符算作1个token）
			tokenCount += len(msg.Content)

			// 发送消息内容
			err = sse.Encode(c.Writer, sse.Event{
				Event: "message",
				Data:  msg.Content,
			})
			if err != nil {
				log.Printf("[Chat] Error publishing message: %v\n", err)
				break outer
			}
			c.Writer.Flush()
		}
	}
}

func (*Chat) GetHistory(c *gin.Context) {
	id := c.Query("user_id")
	if id == "" {
		c.JSON(consts.StatusBadRequest, map[string]string{
			"status": "error",
			"error":  "missing user_id parameter",
		})
		return
	}

	zap.L().Sugar().Infof("[Chat] Getting chat history for ID: %s\n", id)

	// 从Redis获取历史会话
	redisCli := utils.RDB
	history, err := utils.GetChatHistory(c, redisCli, id, 10)
	if err != nil {
		log.Printf("[Chat] Error getting chat history: %v\n", err)
		c.JSON(consts.StatusInternalServerError, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	// 将历史记录转换为前端可用的格式
	type MessageResponse struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}

	responses := make([]MessageResponse, 0, len(history))
	for _, msg := range history {
		role := "assistant"
		if msg.Role == schema.User {
			role = "user"
		}
		responses = append(responses, MessageResponse{
			Role:    role,
			Content: msg.Content,
		})
	}

	c.JSON(consts.StatusOK, map[string]interface{}{
		"status":  "success",
		"history": responses,
	})
}

func RunAgent(ctx context.Context, id string, msg string) (*schema.StreamReader[*schema.Message], error) {

	runner, err := EinoCompile.BuildMy(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build agent graph: %w", err)
	}

	redisCli := utils.RDB

	// 从Redis获取历史会话
	history, err := utils.GetChatHistory(ctx, redisCli, id, 0)
	if err != nil {
		log.Printf("[Chat] 获取历史会话失败: %v，将使用空历史记录", err)
		history = []*schema.Message{}
	}

	// 保存用户消息到Redis
	err = utils.SaveUserMessage(ctx, redisCli, id, msg)
	if err != nil {
		log.Printf("[Chat] 保存用户消息失败: %v", err)
	}

	userMessage := &EinoCompile.UserMessage{
		ID:      id,
		Query:   fmt.Sprintf("[userId:%s]%s", id, msg),
		History: history,
	}

	sr, err := runner.Stream(ctx, userMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to stream: %w", err)
	}

	srs := sr.Copy(2)

	go func() {
		// for save to memory
		fullMsgs := make([]*schema.Message, 0)

		defer func() {
			// close stream if you used it
			srs[1].Close()

			// 合并所有消息片段
			fullMsg, err := schema.ConcatMessages(fullMsgs)
			if err != nil {
				fmt.Println("error concatenating messages: ", err.Error())
				return
			}

			// 保存AI回复到Redis
			err = utils.SaveAIMessage(ctx, redisCli, id, fullMsg)
			if err != nil {
				fmt.Println("保存AI回复到历史记录失败: ", err.Error())
			}
		}()

	outer:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done", ctx.Err())
				return
			default:
				chunk, err := srs[1].Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break outer
					}
					// 处理其他错误
					fmt.Println("接收消息错误:", err.Error())
					continue
				}

				// 添加消息片段到完整消息列表
				fullMsgs = append(fullMsgs, chunk)
			}
		}
	}()

	return srs[0], nil
}
