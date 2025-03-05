package ginblog

import (
	"context"
	g "gin-blog/internal/global"
	"gin-blog/internal/model"
	"gin-blog/internal/utils/logger"
	"log"

	"go.uber.org/zap"

	"github.com/glebarez/sqlite"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 根据配置文件初始化日志
func InitLogger(conf *g.Config) *zap.Logger {
	return logger.InitLogger(conf)
}

// 根据配置文件初始化数据库
func InitDatabase(conf *g.Config) *gorm.DB {
	dbtype := conf.DbType()
	dsn := conf.DbDSN()

	var db *gorm.DB
	var err error

	var level glogger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = glogger.Silent
	case "info":
		level = glogger.Info
	case "warn":
		level = glogger.Warn
	case "error":
		fallthrough
	default:
		level = glogger.Error
	}

	config := &gorm.Config{
		Logger:                                   glogger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
	}

	switch dbtype {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), config)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), config)
	default:
		log.Fatal("不支持的数据库类型: ", dbtype)
	}

	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Println("数据库连接成功", dbtype, dsn)

	if conf.Server.DbAutoMigrate {
		if err := model.MakeMigrate(db); err != nil {
			log.Fatal("数据库迁移失败", err)
		}
		log.Println("数据库自动迁移成功")
	}

	return db
}

// 根据配置文件初始化 Redis
func InitRedis(conf *g.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis 连接失败: ", err)
	}

	log.Println("Redis 连接成功", conf.Redis.Addr, conf.Redis.DB, conf.Redis.Password)
	return rdb
}
