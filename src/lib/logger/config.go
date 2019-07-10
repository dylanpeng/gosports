package logger

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	FileName        string `toml:"log_filename"`
	MaxSize         int    `toml:"log_file_maxsize"` // unit M
	TimeKey         string `toml:"log_time_key"`
	LevelKey        string `toml:"log_level_key"`
	NameKey         string `toml:"log_name_key"`
	CallerKey       string `toml:"log_caller_key"`
	MessageKey      string `toml:"log_message_key"`
	StacktraceKey   string `toml:"stacktrace_key"`
	LogLevel        string `toml:"log_level"`
	StacktraceLevel string `toml:"stacktrace_level"`
	ServiceName     string `toml:"service_name"`
}

func (c *Config) GetHook() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:  c.FileName,
		MaxSize:   c.MaxSize,
		LocalTime: true,
	}
}

func (c *Config) GetZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        c.TimeKey,
		LevelKey:       c.LevelKey,
		NameKey:        c.NameKey,
		CallerKey:      c.CallerKey,
		MessageKey:     c.MessageKey,
		StacktraceKey:  c.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,  // Level编码器 根据level输出不同颜色
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
}
