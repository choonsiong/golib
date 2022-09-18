package web

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
)

func TestUploadFiles(t *testing.T) {
	tests := []struct {
		name          string
		allowedTypes  []string
		renameFile    bool
		errorExpected bool
	}{
		{"allowed no rename", []string{"image/jpeg", "image/png"}, false, false},
		{"allowed rename", []string{"image/jpeg", "image/png"}, true, false},
		{"not allowed", []string{"image/jpeg"}, false, true},
	}

	for _, tt := range tests {
		rd, wd := io.Pipe()
		writer := multipart.NewWriter(wd)

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer writer.Close()
			defer wg.Done()

			part, err := writer.CreateFormFile("file", "./testdata/image.png")
			if err != nil {
				t.Error(err)
			}

			file, err := os.Open("./testdata/image.png")
			if err != nil {
				t.Error(err)
			}

			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				t.Error(err)
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		request := httptest.NewRequest("POST", "/", rd)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		AllowedFileTypes = tt.allowedTypes
		uploadedFiles, err := UploadFiles(request, "./testdata/uploads/", tt.renameFile)
		if err != nil && !tt.errorExpected {
			t.Error(err)
		}

		if !tt.errorExpected {
			if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName)); os.IsNotExist(err) {
				t.Errorf("%s: expected file to exists: %s", tt.name, err.Error())
			}

			_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFiles[0].NewFileName))
		}

		if !tt.errorExpected && err != nil {
			t.Errorf("%s: error expected but nil received", tt.name)
		}

		wg.Wait()
	}
}

func TestUploadFile(t *testing.T) {
	tests := []struct {
		name          string
		allowedTypes  []string
		renameFile    bool
		errorExpected bool
	}{
		{"allowed no rename", []string{"image/jpeg", "image/png"}, false, false},
		{"allowed rename", []string{"image/jpeg", "image/png"}, true, false},
		{"not allowed", []string{"image/jpeg"}, false, true},
	}

	for _, tt := range tests {
		rd, wd := io.Pipe()
		writer := multipart.NewWriter(wd)

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			defer writer.Close()
			defer wg.Done()

			part, err := writer.CreateFormFile("file", "./testdata/image2.png")
			if err != nil {
				t.Error(err)
			}

			file, err := os.Open("./testdata/image2.png")
			if err != nil {
				t.Error(err)
			}

			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				t.Error(err)
			}

			err = png.Encode(part, img)
			if err != nil {
				t.Error(err)
			}
		}()

		request := httptest.NewRequest("POST", "/", rd)
		request.Header.Add("Content-Type", writer.FormDataContentType())

		AllowedFileTypes = tt.allowedTypes
		uploadedFile, err := UploadFile(request, "./testdata/uploads/", tt.renameFile)

		if err != nil && !tt.errorExpected {
			t.Error(err)
		}

		if uploadedFile != nil {
			if _, err := os.Stat(fmt.Sprintf("./testdata/uploads/%s", uploadedFile.NewFileName)); os.IsNotExist(err) {
				t.Errorf("%s: expected file to exists: %s", tt.name, err.Error())
			}

			_ = os.Remove(fmt.Sprintf("./testdata/uploads/%s", uploadedFile.NewFileName))
		}

		wg.Wait()
	}
}

func TestDownloadStaticFile(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	DownloadStaticFile(resp, req, "./testdata", "pic.jpg", "image.jpg")

	data := resp.Result()
	defer data.Body.Close()

	if data.Header["Content-Length"][0] != "98827" { // 98827 bytes
		t.Errorf("DownloadStaticFile(): want %v; got %v", "98827", data.Header["Content-Length"][0])
	}

	if data.Header["Content-Disposition"][0] != "attachment; filename=\"image.jpg\"" {
		t.Errorf("DownloadStaticFile(): want %v; got %v", "image.jpg", data.Header["Content-Disposition"][0])
	}

	_, err := io.ReadAll(data.Body)
	if err != nil {
		t.Fatal(err)
	}
}
