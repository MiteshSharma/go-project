package logger

import (
	"github.com/MiteshSharma/project/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	Zap *zap.Logger
}

func NewLogger(config *model.Config) *ZapLogger {
	zapConfig := generateConfig(config)
	logger, _ := zapConfig.Build(zap.AddCallerSkip(1), zap.AddCaller())
	zapLogger := &ZapLogger{
		Zap: logger,
	}
	return zapLogger
}

func generateConfig(appConfig *model.Config) zap.Config {
	loggerConfig := zap.NewProductionConfig()
	if (appConfig != nil) && ((model.LoggerConfig{}) != appConfig.LoggerConfig) {
		loggerConfig.Encoding = "json"
		loggerConfig.Level = zap.NewAtomicLevelAt(getLevel(appConfig.LoggerConfig.LogLevel))
		loggerConfig.OutputPaths = []string{"stderr", appConfig.LoggerConfig.LogFilePath}
		loggerConfig.ErrorOutputPaths = []string{"stderr", appConfig.LoggerConfig.LogFilePath}
		loggerConfig.EncoderConfig = zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	}
	return loggerConfig
}

func (zl *ZapLogger) Debug(message string, args ...Argument) {
	zl.Zap.Debug(message, args...)
}

func (zl *ZapLogger) Info(message string, args ...Argument) {
	zl.Zap.Info(message, args...)
}

func (zl *ZapLogger) Warn(message string, args ...Argument) {
	zl.Zap.Warn(message, args...)
}

func (zl *ZapLogger) Error(message string, args ...Argument) {
	zl.Zap.Error(message, args...)
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.DebugLevel
	}
}
