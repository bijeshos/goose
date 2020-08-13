package filesync

import (
	"github.com/bijeshos/goose/dirutil"
	"github.com/bijeshos/goose/fileutil"
	"github.com/bijeshos/goose/logwrap"
	"path/filepath"
	"strings"
)

func Perform(srcDir string, targetDir string) {
	logwrap.Infow("initiating sync")
	dirExists, err := dirutil.IsExist(srcDir)
	if err != nil {
		logwrap.Fatalw("error occurred: ", "error", err)
		return
	}
	if !dirExists {
		logwrap.Fatalw("dir does not exist", "dir", srcDir)
		return
	}

	dirSame, err := dirutil.IsSame(srcDir, targetDir)
	if err != nil {
		logwrap.Fatalw("error occurred: ", "error", err)
		return
	}
	if dirSame {
		logwrap.Fatalw("both source and target directories are same. exiting program.")
		return
	}

	files := dirutil.Read(srcDir)
	for _, file := range files {
		logwrap.Infow("copying", "from", file)
		logwrap.Infow("copying", "to", filepath.Join(targetDir, filepath.Base(file)))
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		fileutil.CopyFile(file, targetPath)
		logwrap.Infow("copy completed")
		logwrap.Infow("------")
	}

}
