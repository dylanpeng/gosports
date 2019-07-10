package logger

import (
	"errors"
	zapLog "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	Config   *Config
	log      *zapLog.Logger
	logSugar *zapLog.SugaredLogger
}

func (l *Logger) InitLogger() error {
	if l.Config == nil {
		return errors.New("param error : config is nil")
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(l.Config.GetZapEncoderConfig()),                                    // 编码
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(l.Config.GetHook())), // 打印输出端
		GetZapLevel(l.Config.LogLevel, zapcore.DebugLevel),                                           // 打印日志级别                            // 日志
	)

	// 开启开发模式，调用跟踪
	caller := zapLog.AddCaller()
	// 开启堆栈跟踪
	stacktrace := zapLog.AddStacktrace(GetZapLevel(l.Config.StacktraceLevel, zapcore.ErrorLevel))
	// 开启文件及行号
	development := zapLog.Development()
	// 设置初始化字段
	filed := zapLog.Fields(zapLog.String("serviceName", l.Config.ServiceName))

	// 构建日志
	l.log = zapLog.New(core, caller, stacktrace, filed, development)
	l.logSugar = l.log.Sugar()
	return nil
}

func (l *Logger) Debug(msg string, fields ...zapLog.Field) {
	l.log.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zapLog.Field) {
	l.log.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zapLog.Field) {
	l.log.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zapLog.Field) {
	l.log.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...zapLog.Field) {
	l.log.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zapLog.Field) {
	l.log.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zapLog.Field) {
	l.log.Fatal(msg, fields...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logSugar.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	l.logSugar.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logSugar.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logSugar.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logSugar.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logSugar.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logSugar.Fatalf(template, args...)
}
