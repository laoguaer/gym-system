package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func ShowResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 仅处理 POST 请求
		if c.Request.Method == "POST" {
			// 读取请求体
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Printf("Error reading request body: %v", err)
				c.AbortWithStatus(500)
				return
			}

			// 将请求体重新写回，以便后续中间件和处理器可以读取
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

			// 继续处理请求
			c.Next()

			// 在请求处理完成后记录请求体
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, bodyBytes, "", "  "); err != nil {
				log.Printf("Error formatting JSON: %v", err)
			} else {
				// 确保在 Gin 的日志之后输出请求体日志
				log.Printf("Request Body: %s", prettyJSON.String())
			}
		} else {
			c.Next()
		}
	}
}