package fileutil

import "testing"

func TestRead(t *testing.T) {

	Read(".")

	got := 2
	if got != 1 {
		t.Errorf("Abs(-1) = %d; want 1", got)
	}
}
