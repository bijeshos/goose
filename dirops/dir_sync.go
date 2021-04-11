package dirops

import (
	gdu "github.com/bijeshos/guppy/dirutil"
	gfu "github.com/bijeshos/guppy/fileutil"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

func Move(srcDir string, targetDir string) {
	zap.S().Infow("initiating directory move", "source", srcDir, "target", targetDir)
	checkSrcDirExists(srcDir)
	checkSrcAndTargetDirDiffer(srcDir, targetDir)
	moveFiles(srcDir, targetDir)
	//deleteEmptyDirs(srcDir)
}

func Sync(srcDir string, targetDir string) {
	zap.S().Infow("initiating directory sync", "source", srcDir, "target", targetDir)
	checkSrcDirExists(srcDir)
	checkSrcAndTargetDirDiffer(srcDir, targetDir)
	copyFiles(srcDir, targetDir)
}

func checkSrcDirExists(srcDir string) {
	zap.S().Infow("verifying existence of source directory")
	//check whether source directory exists
	dirExists, err := gdu.IsExist(srcDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if !dirExists {
		zap.S().Fatalw("dir does not exist", "dir", srcDir)
	}
}

func checkSrcAndTargetDirDiffer(srcDir string, targetDir string) {
	zap.S().Infow("verifying source and target directories are different")
	//check whether source and target directories are same
	dirSame, err := gdu.IsSame(srcDir, targetDir)
	if err != nil {
		zap.S().Fatalw("error occurred: ", "error", err)
	}
	if dirSame {
		zap.S().Fatalw("both source and target directories are same. exiting program.")
	}
}

func copyFiles(srcDir string, targetDir string) {

	ignoreList := getIgnoreDirs()

	zap.S().Infow("retrieving file details from source directory")
	files, dirReadErr := gdu.ReadFiles(srcDir, ignoreList)
	if dirReadErr != nil {
		zap.S().Fatalw("Error occurred while reading src dir", "error", dirReadErr)
	}

	zap.S().Infow("initiating file copying")
	for _, file := range files {
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		gfu.CopyFile(file, targetPath, false)

	}
}

func moveFiles(srcDir string, targetDir string) {
	ignoreList := getIgnoreDirs()
	zap.S().Infow("retrieving file details from source directory")
	files, dirReadErr := gdu.ReadFiles(srcDir, ignoreList)
	if dirReadErr != nil {
		zap.S().Fatalw("Error occurred while reading src dir", "error", dirReadErr)
	}
	zap.S().Infow("initiating file moving")
	for _, file := range files {
		relativePath := strings.Replace(file, srcDir, "", 1)
		targetPath := filepath.Join(targetDir, relativePath)
		gfu.MoveFile(file, targetPath, false)

	}
}

func getIgnoreDirs() []string {
	zap.S().Infow("reading ignore file list")
	ignoreFileName := viper.GetString("sync.ignore-file-list")
	ignoreList, fileReadErr := gfu.ReadFileContent(ignoreFileName)
	if fileReadErr != nil {
		zap.S().Errorw("Error occurred while reading ignore file", fileReadErr)
	}
	return ignoreList

}

/*func deleteEmptyDirs(srcDir string) {
	ignoreList := getIgnoreDirs()
	dirList, _ := gdu.ReadDirs(srcDir, ignoreList)
	sort.Reverse(dirList)
}*/
