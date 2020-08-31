package logwrap

import (
	"github.com/bijeshos/goose/fileutil"
	"go.uber.org/zap"
	"log"
	"path/filepath"
)

func InitZapLogger(logDir string, logFile string) *zap.SugaredLogger {
	log.Println("initializing zap logger")
	fileutil.CreateFile(logDir, logFile)
	//configure output log file
	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		filepath.Join(logDir, logFile),
		"stdout",
	}
	logger, err := cfg.Build()
	//"./../.logs/goose.log", //todo: need to externalize the log file name
	if err != nil {
		log.Fatal("error occurred while creating log", err)
	}

	sugar := logger.Sugar()
	defer sugar.Sync()
	zap.ReplaceGlobals(logger)
	return sugar
}
