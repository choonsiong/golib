// Package file provides helpers to work with files.
package file

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Exists return true if the filename is an executable and exists in the user PATH.
func Exists(filename string) (bool, error) {
	found := false

	path := os.Getenv("PATH")
	pathSlice := strings.Split(path, ":")

	for _, dir := range pathSlice {
		fullPath := dir + "/" + filename
		fileInfo, err := os.Stat(fullPath)
		if err == nil { // found!
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					found = true
				}
			}
		}
	}

	if !found {
		return false, errors.New(fmt.Sprintf("%s not found", filename))
	}

	return found, nil
}
