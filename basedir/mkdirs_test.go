package basedir

import (
	"testing"
	"os"
)

func TestMakeEnvPathFunc(t *testing.T) {
	fun := makeEnvPathFunc("/tmp")
	err := fun("foo/bar")
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	f, err := os.Open("/tmp/foo/bar")
	if err != nil {
		t.Errorf("file couldn't be opened, got %v", err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		t.Errorf("couldn't get file info, got %v", err)
	}
	if !fileInfo.IsDir() {
		t.Errorf("touched file /tmp/foo/bar is not a directory")
	}
}
