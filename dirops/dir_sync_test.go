package dirops

import (
	"fmt"
	"github.com/bijeshos/goose/gooseinit"
	"os"
	"testing"
)

func TestRead(t *testing.T) {

	fmt.Println(os.Getwd())

	gooseinit.ZapLogger("./.logs", "goose.log")

	srcDir := "/home/bos/1-bos/tmp/test/goose/src-1"
	targetDir := "/home/bos/1-bos/tmp/test/goose/target-1"
	//targetDir := "/home/bos/1-bos/tmp/test/goose/src-1"

	Sync(srcDir, targetDir)

}
