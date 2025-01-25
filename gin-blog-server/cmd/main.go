package main

import (
	"bytes"
	"encoding/json"
	"flag"
	ginblog "gin-blog/internal"
	g "gin-blog/internal/global"
	"gin-blog/internal/middleware"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func logPostRequestBody() gin.HandlerFunc {
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

func main() {
	configPath := flag.String("c", "../config.yml", "配置文件路径")
	flag.Parse()

	// 根据命令行参数读取配置文件, 其他变量的初始化依赖于配置文件对象
	conf := g.ReadConfig(*configPath)

	_ = ginblog.InitLogger(conf)
	db := ginblog.InitDatabase(conf)
	rdb := ginblog.InitRedis(conf)

	// 初始化 gin 服务
	gin.SetMode(conf.Server.Mode)
	r := gin.New()
	r.SetTrustedProxies([]string{"*"})
	// 开发模式使用 gin 自带的日志和恢复中间件, 生产模式使用自定义的中间件
	if conf.Server.Mode == "debug" {
		r.Use(logPostRequestBody())         // 将 logPostRequestBody 放在 gin.Logger() 之后
		r.Use(gin.Logger(), gin.Recovery()) // gin 自带的日志和恢复中间件, 挺好用的
	} else {
		r.Use(logPostRequestBody()) // 将 logPostRequestBody 放在自定义 Logger 之后
		r.Use(middleware.Recovery(true), middleware.Logger())
	}
	r.Use(middleware.CORS())
	r.Use(middleware.WithGormDB(db))
	r.Use(middleware.WithRedisDB(rdb))
	r.Use(middleware.WithCookieStore(conf.Session.Name, conf.Session.Salt))
	ginblog.RegisterHandlers(r)

	// 使用本地文件上传, 需要静态文件服务, 使用七牛云不需要
	if conf.Upload.OssType == "local" {
		r.Static(conf.Upload.Path, conf.Upload.StorePath)
	}

	serverAddr := conf.Server.Port
	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr, "0.0.0.0:") {
		log.Printf("Serving HTTP on (http://localhost:%s/) ... \n", strings.Split(serverAddr, ":")[1])
	} else {
		log.Printf("Serving HTTP on (http://%s/) ... \n", serverAddr)
	}
	r.Run(serverAddr)
}
