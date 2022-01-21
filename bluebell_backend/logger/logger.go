package logger

import (
	"bluebell_backend/settings"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

// Init 初始化lg
func Init(cfg *settings.LogConfig, mode string) (err error) {
	// 定义好 core 的三大参数 —— writeSyncer， encoder， Level
	// writeSyncer 指定日志将写到哪里去
	writeSyncer := getLogWriter(
		cfg.Filename,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge,
	)
	encoder := getEncoder()    // 编码器
	var l = new(zapcore.Level) // 日志等级
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	// 根据不同情况来修改 core 的内置类型
	var core zapcore.Core
	if mode == "dev" {
		// 进入开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件名称
		MaxSize:    maxSize,   // 最大内存（超出则切割）
		MaxBackups: maxBackup, // 最大备份数量
		MaxAge:     maxAge,    // 最大备份天数
		Compress:   false,     // 是否将文件压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
