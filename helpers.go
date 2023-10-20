package logger

import (
	"context"

	"go.uber.org/zap"
)

func NewGlobals(zapLogger *zap.Logger) {
	zap.ReplaceGlobals(zapLogger)
	ReplaceGlobals(New(zapLogger))
}

func NewNamedLogger(name string) Logger {
	return New(Named(name))
}

func NewNamedSugarLogger(name string) *zap.SugaredLogger {
	return NewNamedLogger(name).Sugar()
}

func SugaredLoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	return FromContext(ctx).Sugar()
}

func NamedSugaredLoggerFromContext(ctx context.Context, name string) *zap.SugaredLogger {
	return FromContext(ctx).Sugar().Named(name)
}
