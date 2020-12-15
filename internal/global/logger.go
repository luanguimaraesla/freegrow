package global

import "go.uber.org/zap"

var GlobalLogger *zap.Logger

type Logger struct {
	L *zap.Logger
}

func NewLogger(fields ...zap.Field) *Logger {
	return &Logger{
		GlobalLogger.With(fields...),
	}
}
