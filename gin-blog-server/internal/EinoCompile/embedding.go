package EinoCompile

import (
	"context"
	"errors"
	"gin-blog/internal/handle"
	"gin-blog/internal/model"

	"github.com/gin-gonic/gin"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/components/embedding"
	"go.uber.org/zap"
)

func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		return nil, errors.New("上下文类型错误：需要 *gin.Context 类型")
	}

	configMap, err := model.GetConfigMap(handle.GetDB(ginCtx))
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
