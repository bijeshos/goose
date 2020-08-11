package backup

import "testing"

func TestRead(t *testing.T) {

	srcDir := "/home/bos/1-bos/tmp/test/goose/src-1"
	targetDir := "/home/bos/1-bos/tmp/test/goose/target-1"

	performPrimary(srcDir, targetDir)

}
