package fileutil

import (
	"fmt"
	"github.com/bijeshos/goose/io/dirutil"
	"io"
	"log"
	"os"
	"path/filepath"
)

//Read to read from dir
func Read(srcDir string) []string {
	fmt.Println("reading files from: ", srcDir)

	var files []string
	var dirs []string

	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		//fmt.Println("----")
		if info.IsDir() {
			//fmt.Println("dir: ", path)
			dirs = append(dirs, path)
		} else {
			/*fmt.Println("file: ", path)
			fmt.Println("size: ", info.Size(), "Bytes")
			fmt.Println("modified time: ", info.ModTime())*/
			files = append(files, path)
		}
		/*fmt.Println("path: ", path)
		fmt.Println("file-info: ", info)
		if filepath.Ext(path) == "txt"{
			return nil
		}*/

		return nil
	})
	if err != nil {
		panic(err)
	}
	/*for _, file := range files {
		fmt.Println(file)
	}*/
	return files
}

func CopyFile(srcFilePath, targetFilePath string) {
	//open source file
	src, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	//create sub directories if needed
	targetSubDir := filepath.Dir(targetFilePath)
	dirutil.MkDirAll(targetSubDir)

	//open target file
	target, err := os.OpenFile(targetFilePath, os.O_RDWR|os.O_CREATE, 0666)
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
