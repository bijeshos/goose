package dirsync

import (
	"github.com/bijeshos/goose/dirutil"
	"github.com/bijeshos/goose/fileutil"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

func Perform(srcDir string, targetDir string) {
	zap.S().Infow("initiating directory sync", "source", srcDir, "target", targetDir)
	checkSrcDirExists(srcDir)
	checkSrcAndTargetDirDiffer(srcDir, targetDir)
	copyFiles(srcDir, targetDir)
}

func checkSrcDirExists(srcDir string) {
	//check whether source directory exists
	dirExists, err := dirutil.IsExist(srcDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if !dirExists {
		zap.S().Fatalw("dir does not exist", "dir", srcDir)
	}
}

func checkSrcAndTargetDirDiffer(srcDir string, targetDir string) {
	//check whether source and target directories are same
	dirSame, err := dirutil.IsSame(srcDir, targetDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if dirSame {
		zap.S().Fatalw("both source and target directories are same. exiting program.")
	}
}

func copyFiles(srcDir string, targetDir string) {
	ignoreFileName := viper.GetString("sync.ignore-file-list")
	ignoreList, fileReadErr := fileutil.ReadFile(ignoreFileName)
	if fileReadErr != nil {
		zap.S().Errorw("Error occurred while reading ignore file", fileReadErr)
	}

	files, dirReadErr := dirutil.Read(srcDir, ignoreList)
	if dirReadErr != nil {
		zap.S().Fatalw("Error occurred while reading src dir", "error", dirReadErr)
	}

	for _, file := range files {
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		fileutil.CopyFile(file, targetPath, false)

	}
}
