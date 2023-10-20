package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type wrapLogger zap.Logger

func New(logger *zap.Logger) Logger {
	return (*wrapLogger)(logger)
}

func (log *wrapLogger) Sugar() *zap.SugaredLogger {
	return ((*zap.Logger)(log)).Sugar()
}

func (log *wrapLogger) Named(s string) *zap.Logger {
	return ((*zap.Logger)(log)).Named(s)
}

func (log *wrapLogger) WithOptions(opts ...zap.Option) *zap.Logger {
	return ((*zap.Logger)(log)).WithOptions(opts...)
}

func (log *wrapLogger) With(fields ...zap.Field) *zap.Logger {
	return ((*zap.Logger)(log)).With(fields...)
}

func (log *wrapLogger) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return ((*zap.Logger)(log)).Check(lvl, msg)
}

func (log *wrapLogger) logger() *zap.Logger {
	return ((*zap.Logger)(log)).WithOptions(zap.AddCallerSkip(1))
}

func (log *wrapLogger) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	log.logger().Log(lvl, msg, fields...)
}

func (log *wrapLogger) Debug(msg string, fields ...zap.Field) {
	log.logger().Debug(msg, fields...)
}

func (log *wrapLogger) Info(msg string, fields ...zap.Field) {
	log.logger().Info(msg, fields...)
}

func (log *wrapLogger) Warn(msg string, fields ...zap.Field) {
	log.logger().Warn(msg, fields...)
}

func (log *wrapLogger) Error(msg string, fields ...zap.Field) {
	log.logger().Error(msg, fields...)
}

func (log *wrapLogger) DPanic(msg string, fields ...zap.Field) {
	log.logger().DPanic(msg, fields...)
}

func (log *wrapLogger) Panic(msg string, fields ...zap.Field) {
	log.logger().Panic(msg, fields...)
}

func (log *wrapLogger) Fatal(msg string, fields ...zap.Field) {
	log.logger().Fatal(msg, fields...)
}

func (log *wrapLogger) Sync() error {
	return ((*zap.Logger)(log)).Sync()
}

func (log *wrapLogger) Core() zapcore.Core {
	return ((*zap.Logger)(log)).Core()
}

func (log *wrapLogger) WithContext(ctx context.Context) Logger {
	level := LevelFromContext(ctx)
	fields := FieldsFromContext(ctx)
	switch {
	case level == nil && len(fields) == 0:
		return log
	case level != nil && len(fields) == 0:
		return New(((*zap.Logger)(log)).WithOptions(zap.IncreaseLevel(level)))
	case level == nil && len(fields) != 0:
		return New(((*zap.Logger)(log)).With(fields...))
	case level != nil && len(fields) != 0:
		return New(((*zap.Logger)(log)).WithOptions(zap.IncreaseLevel(level), zap.Fields(fields...)))
	default:
		return log
	}
}

func (log *wrapLogger) OnError(err error, fields ...zap.Field) IZapLog {
	switch {
	case err == nil:
		return NewNopIZapLog()
	case len(fields) == 0:
		return ((*zap.Logger)(log)).With(zap.Error(err))
	default:
		return ((*zap.Logger)(log)).With(append(fields, zap.Error(err))...)
	}
}
