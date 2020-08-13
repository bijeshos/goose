package fileutil

import (
	"github.com/bijeshos/goose/dirutil"
	"github.com/bijeshos/goose/logwrap"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(srcPath, targetPath string) {
	//open source file
	src, err := os.Open(srcPath)
	if err != nil {
		logwrap.Fatalw("error occurred: ", "error", err)
	}
	defer src.Close()

	//create sub directories if needed
	targetSubDir := filepath.Dir(targetPath)
	dirutil.MkDirAll(targetSubDir)

	//open target file
	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logwrap.Fatalw("error occurred: ", "error", err)
	}
	defer target.Close()

	//perform copying
	_, err = io.Copy(target, src)
	if err != nil {
		logwrap.Fatalw("error occurred: ", "error", err)
	}
}
