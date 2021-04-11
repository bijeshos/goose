package gooseinit

import (
	gfu "github.com/bijeshos/guppy/fileutil"
	"go.uber.org/zap"
	"log"
	"path/filepath"
)

func ZapLogger(logDir string, logFile string) *zap.SugaredLogger {
	log.Println("initializing zap logger")
	gfu.CreateFile(logDir, logFile)
	//configure output log file
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		filepath.Join(logDir, logFile),
		"stdout",
	}
	logger, err := cfg.Build()
	if err != nil {
		log.Fatal("error occurred while creating log", err)
	}

	sugar := logger.Sugar()
	defer sugar.Sync()
	zap.ReplaceGlobals(logger)
	return sugar
}
