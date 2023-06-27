package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type IZapLog interface {
	Log(lvl zapcore.Level, msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type ZapLogger interface {
	IZapLog
	Core() zapcore.Core
	Sync() error
	Sugar() *zap.SugaredLogger
	Named(s string) *zap.Logger
	WithOptions(opts ...zap.Option) *zap.Logger
	With(fields ...zap.Field) *zap.Logger
	Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
}

type Logger interface {
	ZapLogger
	WithContext(ctx context.Context) Logger
	OnError(err error, fields ...zap.Field) IZapLog
}
