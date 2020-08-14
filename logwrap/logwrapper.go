package logwrap

import (
	"go.uber.org/zap"
	"log"
)

var (
	sugar = initZap()
)

func initZap() *zap.SugaredLogger {
	//configure output log file
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		"./../.logs/goose.log", //todo: need to externalize the log file name
		"stdout",
	}
	logger, err := cfg.Build()

	if err != nil {
		log.Fatal("error occurred while creating log", err)
	}

	sugar := logger.Sugar()
	defer sugar.Sync()
	return sugar

}

func Debugw(msg string, keysAndValues ...interface{}) {
	sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	sugar.Warnw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	sugar.Fatalw(msg, keysAndValues...)
}
