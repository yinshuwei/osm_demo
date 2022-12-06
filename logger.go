package main

import "go.uber.org/zap"

// Logger 适配zap logger
type Logger struct {
	zapLogger *zap.Logger
}

func (l *Logger) fields(data map[string]string) []zap.Field {
	var fields []zap.Field
	for key, val := range data {
		fields = append(fields, zap.String(key, val))
	}
	return fields
}

func (l *Logger) Error(msg string, data map[string]string) {
	if l == nil || l.zapLogger == nil {
		return
	}
	l.zapLogger.Error(msg, l.fields(data)...)
}

func (l *Logger) Info(msg string, data map[string]string) {
	if l == nil || l.zapLogger == nil {
		return
	}
	l.zapLogger.Info(msg, l.fields(data)...)
}

func (l *Logger) Warn(msg string, data map[string]string) {
	if l == nil || l.zapLogger == nil {
		return
	}
	l.zapLogger.Warn(msg, l.fields(data)...)
}
