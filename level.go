package logger

import (
	"go.uber.org/zap/zapcore"
)

var (
	_ zapcore.Core = (*contextLevelCore)(nil)
)

type contextLevelCore struct {
	core  zapcore.Core
	level zapcore.Level
}

func (c *contextLevelCore) Enabled(lvl zapcore.Level) bool {
	return c.level.Enabled(lvl)
}

func (c *contextLevelCore) Level() zapcore.Level {
	return c.level
}

func (c *contextLevelCore) With(fields []zapcore.Field) zapcore.Core {
	return &contextLevelCore{c.core.With(fields), c.level}
}

func (c *contextLevelCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *contextLevelCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	return c.core.Write(ent, fields)
}

func (c *contextLevelCore) Sync() error {
	return c.core.Sync()
}
