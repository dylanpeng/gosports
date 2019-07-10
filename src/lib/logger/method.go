package logger

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

var mapLogLevel = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func GetZapLevel(levelKey string, defaultLevel zapcore.Level) zapcore.Level {
	levelKey = strings.ToLower(levelKey)
	result, ok := mapLogLevel[levelKey]
	if ok {
		return result
	}
	return defaultLevel
}
