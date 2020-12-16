package log

import "go.uber.org/zap"

var L *zap.Logger

type Logger struct {
	L *zap.Logger
}

func NewLogger(fields ...zap.Field) *Logger {
	return &Logger{
		L.With(fields...),
	}
}
