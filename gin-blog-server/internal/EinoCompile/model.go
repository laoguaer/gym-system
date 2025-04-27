package EinoCompile

import (
	"context"
	"errors"
	"gin-blog/internal/handle"

	model2 "gin-blog/internal/model"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func newChatModel(ctx context.Context) (cm model.ChatModel, err error) {
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, errors.New("上下文类型错误：需要 *gin.Context 类型")
	}

	configMap, err := model2.GetConfigMap(handle.GetDB(ginCtx))
	if err != nil {
		zap.L().Error("获取配置失败", zap.Error(err))
		return
	}
	ARK_CHAT_MODEL := configMap["ARK_CHAT_MODEL"]
	ARK_API_KEY := configMap["ARK_API_KEY"]
	config := &ark.ChatModelConfig{
		Model:  ARK_CHAT_MODEL,
		APIKey: ARK_API_KEY,
	}
	cm, err = ark.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
