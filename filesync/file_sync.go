package filesync

import (
	"github.com/bijeshos/goose/dirutil"
	"github.com/bijeshos/goose/fileutil"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

func Perform(srcDir string, targetDir string) {
	zap.S().Infow("initiating sync")
	checkSrcDirExists(srcDir)
	checkSrcAndTargetDirDiffer(srcDir, targetDir)
	copyFiles(srcDir, targetDir)

}

func checkSrcDirExists(srcDir string) {
	//check whether source directory exists
	dirExists, err := dirutil.IsExist(srcDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if !dirExists {
		zap.S().Fatalw("dir does not exist", "dir", srcDir)
	}
}

func checkSrcAndTargetDirDiffer(srcDir string, targetDir string) {
	//check whether source and target directories are same
	dirSame, err := dirutil.IsSame(srcDir, targetDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if dirSame {
		zap.S().Fatalw("both source and target directories are same. exiting program.")
	}
}

func copyFiles(srcDir string, targetDir string) {
	files := dirutil.Read(srcDir)
	for _, file := range files {
		zap.S().Infow("copying", "from", file)
		zap.S().Infow("copying", "to", filepath.Join(targetDir, filepath.Base(file)))
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		fileutil.CopyFile(file, targetPath, false)
		zap.S().Infow("copy completed")
		zap.S().Infow("------")
	}
}
