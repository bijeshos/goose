package dirutil

import (
	"fmt"
	"os"
)

//Read to read from dir
func Read() {
	fmt.Println("inside dirutil:read")
}

//MkDirAll to create directories
func MkDirAll(targetDir string) {
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Println("creating directory: ", targetDir)
		os.MkdirAll(targetDir, os.ModePerm)
	}
}
