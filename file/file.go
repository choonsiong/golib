// Package file provides helpers to work with files.
package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	UserPath = "PATH"
)

// BinaryMode returns the file mode of filename in binary digits.
func BinaryMode(filename string) (string, error) {
	fileInfo, err := os.Stat(filename)

	if err != nil {
		return "", fmt.Errorf("%w: %q", ErrFileNotFound, filename)
	}

	fileMode := fileInfo.Mode()

	return convertToBinary(fileMode.String())
}

// IsExecutableInPath returns true if filename is an executable and exists
// in the user PATH.
func IsExecutableInPath(filename string) (bool, error) {
	found := false
	path := os.Getenv(UserPath)
	pathSlice := strings.Split(path, ":")

	for _, dir := range pathSlice {
		fullPath := dir + "/" + filename
		fileInfo, err := os.Stat(fullPath)

		if err == nil { // file found in user path
			mode := fileInfo.Mode()

			if mode.IsRegular() {
				if mode&0111 != 0 { // check executable bits
					found = true
				}
			}
		}
	}

	if !found {
		return false, fmt.Errorf("%w: %q", ErrFileNotFound, filename)
	}

	return found, nil
}

// GetStrings reads all lines from filename and returns a slice of string.
func GetStrings(filename string, ignoreCase bool) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOpenFile, err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if ignoreCase {
			line = strings.ToLower(line)
		}

		lines = append(lines, line)
	}

	err = file.Close()

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCloseFile, err)
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("%w: %v", ErrScanFile, scanner.Err())
	}

	return lines, nil
}
