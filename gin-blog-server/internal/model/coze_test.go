package model_test

import (
	"context"
	"errors"
	"fmt"
	ginblog "gin-blog/internal"

	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"io"
	"testing"

	"github.com/coze-dev/coze-go"
	"github.com/stretchr/testify/assert"
)

func TestCozeChat(t *testing.T) {
	ctx := context.Background()
	db := ginblog.InitDatabase(g.ReadConfig("../../config.yml"))
	client, err := model.GetCozeClient(db)
	assert.Nil(t, err)
	appid := "7465547666353684507"
	// botId := "7452369603411918883"
	conversationID := "7467427656322105363"
	assert.Nil(t, err)
	fmt.Printf("conversationID:%s\n", conversationID)
	// Step one, create chats
	req := &coze.WorkflowsChatStreamReq{
		AppID:      &appid,
		// BotID:      &botId,
		WorkflowID: "7454209939809320995",
		AdditionalMessages: []*coze.Message{
			coze.BuildUserQuestionText("请问我有几个哥哥", nil),
		},
		ConversationID: &conversationID,
	}

	resp, err := client.Workflows.Chat.Stream(ctx, req)
	if err != nil {
		fmt.Printf("Error starting chats: %v\n", err)
		return
	}

	defer resp.Close()
	for {
		event, err := resp.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if event.Event == coze.ChatEventConversationMessageDelta {
			fmt.Print(event.Message.Content)
		} else if event.Event == coze.ChatEventConversationChatCompleted {
			fmt.Printf("Token usage:%d\n", event.Chat.Usage.TokenCount)
		} else {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("done, log:%s\n", resp.Response().LogID())
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestCozeMessage(t *testing.T) {
	ctx := context.Background()
	db := ginblog.InitDatabase(g.ReadConfig("../../config.yml"))
	client, err := model.GetCozeClient(db)
	assert.Nil(t, err)
	// botId := "7452369603411918883"
	// appid := "7465547666353684507"
	conversationID := "7467427656322105363"
	assert.Nil(t, err)
	fmt.Printf("conversationID:%s\n", conversationID)

	message, err := client.Conversations.Messages.List(ctx, &coze.ListConversationsMessagesReq{ConversationID: conversationID})
	if err != nil {
		fmt.Println("Error fetching message:", err)
		return
	}
	for message.Next() {
		fmt.Printf("type:%+v %v\n", message.Current().Type, message.Current().Content)
	}
	if message.Err() != nil {
		fmt.Println("Error fetching message:", message.Err())
		return
	}

	// the page result will return followed information
	fmt.Println("has_more:", message.HasMore())
	
}
