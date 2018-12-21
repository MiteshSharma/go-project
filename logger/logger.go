package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Argument = zapcore.Field

var Int = zap.Int
var String = zap.String
var Error = zap.Error

type Logger interface {
	Debug(message string, args ...Argument)
	Info(message string, args ...Argument)
	Warn(message string, args ...Argument)
	Error(message string, args ...Argument)
}
