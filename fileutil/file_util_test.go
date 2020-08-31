package fileutil

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	t.Skip("testing this test for now")
	//Read(".")
	//placeholder

	got := 2
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}

func TestReadFile(t *testing.T) {
	filePath := "/home/bos/.goose/goose.ignore"
	lines := ReadFile(filePath)
	fmt.Println(lines)
	fmt.Println("--")
}
