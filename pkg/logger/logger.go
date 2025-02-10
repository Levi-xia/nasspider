package logger

import (
	"fmt"
	"nasspider/config"
	"os"
	"path/filepath"
	"time"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

type LogConfig struct {
	DebugFileName string `json:"debugFileName"`
	InfoFileName  string `json:"infoFileName"`
	WarnFileName  string `json:"warnFileName"`
	ErrorFileName string `json:"errorFileName"`
	MaxSize       int    `json:"maxsize"`
	MaxAge        int    `json:"max_age"`
	MaxBackups    int    `json:"max_backups"`
}

func InitLog() (err error) {
	loggerCnf := config.Conf.Logger
	workDir, _ := os.Getwd()
	// 初始化日志
	subfix := time.Now().Format("2006010215")
	cfg := &LogConfig{
		DebugFileName: filepath.Join(workDir, fmt.Sprintf("%s.%s", loggerCnf.DebugFileName, subfix)),
		InfoFileName:  filepath.Join(workDir, fmt.Sprintf("%s.%s", loggerCnf.InfoFileName, subfix)),
		WarnFileName:  filepath.Join(workDir, fmt.Sprintf("%s.%s", loggerCnf.WarnFileName, subfix)),
		ErrorFileName: filepath.Join(workDir, fmt.Sprintf("%s.%s", loggerCnf.ErrorFileName, subfix)),
		MaxSize:       loggerCnf.MaxSize,
		MaxAge:        loggerCnf.MaxAge,
		MaxBackups:    loggerCnf.MaxBackups,
	}
	writeSyncerDebug := getLogWriter(cfg.DebugFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerInfo := getLogWriter(cfg.InfoFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerWarn := getLogWriter(cfg.WarnFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncerError := getLogWriter(cfg.ErrorFileName, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)

    // 定义日志级别
    debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.DebugLevel
    })
    infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.InfoLevel
    })
    warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.WarnLevel
    })
    errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
        return lvl == zapcore.ErrorLevel
    })

    encoder := getEncoder()
    // 文件输出
    debugCore := zapcore.NewCore(encoder, writeSyncerDebug, debugLevel)
    infoCore := zapcore.NewCore(encoder, writeSyncerInfo, infoLevel)
    warnCore := zapcore.NewCore(encoder, writeSyncerWarn, warnLevel)
    errorCore := zapcore.NewCore(encoder, writeSyncerError, errorLevel)
    
	// 标准输出
	var core zapcore.Core
	if config.Conf.Server.Debug {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		std := zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zap.DebugLevel)
		core = zapcore.NewTee(infoCore, debugCore, warnCore, errorCore, std)
	} else {
		core = zapcore.NewTee(infoCore, debugCore, warnCore, errorCore)
	}
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	Logger = logger.Sugar()
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(time.Local().Format("2006-01-02 15:04:05"))
	}
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
