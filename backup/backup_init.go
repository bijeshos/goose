package backup

import (
	"fmt"
	"github.com/bijeshos/goose/io/fileutil"
	"path/filepath"
	"strings"
)

func performPrimary(srcDir string, targetDir string) {
	files := fileutil.Read(srcDir)

	//dirutil.MkDirAll(targetDir)

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
