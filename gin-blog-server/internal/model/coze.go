package model

import (
	"context"
	"errors"
	"fmt"
	g "gin-blog/internal/global"
	"net/http"
	"time"

	"github.com/coze-dev/coze-go"
	"gorm.io/gorm"
)

func GetCozeToken(db *gorm.DB) (string, error) {
	configMap, err := GetConfigMap(db)
	if err != nil {
		return "", err
	}
	token := configMap["coze_token"]
	return token, nil
}

func GetCozeWorkflowId(db *gorm.DB) (string, error) {
	workflowId := g.GetConfig().Coze.WorkflowId
	if workflowId == "" {
		return "", errors.New("coze_workflow_id is empty")
	}
	return workflowId, nil
}

func GetCozeUrl(db *gorm.DB) (string, error) {
	url := g.GetConfig().Coze.Url
	if url == "" {
		return "", errors.New("coze_url is empty")
	}
	return url, nil
}

func GetCozeClient(db *gorm.DB) (coze.CozeAPI, error) {
	cozeAPIToken, err := GetCozeToken(db)
	if err != nil {
		return coze.CozeAPI{}, err
	}
	baseUrl := g.GetConfig().Coze.BaseUrl
	authCli := coze.NewTokenAuth(cozeAPIToken)
	customClient := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
	cozeCli3 := coze.NewCozeAPI(authCli,
		coze.WithBaseURL(baseUrl),
		coze.WithHttpClient(customClient),
	)
	return cozeCli3, nil
}

func NewCozeConversation(db *gorm.DB) (string, error) {
	client, err := GetCozeClient(db)
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	botId := g.GetConfig().Coze.BotId
	// Create a new conversation
	resp, err := client.Conversations.Create(ctx, &coze.CreateConversationsReq{BotID: botId})
	if err != nil {
		return "", err
	}
	fmt.Println("create conversations:", resp.Conversation)
	fmt.Println(resp.LogID())

	conversationID := resp.Conversation.ID
	return conversationID, nil
}

func GetHistoryMessage(db *gorm.DB, conversationId string) ([]*coze.Message, error) {
	client, err := GetCozeClient(db)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	// you can use iterator to automatically retrieve next page
	message, err := client.Conversations.Messages.List(ctx, &coze.ListConversationsMessagesReq{Limit: 2, ConversationID: conversationId})
	if err != nil {
		fmt.Println("Error fetching message:", err)
		return nil, err
	}
	for message.Next() {
		fmt.Println(message.Current())
	}
	if message.Err() != nil {
		fmt.Println("Error fetching message:", message.Err())
		return nil, message.Err()
	}

	// the page result will return followed information
	fmt.Println("has_more:", message.HasMore())
	return nil, nil
}
