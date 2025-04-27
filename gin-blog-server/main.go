package main

import (
	"flag"
	ginblog "gin-blog/internal"
	g "gin-blog/internal/global"
	"gin-blog/internal/middleware"
	"gin-blog/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("c", "./config.yml", "配置文件路径")
	flag.Parse()

	// 根据命令行参数读取配置文件, 其他变量的初始化依赖于配置文件对象
	conf := g.ReadConfig(*configPath)

	_ = ginblog.InitLogger(conf)
	utils.DB = ginblog.InitDatabase(conf)
	utils.RDB = ginblog.InitRedis(conf)

	// 初始化 gin 服务
	gin.SetMode(conf.Server.Mode)
	r := gin.New()
	r.SetTrustedProxies([]string{"*"})
	// 开发模式使用 gin 自带的日志和恢复中间件, 生产模式使用自定义的中间件
	if conf.Server.Mode == "debug" {
		r.Use(gin.Logger(), gin.Recovery()) // gin 自带的日志和恢复中间件, 挺好用的
		r.Use(middleware.ShowResponse())    // 将 logPostRequestBody 放在 gin.Logger() 之后
	} else {
		r.Use(middleware.Recovery(true), middleware.Logger())
		r.Use(middleware.ShowResponse()) // 将 logPostRequestBody 放在自定义 Logger 之后
	}
	r.Use(middleware.CORS())
	r.Use(middleware.WithGormDB(utils.DB))
	r.Use(middleware.WithRedisDB(utils.RDB))
	r.Use(middleware.WithCookieStore(conf.Session.Name, conf.Session.Salt))
	ginblog.RegisterHandlers(r)

	// 使用本地文件上传, 需要静态文件服务, 使用七牛云不需要
	if conf.Upload.OssType == "local" {
		r.Static(conf.Upload.Path, conf.Upload.StorePath)
	}

	serverAddr := conf.Server.Port
	if serverAddr[0] == ':' || strings.HasPrefix(serverAddr, "0.0.0.0:") {
		zap.L().Info("Starting server", zap.String("url", "http://localhost:"+strings.Split(serverAddr, ":")[1]))
	} else {
		zap.L().Info("Starting server", zap.String("url", "http://"+serverAddr))
	}
	r.Run(serverAddr)
}
