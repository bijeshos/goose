package dirsync

import (
	"fmt"
	"github.com/bijeshos/goose/logwrap"
	"os"
	"testing"
)

func TestRead(t *testing.T) {

	fmt.Println(os.Getwd())

	logwrap.InitZapLogger("./.logs", "goose.log")

	srcDir := "/home/bos/1-bos/tmp/test/goose/src-1"
	targetDir := "/home/bos/1-bos/tmp/test/goose/target-1"
	//targetDir := "/home/bos/1-bos/tmp/test/goose/src-1"

	Perform(srcDir, targetDir)

}
