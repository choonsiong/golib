// Package file provides helpers to work with files.
package file

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/choonsiong/golib/stringx"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

var (
	AllowedFileTypes  []string
	MaxUploadFileSize int
)

// UploadFiles uploads one or more files to the specified directory with
// random file names. If rename is true, then original file names are use.
func UploadFiles(r *http.Request, uploadDir string, rename ...bool) ([]*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	var uploadedFiles []*UploadedFile

	if MaxUploadFileSize == 0 {
		MaxUploadFileSize = 1024 * 1024 * 1024
	}

	err := r.ParseMultipartForm(int64(MaxUploadFileSize))
	if err != nil {
		return nil, errors.New("upload file size too big")
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fh := range fileHeaders {
			uploadedFiles, err = func(uploadedFiles []*UploadedFile) ([]*UploadedFile, error) {
				var uploadedFile UploadedFile

				f, err := fh.Open()
				if err != nil {
					return nil, err
				}
				defer f.Close()

				buff := make([]byte, 512)
				_, err = f.Read(buff)
				if err != nil {
					return nil, err
				}

				allowed := false
				fileType := http.DetectContentType(buff)

				if len(AllowedFileTypes) > 0 {
					for _, allowedFileType := range AllowedFileTypes {
						if strings.EqualFold(fileType, allowedFileType) {
							allowed = true
						}
					}
				} else {
					allowed = true
				}

				if !allowed {
					return nil, errors.New("upload file type is not supported")
				}

				_, err = f.Seek(0, 0)
				if err != nil {
					return nil, err
				}

				if renameFile {
					uploadedFile.NewFileName = fmt.Sprintf("%s%s", stringx.RandomStringIgnoreError(25), filepath.Ext(fh.Filename))
				} else {
					uploadedFile.NewFileName = fh.Filename
				}

				uploadedFile.OriginalFileName = fh.Filename

				var outputFile *os.File
				defer outputFile.Close()

				if outputFile, err = os.Create(filepath.Join(uploadDir, uploadedFile.NewFileName)); err != nil {
					return nil, err
				} else {
					fileSize, err := io.Copy(outputFile, f)
					if err != nil {
						return nil, err
					}
					uploadedFile.FileSize = fileSize
				}

				uploadedFiles = append(uploadedFiles, &uploadedFile)
				return uploadedFiles, nil
			}(uploadedFiles)
			if err != nil {
				return uploadedFiles, err
			}
		}
	}
	return uploadedFiles, nil
}
