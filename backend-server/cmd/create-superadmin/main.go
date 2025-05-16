package main

import (
	"errors"
	"flag"
	ginblog "gin-blog/internal"
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"gin-blog/internal/utils"
	"os"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	username := flag.String("username", "", "超级管理员账户")
	password := flag.String("password", "", "超级管理员密码")
	configPath := flag.String("c", "../../config.yml", "配置文件路径")
	flag.Parse()

	// 根据命令行参数读取配置文件, 其他变量的初始化依赖于配置文件对象
	conf := g.ReadConfig(*configPath)

	//! 处理 sqlite3 数据库路径
	conf.SQLite.Dsn = "../" + conf.SQLite.Dsn
	conf.Server.DbLogMode = "silent"

	db := ginblog.InitDatabase(conf)

	if *username == "" || *password == "" {
		zap.L().Fatal("请指定超级管理员账户和密码")
	}

	createSuperAdmin(db, *username, *password)
}

// 创建超级管理员
func createSuperAdmin(db *gorm.DB, username, password string) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var userAuth model.UserAuth
		err := db.Where("username = ?", username).First(&userAuth).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		if userAuth.ID != 0 {
			return errors.New(userAuth.Username + " 账户已存在")
		}

		zap.L().Info("开始创建超级管理员")

		// 默认生成一个 super admin 用户
		hashPassword, err := utils.BcryptHash(password)
		if err != nil {
			return errors.New("密码生成失败: " + err.Error())
		}

		userAuth = model.UserAuth{
			Username: username,
			Password: hashPassword,
			IsSuper:  true,
			UserInfo: &model.UserInfo{
				Nickname: username,
				Avatar:   "https://cdn.hahacode.cn/config/superadmin_avatar.jpg",
				Intro:    "这个人很懒，什么都没有留下",
				Website:  "https://www.hahacode.cn",
			},
		}
		if err := db.Create(&userAuth).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		zap.L().Error("创建超级管理员失败", zap.Error(err))
		os.Exit(0)
	}

	zap.L().Info("创建超级管理员成功", zap.String("username", username), zap.String("password", password))
}
