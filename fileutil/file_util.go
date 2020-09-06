package fileutil

import (
	"bufio"
	"github.com/bijeshos/goose/dirutil"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//CreateFile
func CreateFile(targetDir, fileName string) error {
	dirErr := dirutil.MkDirAll(targetDir)
	if dirErr != nil {
		return dirErr
	}
	//_, err := os.OpenFile(filepath.Join(targetDir, fileName), os.O_RDWR|os.O_CREATE, 0666)
	_, err := os.Create(filepath.Join(targetDir, fileName))
	if err != nil {
		//zap.S().Fatalw("error occurred: ", "error", err)
		return err
	}
	return nil
}

//CopyFile
func CopyFile(srcPath, targetPath string, forceReplace bool) error {

	//open source file
	src, err := os.Open(srcPath)
	if err != nil {
		//zap.S().Errorw("error occurred: ", "error", err)
		return err
	}
	defer src.Close()

	//create sub directories at target if needed
	targetSubDir := filepath.Dir(targetPath)
	dirutil.MkDirAll(targetSubDir)

	//open target file
	target, err := os.OpenFile(targetPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		//zap.S().Errorw("error occurred: ", "error", err)
		return err
	}
	defer target.Close()
	var proceedCopy bool = false

	if forceReplace {
		proceedCopy = true
	} else {
		isSame, _ := isSameMetadata(srcPath, targetPath)
		proceedCopy = !isSame
	}
	if proceedCopy {
		zap.S().Infow("copying", "from", srcPath, "to", targetPath)
		//perform copying
		_, err = io.Copy(target, src)
		if err != nil {
			//zap.S().Errorw("error occurred: ", "error", err)
			return err
		}
		zap.S().Infow("copy completed")
		zap.S().Infow("------")
	} else {
		zap.S().Debugw("Skipping copy", "file", srcPath, "reason", "Another file with same name and size exists at target")
	}
	return nil
}

func isSameMetadata(srcFilePath string, targetFilePath string) (bool, error) {
	srcFileInfo, srcErr := os.Stat(srcFilePath)
	targetFileInfo, targetErr := os.Stat(targetFilePath)

	if srcErr != nil {
		return false, srcErr
	}

	//zap.S().Debugw("src file", "info", srcFileInfo)

	if targetErr != nil {
		if os.IsNotExist(targetErr) {
			return false, targetErr
		}
	}
	//zap.S().Debugw("target file", "info", targetFileInfo)

	if strings.Compare(srcFileInfo.Name(), targetFileInfo.Name()) == 0 && srcFileInfo.Size() == targetFileInfo.Size() {
		return true, nil
	} else {
		return false, nil
	}

}

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		//zap.S().Errorw("error occurred: ", "error", err)
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//fmt.Println(lines)
	if err := scanner.Err(); err != nil {
		//zap.S().Errorw("error occurred: ", "error", err)
		return nil, err
	}
	return lines, nil
}
