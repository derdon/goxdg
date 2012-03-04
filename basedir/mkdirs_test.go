package basedir

import (
	"testing"
	"os"
	"fmt"
	"io/ioutil"
)

func TestMakeEnvPathFunc(t *testing.T) {
	tmpDir := os.TempDir()
	fun := makeEnvPathFunc(tmpDir)
	dirName, err := ioutil.TempDir(tmpDir, "")
	defer os.Remove(dirName)
	if err != nil {
		t.Errorf(
			"couldn't create temporary directory %s. Reason: %s",
			dirName, err)
	}
	err = fun(dirName)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	f, err := os.Open(fmt.Sprintf("%s/%s", tmpDir, dirName))
	if err != nil {
		t.Errorf("file couldn't be opened, got %v", err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		t.Errorf("couldn't get file info, got %v", err)
	}
	if !fileInfo.IsDir() {
		t.Errorf("touched file %s/%s is not a directory", tmpDir, dirName)
	}
}
