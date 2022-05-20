package cli

import (
	"errors"
	"os"
)

// NeedFilename returns the filename from os.Args.
func NeedFilename() (string, error) {
	if len(os.Args) != 2 {
		return "", errors.New("please provide a filename")
	}
	return os.Args[1], nil
}
