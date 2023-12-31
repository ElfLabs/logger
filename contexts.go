package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	loggerCtx    struct{}
	fieldCtx     struct{}
	fieldFuncCtx struct{}
	levelCtx     struct{}
)

type (
	FieldFunc func(ctx context.Context) zap.Field
)

func NewContext(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		logger = L()
	}
	return context.WithValue(ctx, loggerCtx{}, logger.WithContext(ctx))
}

func FromContext(ctx context.Context) Logger {
	logger, _ := ctx.Value(loggerCtx{}).(Logger)
	if logger == nil {
		logger = L()
	}
	return logger.WithContext(ctx)
}

func NewFieldContext(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}
	exists := FieldsFromContext(ctx)
	if len(exists) == 0 {
		return context.WithValue(ctx, fieldCtx{}, fields)
	}
	return context.WithValue(ctx, fieldCtx{}, append(exists, fields...))
}

func WithField(ctx context.Context, fields ...zap.Field) context.Context {
	return NewFieldContext(ctx, fields...)
}

func FieldsFromContext(ctx context.Context) []zap.Field {
	fields, _ := ctx.Value(fieldCtx{}).([]zap.Field)
	fs := FieldFuncFromContext(ctx)
	if len(fs) == 0 {
		return fields
	}
	for _, f := range fs {
		fields = append(fields, f(ctx))
	}
	return fields
}

func NewFieldFuncContext(ctx context.Context, fs ...FieldFunc) context.Context {
	if len(fs) == 0 {
		return ctx
	}
	exists := FieldFuncFromContext(ctx)
	if len(exists) == 0 {
		return context.WithValue(ctx, fieldFuncCtx{}, fs)
	}
	return context.WithValue(ctx, fieldFuncCtx{}, append(exists, fs...))
}

func FieldFuncFromContext(ctx context.Context) []FieldFunc {
	fs, _ := ctx.Value(fieldFuncCtx{}).([]FieldFunc)
	if len(fs) != 0 {
		return fs
	}
	return nil
}

func WithLevel(ctx context.Context, level zapcore.Level) context.Context {
	return context.WithValue(ctx, levelCtx{}, level)
}

func LevelFromContext(ctx context.Context) zapcore.Level {
	level, ok := ctx.Value(levelCtx{}).(zapcore.Level)
	if ok {
		return level
	}
	return zapcore.InvalidLevel
}
