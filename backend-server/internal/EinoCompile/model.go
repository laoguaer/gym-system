package EinoCompile

import (
	"context"

	model2 "gin-blog/internal/model"
	"gin-blog/internal/utils"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino/components/model"
	"go.uber.org/zap"
)

func newChatModel(ctx context.Context) (cm model.ChatModel, err error) {

	configMap, err := model2.GetConfigMap(utils.DB)
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
