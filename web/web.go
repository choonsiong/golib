package web

import (
	"errors"
	"fmt"
	"github.com/choonsiong/golib/stringx"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// UploadedFile is a struct used to save information about an uploaded file.
type UploadedFile struct {
	FileSize         int64
	NewFileName      string
	OriginalFileName string
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
				defer func(f multipart.File) {
					err := f.Close()
					if err != nil {
						log.Fatal(err)
					}
				}(f)

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
				defer func(outputFile *os.File) {
					err := outputFile.Close()
					if err != nil {
						log.Fatal(err)
					}
				}(outputFile)

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

// UploadFile uploads one file to the specified directory with random file
// name. If rename is true, then original file name is use.
func UploadFile(r *http.Request, uploadDir string, rename ...bool) (*UploadedFile, error) {
	renameFile := true
	if len(rename) > 0 {
		renameFile = rename[0]
	}

	files, err := UploadFiles(r, uploadDir, renameFile)
	if err != nil {
		return nil, err
	}

	return files[0], nil
}
