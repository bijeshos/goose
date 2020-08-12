package fileutil

import (
	"github.com/bijeshos/goose/dirutil"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CopyFile(srcPath, targetPath string) {
	//open source file
	src, err := os.Open(srcPath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	//create sub directories if needed
	targetSubDir := filepath.Dir(targetPath)
	dirutil.MkDirAll(targetSubDir)

	//open target file
	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer target.Close()

	//perform copying
	_, err = io.Copy(target, src)
	if err != nil {
		log.Fatal(err)
	}
}
