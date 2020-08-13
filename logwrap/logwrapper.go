package logwrap

import (
	"go.uber.org/zap"
)

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Warnw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Fatalw(msg, keysAndValues...)
}
