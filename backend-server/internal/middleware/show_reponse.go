package middleware

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CustomResponseWriter 包装gin的ResponseWriter以捕获响应数据
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ShowResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 读取请求体
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			zap.L().Sugar().Errorf("Error reading request body: %v", err)
			c.AbortWithStatus(500)
			return
		}

		// 将请求体重新写回，以便后续中间件和处理器可以读取
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// 创建自定义ResponseWriter
		blw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// 继续处理请求
		c.Next()

		// 记录请求和响应信息
		zap.L().Info("HTTP Request/Response",
			zap.String("request_body", string(bodyBytes)),
			zap.String("response_body", blw.body.String()),
		)
	}
}
