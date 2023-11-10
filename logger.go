package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	*zap.Logger

	ctx context.Context
}

func New(zapLogger *zap.Logger) Logger {
	return &logger{Logger: zapLogger}
}

func (log *logger) applyContext() *zap.Logger {
	if log.ctx == nil {
		return log.Logger.WithOptions(zap.AddCallerSkip(1))
	}

	logger := log.Logger

	level := LevelFromContext(log.ctx)
	if level != zapcore.InvalidLevel {
		logger = logger.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return &contextLevelCore{
				core:  core,
				level: level,
			}
		}))
	}

	fields := FieldsFromContext(log.ctx)
	if len(fields) != 0 {
		return logger.WithOptions(zap.AddCallerSkip(1)).WithLazy(fields...)
	}
	return logger.WithOptions(zap.AddCallerSkip(1))
}

func (log *logger) Log(lvl zapcore.Level, msg string, fields ...zap.Field) {
	log.applyContext().Log(lvl, msg, fields...)
}

func (log *logger) Debug(msg string, fields ...zap.Field) {
	log.applyContext().Debug(msg, fields...)
}

func (log *logger) Info(msg string, fields ...zap.Field) {
	log.applyContext().Info(msg, fields...)
}

func (log *logger) Warn(msg string, fields ...zap.Field) {
	log.applyContext().Warn(msg, fields...)
}

func (log *logger) Error(msg string, fields ...zap.Field) {
	log.applyContext().Error(msg, fields...)
}

func (log *logger) DPanic(msg string, fields ...zap.Field) {
	log.applyContext().DPanic(msg, fields...)
}

func (log *logger) Panic(msg string, fields ...zap.Field) {
	log.applyContext().Panic(msg, fields...)
}

func (log *logger) Fatal(msg string, fields ...zap.Field) {
	log.applyContext().Fatal(msg, fields...)
}

func (log *logger) WithContext(ctx context.Context) Logger {
	return &logger{
		Logger: log.WithOptions(), // clone
		ctx:    ctx,
	}
}

func (log *logger) OnError(err error, fields ...zap.Field) IZapLog {
	switch {
	case err == nil:
		return NewNopIZapLog()
	case len(fields) == 0:
		return log.With(zap.Error(err))
	default:
		return log.With(append(fields, zap.Error(err))...)
	}
}
