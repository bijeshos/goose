package filesync

import (
	"fmt"
	"github.com/bijeshos/goose/dirutil"
	"github.com/bijeshos/goose/fileutil"
	"path/filepath"
	"strings"
)

func Perform(srcDir string, targetDir string) {

	dirExists, err := dirutil.IsExist(srcDir)
	if err != nil {
		fmt.Println("error occurred: ", err)
		return
	}
	if !dirExists {
		fmt.Println(srcDir, "does not exist. exiting program.")
		return
	}

	dirSame, err := dirutil.IsSame(srcDir, targetDir)
	if err != nil {
		fmt.Println("error occurred: ", err)
		return
	}
	if dirSame {
		fmt.Println("both source and target directories are same. exiting program.")
		return
	}

	files := dirutil.Read(srcDir)
	for _, file := range files {
		fmt.Println("copying: from: ", file)
		fmt.Println("copying: to  : ", filepath.Join(targetDir, filepath.Base(file)))
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		fileutil.CopyFile(file, targetPath)
		fmt.Println("copy completed")
		fmt.Println("------")
	}

}
