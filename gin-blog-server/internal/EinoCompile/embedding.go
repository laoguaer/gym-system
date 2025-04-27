package EinoCompile

import (
	"context"
	"gin-blog/internal/model"
	"gin-blog/internal/utils"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/components/embedding"
	"go.uber.org/zap"
)

func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {

	configMap, err := model.GetConfigMap(utils.DB)
	if err != nil {
		zap.L().Error("获取配置失败", zap.Error(err))
		return
	}
	ARK_EMBEDDING_MODEL := configMap["ARK_EMBEDDING_MODEL"]
	ARK_API_KEY := configMap["ARK_API_KEY"]
	config := &ark.EmbeddingConfig{
		Model:  ARK_EMBEDDING_MODEL,
		APIKey: ARK_API_KEY,
	}
	eb, err = ark.NewEmbedder(ctx, config)
	if err != nil {
		return nil, err
	}
	return eb, nil
}
