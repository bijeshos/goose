package fileutil

import (
	"github.com/bijeshos/goose/dirutil"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
)

//CreateFile
func CreateFile(targetDir, fileName string) {
	dirutil.MkDirAll(targetDir)
	_, err := os.OpenFile(filepath.Join(targetDir, fileName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
}

//CopyFile
func CopyFile(srcPath, targetPath string) {
	//open source file
	src, err := os.Open(srcPath)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	defer src.Close()

	//create sub directories at target if needed
	targetSubDir := filepath.Dir(targetPath)
	dirutil.MkDirAll(targetSubDir)

	//open target file
	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	defer target.Close()

	if preValidation(srcPath, targetPath) {
		//perform copying
		_, err = io.Copy(target, src)
		if err != nil {
			zap.S().Fatalw("error occurred: ", "error", err)
		}
	}

}

func preValidation(srcFilePath string, targetFilePath string) bool {

	return true //todo
}
