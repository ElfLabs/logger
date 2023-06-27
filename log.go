package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	L().Log(lvl, msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	L().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	L().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	L().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	L().Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	L().DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	L().Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	L().Fatal(msg, fields...)
}

func Core() zapcore.Core {
	return L().Core()
}

func Sync() error {
	return L().Sync()
}

func Sugar() *zap.SugaredLogger {
	return L().Sugar()
}

func Named(s string) *zap.Logger {
	return L().Named(s)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return L().WithOptions(opts...)
}

func With(fields ...zap.Field) *zap.Logger {
	return L().With(fields...)
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return L().Check(lvl, msg)
}

func WithContext(ctx context.Context) Logger {
	return L().WithContext(ctx)
}

func OnError(err error, fields ...zap.Field) IZapLog {
	return L().OnError(err, fields...)
}
