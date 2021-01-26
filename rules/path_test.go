package rules

import (
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v5"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	gofakeit.Seed(0)
	defer filet.CleanUp(t)
	dirName := gofakeit.BS()
	filename := gofakeit.BS()
	path := fmt.Sprintf(dirName)
	if IsFile(path) {
		t.Errorf("expected file %s to not exist", filename)
	}
	filet.File(t, path, "")
	if !IsFile(path) {
		t.Errorf("expected file %s to exist", filename)
	}
}

func TestDirExists(t *testing.T) {
	gofakeit.Seed(2)
	currentDir, err := os.Getwd()
	if err != nil {
		t.Skipf("can't find current directory, skipping directory test: %s", err.Error())
	}
	if !IsDir(currentDir) {
		t.Errorf("expected currentDir %s to exist", currentDir)
	}
	if IsDir(gofakeit.AppName()) {
		t.Errorf("expected dir %s to not exists exist", currentDir)
	}
}

func TestIsValidPath(t *testing.T) {
	gofakeit.Seed(1)
	defer filet.CleanUp(t)
	dirName := gofakeit.BS()
	filename := gofakeit.BS()
	path := fmt.Sprintf(dirName)
	if IsValidPath(path) {
		t.Errorf("expected file %s to not exist", filename)
	}
	filet.File(t, path, "")
	if !IsValidPath(path) {
		t.Errorf("expected file %s to exist", filename)
	}
}
