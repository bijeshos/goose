package filesync

import (
	"fmt"
	"github.com/bijeshos/goose/fileutil"
	"path/filepath"
	"strings"
)

func Perform(srcDir string, targetDir string) {
	files := fileutil.Read(srcDir)
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
