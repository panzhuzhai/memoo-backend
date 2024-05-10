package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

var ZapLogger *zap.SugaredLogger
var GormLogger Logger

// Logger is an alternative implementation of *gorm.Logger
type Logger interface {
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Info(ctx context.Context, msg string, data ...interface{})
	Warn(ctx context.Context, msg string, data ...interface{})
	Error(ctx context.Context, msg string, data ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

type logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(sugar *zap.SugaredLogger) Logger {
	return &logger{Zap: sugar}
}

func InitLogger(ginMode string) {
	config := zap.NewDevelopmentConfig()
	if ginMode == gin.ReleaseMode {
		config = zap.NewProductionConfig()
		encoderConfig := zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}
		config.EncoderConfig = encoderConfig
		config.OutputPaths = []string{"stdout", "/app/logs/application.log"}
		config.ErrorOutputPaths = []string{"stdout", "/app/logs/application.log"}
	}
	logger, _ := config.Build()
	sugar := logger.Sugar()
	GormLogger = NewLogger(sugar)
	ZapLogger = sugar

	_ = logger.Sync()
}
