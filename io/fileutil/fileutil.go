package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

//Read to read from dir
func Read(srcDir string) {
	fmt.Println("reading files from: ", srcDir)

	var files []string

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		fmt.Println("----")
		fmt.Println("path", path)
		fmt.Println("file-info", info)
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
