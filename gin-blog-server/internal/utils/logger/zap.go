package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	g "gin-blog/internal/global"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 初始化zap日志
func InitLogger(conf *g.Config) *zap.Logger {
	// 创建日志目录
	logDir := conf.Log.Directory
	if err := os.MkdirAll(logDir, 0755); err != nil {
		fmt.Printf("创建日志目录失败: %v\n", err)
		os.Exit(1)
	}

	// 设置日志级别
	var level zapcore.Level
	switch conf.Log.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// 配置日志编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.DateTime),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 配置日志输出
	var encoder zapcore.Encoder
	switch conf.Log.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 配置日志轮转
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path.Join(logDir, "app.log"),
		MaxSize:    100, // megabytes
		MaxBackups: 30,
		MaxAge:     7,    // days
		Compress:   true, // 是否压缩
	}

	// 创建核心
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberJackLogger),
		),
		level,
	)

	// 创建logger
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(0),
	)

	// 替换全局logger
	zap.ReplaceGlobals(logger)

	return logger
}
