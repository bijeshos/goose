package dirutil

import (
	"fmt"
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

//MkDirAll to create directories
func MkDirAll(targetDir string) {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Println("creating directory: ", targetDir)
		os.MkdirAll(targetDir, os.ModePerm)
	}
}

//IsExist to check if the directory exists
func IsExist(dir string) (bool, error) {
	//check whether directory exists
	_, err := os.Stat(dir)
	if err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, nil
}