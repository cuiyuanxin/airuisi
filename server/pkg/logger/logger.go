package logger

import (
	"path/filepath"

	"github.com/cuiyuanxin/airuisi/pkg/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 组合日志路径
func getFilePath(dir, level string) string {
	return filepath.Join(dir, level) + ".log"
}

// 获取日志写入路径
func getLogWriter(filepath string, maxsize, maxage, maxbackups int, localtime, compress bool) (zapcore.WriteSyncer, zapcore.WriteSyncer) {
	// 日志路径
	filePath := "./runtime/logs"

	if filepath != "" {
		filePath = filepath
	}
	// 每个日志文件长度的最大大小
	maxSize := 100
	if maxsize > 0 {
		maxSize = maxsize
	}
	// 日志保留的最大天数
	maxAge := 7
	if maxage > 0 {
		maxAge = maxage
	}
	// 只保留最近多少个日志文件，用于控制程序总日志的大小
	maxBackups := 30
	if maxbackups > 0 {
		maxBackups = maxbackups
	}
	// 是否使用本地时间
	localTime := localtime

	// info日志文件
	infoLog := &lumberjack.Logger{
		Filename:   getFilePath(filePath, zap.InfoLevel.String()), // 日志文件路径
		MaxSize:    maxSize,                                       // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups,                                    // 日志文件最多保存多少个备份
		MaxAge:     maxAge,                                        // 文件最多保存多少天
		LocalTime:  localTime,                                     // 是否使用本地时间
		Compress:   compress,                                      // 是否压缩
	}
	// error日志文件
	errLog := &lumberjack.Logger{
		Filename:   getFilePath(filePath, zap.ErrorLevel.String()), // 日志文件路径
		MaxSize:    maxSize,                                        // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups,                                     // 日志文件最多保存多少个备份
		MaxAge:     maxAge,                                         // 文件最多保存多少天
		LocalTime:  localTime,                                      // 是否使用本地时间
		Compress:   compress,                                       // 是否压缩
	}

	return zapcore.AddSync(infoLog), zapcore.AddSync(errLog)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 全路径编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func NewLogger(loggerSetting *setting.Logger) (*zap.Logger, error) {
	infoWriteSyncer, errWriteSyncer := getLogWriter(
		loggerSetting.FilePath,
		loggerSetting.MaxSize,
		loggerSetting.MaxAge,
		loggerSetting.MaxBackups,
		loggerSetting.LocalTime,
		loggerSetting.Compress,
	)
	encoder := getEncoder()

	// 实现两个判断日志等级的interface
	// 优先error
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.DebugLevel && lvl < zap.ErrorLevel
	})

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, errWriteSyncer, highPriority),
		zapcore.NewCore(encoder, infoWriteSyncer, lowPriority),
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	logger := zap.New(core, caller, development)
	zap.ReplaceGlobals(logger)

	return logger, nil
}
