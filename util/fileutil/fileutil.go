package fileutil

import (
	"fmt"
	"github.com/bijeshos/goose/util/dirutil"
	"io"
	"log"
	"os"
	"path/filepath"
)

//Read to read from dir
func Read(srcDir string) []string {
	fmt.Println("reading files from: ", srcDir)
	var files []string
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

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
