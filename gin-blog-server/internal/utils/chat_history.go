package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/eino/schema"
	"github.com/go-redis/redis/v9"
)

const (
	// 聊天历史记录的Redis键前缀
	ChatHistoryKey = "chat:history:"
	// 历史记录保存的最大条数
	MaxHistorySize = 50
	// 历史记录过期时间（7天）
	HistoryExpiration = 7 * 24 * time.Hour
)

// 保存用户消息到历史记录
func SaveUserMessage(ctx context.Context, rdb *redis.Client, userId string, message string) error {
	key := ChatHistoryKey + userId

	// 创建用户消息
	userMsg := schema.UserMessage(message)

	// 序列化消息
	msgBytes, err := json.Marshal(userMsg)
	if err != nil {
		return fmt.Errorf("序列化用户消息失败: %w", err)
	}

	// 添加到Redis列表
	err = rdb.RPush(ctx, key, string(msgBytes)).Err()
	if err != nil {
		return fmt.Errorf("保存用户消息到Redis失败: %w", err)
	}

	// 设置过期时间
	rdb.Expire(ctx, key, HistoryExpiration)

	// 保持列表长度不超过最大值
	current, err := rdb.LLen(ctx, key).Result()
	if err == nil && current > MaxHistorySize {
		// 移除最早的消息
		rdb.LTrim(ctx, key, current-MaxHistorySize, -1)
	}

	return nil
}

// 保存AI回复到历史记录
func SaveAIMessage(ctx context.Context, rdb *redis.Client, userId string, message *schema.Message) error {
	key := ChatHistoryKey + userId

	// 序列化消息
	msgBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("序列化AI消息失败: %w", err)
	}

	// 添加到Redis列表
	err = rdb.RPush(ctx, key, string(msgBytes)).Err()
	if err != nil {
		return fmt.Errorf("保存AI消息到Redis失败: %w", err)
	}

	// 设置过期时间
	rdb.Expire(ctx, key, HistoryExpiration)

	// 保持列表长度不超过最大值
	current, err := rdb.LLen(ctx, key).Result()
	if err == nil && current > MaxHistorySize {
		// 移除最早的消息
		rdb.LTrim(ctx, key, current-MaxHistorySize, -1)
	}

	return nil
}

// 获取用户的聊天历史记录
func GetChatHistory(ctx context.Context, rdb *redis.Client, userId string) ([]*schema.Message, error) {
	key := ChatHistoryKey + userId

	// 从Redis获取历史记录
	msgs, err := rdb.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, fmt.Errorf("从Redis获取历史记录失败: %w", err)
	}

	// 如果没有历史记录，返回空切片
	if len(msgs) == 0 {
		return []*schema.Message{}, nil
	}

	// 反序列化消息
	history := make([]*schema.Message, 0, len(msgs))
	for _, msgStr := range msgs {
		var msg schema.Message
		err := json.Unmarshal([]byte(msgStr), &msg)
		if err != nil {
			// 记录错误但继续处理其他消息
			fmt.Printf("反序列化消息失败: %v\n", err)
			continue
		}
		history = append(history, &msg)
	}

	return history, nil
}

// 清除用户的聊天历史记录
func ClearChatHistory(ctx context.Context, rdb *redis.Client, userId string) error {
	key := ChatHistoryKey + userId

	// 删除Redis中的历史记录
	_, err := rdb.Del(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("清除历史记录失败: %w", err)
	}

	return nil
}
