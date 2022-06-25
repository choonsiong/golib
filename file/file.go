// Package file provides helpers to work with files.
package file

import (
	"errors"
	"os"
	"strings"
)

const (
	UserPath = "PATH"
)

var (
	ErrInvalidTriplet  = errors.New("file: invalid triplet")
	ErrInvalidFilename = errors.New("file: filename not found")
)

// IsExecutableInPath returns true if the filename is an executable and
// exists in the user PATH.
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
		return false, ErrInvalidFilename
	}
	return found, nil
}

// BinaryMode returns the file mode of filename in binary mode.
func BinaryMode(filename string) (string, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return "", ErrInvalidFilename
	}
	fileMode := fileInfo.Mode()
	return convertToBinary(fileMode.String())
}

// convertToBinary returns the permissions given in binary form.
func convertToBinary(permissions string) (string, error) {
	if permissions == "" {
		return "", ErrInvalidTriplet
	}

	binaryPermissions := permissions[1:]
	p1, err := tripletToBinary(binaryPermissions[0:3])
	if err != nil {
		return "", err
	}
	p2, err := tripletToBinary(binaryPermissions[3:6])
	if err != nil {
		return "", err
	}
	p3, err := tripletToBinary(binaryPermissions[6:9])
	if err != nil {
		return "", err
	}
	return p1 + p2 + p3, nil
}

// tripletToBinary returns the single triplet in three digits binary value.
func tripletToBinary(triplet string) (string, error) {
	if triplet == "rwx" {
		return "111", nil
	}
	if triplet == "-wx" {
		return "011", nil
	}
	if triplet == "--x" {
		return "001", nil
	}
	if triplet == "---" {
		return "000", nil
	}
	if triplet == "r-x" {
		return "101", nil
	}
	if triplet == "r--" {
		return "100", nil
	}
	if triplet == "rw-" {
		return "110", nil
	}
	if triplet == "-w-" {
		return "010", nil
	}
	return "", ErrInvalidTriplet
}
