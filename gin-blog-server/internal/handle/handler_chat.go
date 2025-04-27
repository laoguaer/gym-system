package handle

import (
	"context"
	"errors"
	"fmt"
	"gin-blog/internal/EinoCompile"
	"io"
	"log"

	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/*
curl -X POST 'https://api.coze.cn/v1/workflow/stream_run' \
-H "Authorization: Bearer pat_Ozo3WylcPI18kUqTY7wYNucGrjUtQziKsUy73dALUhRTfu0VZmzLktyabPs132hP" \
-H "Content-Type: application/json" \

	-d '{
	  "parameters": {
	    "start": "默认"
	  },
	  "workflow_id": "7454209939809320995"
	}'

{"code":0,"cost":"0","data":"{\"content_type\":1,\"data\":\"{\\\"output\\\":\\\"{\\\\\\\"回答\\\\\\\":\\\\\\\"这是默认消息\\\\\\\",\\\\\\\"问题\\\\\\\":\\\\\\\"默认\\\\\\\"}\\\"}\",\"original_result\":null,\"type_for_model\":2}","debug_url":"https://www.coze.cn/work_flow?execute_id=7464449491292700682\u0026space_id=7452369603411329059\u0026workflow_id=7454209939809320995","msg":"Success","token":269}
*/

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

type CozeReq struct {
	Parameters map[string]string `json:"parameters" binding:"required"`
	WorkflowId string            `json:"workflow_id" binding:"required"`
}
type CozeResp struct {
	Code     int    `json:"code"`
	Cost     string `json:"cost"`
	Data     string `json:"data"`
	DebugUrl string `json:"debug_url"`
	Msg      string `json:"msg"`
	Token    int    `json:"token"`
}
type CozeData struct {
	ContentType    int    `json:"content_type"`
	Data           string `json:"data"`
	OriginalResult string `json:"original_result"`
	TypeForModel   int    `json:"type_for_model"`
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

	_, err := RunAgent(c, id, message)
	if err != nil {
		log.Printf("[Chat] Error running agent: %v\n", err)
		c.JSON(consts.StatusInternalServerError, map[string]string{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

outer:
	for {
		select {
		case <-c.Done():
			log.Printf("[Chat] Context done for chat ID: %s\n", id)
			return
		default:
			// msg, err := sr.Recv()
			if errors.Is(err, io.EOF) {
				log.Printf("[Chat] EOF received for chat ID: %s\n", id)
				break outer
			}
			if err != nil {
				log.Printf("[Chat] Error receiving message: %v\n", err)
				break outer
			}

			// err = s.Publish(&sse.Event{
			// 	Data: []byte(msg.Content),
			// })
			if err != nil {
				log.Printf("[Chat] Error publishing message: %v\n", err)
				break outer
			}
		}
	}

	// log.Printf("chat Request: %v\n response: %v", string(requestBody), apiResp)
	// c.JSON(http.StatusOK, apiResp)
}

func (*Chat) GetHistory(c *gin.Context) {

}

func RunAgent(ctx context.Context, id string, msg string) (*schema.StreamReader[*schema.Message], error) {

	runner, err := EinoCompile.BuildMy(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build agent graph: %w", err)
	}

	// conversation := memory.GetConversation(id, true)

	userMessage := &EinoCompile.UserMessage{
		ID:      id,
		Query:   msg,
		History: nil,
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

			// add user input to history
			// conversation.Append(schema.UserMessage(msg))

			// fullMsg, err := schema.ConcatMessages(fullMsgs)
			if err != nil {
				fmt.Println("error concatenating messages: ", err.Error())
			}
			// add agent response to history
			// conversation.Append(fullMsg)
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
				}

				fullMsgs = append(fullMsgs, chunk)
			}
		}
	}()

	return srs[0], nil
}
