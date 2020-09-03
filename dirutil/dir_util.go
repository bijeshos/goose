package dirutil

import (
	"github.com/bijeshos/goose/arrayutil"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

//Read to read from dir
func Read(srcDir string, ignoreList []string) ([]string, error) {
	zap.S().Infow("reading files", "from", srcDir)
	var files []string
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && arrayutil.IsPresent(ignoreList, info.Name()) {
			zap.S().Infow("ignoring dir", "dir name", info.Name())
			return filepath.SkipDir
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

//MkDirAll to create directories
func MkDirAll(targetDir string) error {
	_, err := os.Stat(targetDir)
	if err != nil {
		return err
	}

	if os.IsNotExist(err) {
		zap.S().Infow("creating directory", "dir", targetDir)
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

//IsExist to check if the directory exists
func IsExist(dir string) (bool, error) {
	//check whether directory exists
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

//IsSame to check if both directories are same
func IsSame(srcDir string, targetDir string) (bool, error) {

	srcInfo, srcErr := os.Stat(srcDir)
	if srcErr != nil {
		return false, srcErr
	}
	targetInfo, targetErr := os.Stat(targetDir)
	if targetErr != nil {
		if os.IsNotExist(targetErr) {
			return false, nil
		}
		return false, targetErr
	}
	// check if source and target dirs are same
	if os.SameFile(srcInfo, targetInfo) {
		return true, nil
	} else {
		return false, nil
	}
}
