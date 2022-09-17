package file

import (
	"errors"
	"os"
	"testing"
)

func TestCreateDirIfNotExists(t *testing.T) {
	err := CreateDirIfNotExists("./testdata/testDir")
	if err != nil {
		if !errors.Is(err, ErrCreateDir) {
			t.Error(err)
		}
	}

	err = CreateDirIfNotExists("./testdata/testDir")
	if err != nil {
		if !errors.Is(err, ErrCreateDir) {
			t.Error(err)
		}
	}

	err = CreateDirIfNotExists("/testDir")
	if err != nil {
		if !errors.Is(err, ErrCreateDir) {
			t.Error(err)
		}
	}

	_ = os.Remove("./testdata/testDir")
}
