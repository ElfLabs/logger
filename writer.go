package logger

import (
	"bytes"

	"go.uber.org/zap"
)

type Writer struct {
	Fields  []zap.Field
	LogFunc func(msg string, fields ...zap.Field)
}

func (l *Writer) Write(p []byte) (int, error) {
	p = bytes.TrimSpace(p)
	l.LogFunc(string(p), l.Fields...)
	return len(p), nil
}
