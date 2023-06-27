package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type nopZapLog struct{}

func NewNopIZapLog() IZapLog {
	return nopZapLog{}
}

func (n nopZapLog) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Debug(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Info(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Warn(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Error(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) DPanic(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Panic(msg string, fields ...zap.Field) {
	return
}

func (n nopZapLog) Fatal(msg string, fields ...zap.Field) {
	return
}

type nopILogger struct {
	*zap.Logger
}

func (n nopILogger) WithContext(ctx context.Context) Logger {
	return n
}

func (n nopILogger) OnError(err error, fields ...zap.Field) IZapLog {
	return NewNopIZapLog()
}

func NewNopILogger() Logger {
	return nopILogger{
		Logger: zap.NewNop(),
	}
}
