package file

import (
	"errors"
	"github.com/choonsiong/golib/stringx"
	"os"
	"testing"
)

func TestIsExecutableInPath(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
		wantErr  error
	}{
		{"go", "go", true, nil},
		{"invalid filename", "abc", false, ErrFileNotFound},
		{"empty filename", "", false, ErrFileNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsExecutableInPath(tt.filename)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("IsExecutableInPath(%v), want error %v; got nil", tt.filename, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("IsExecutableInPath(%v), want error %v; got %v", tt.filename, tt.wantErr, err)
				}
			}

			if got != tt.want {
				t.Errorf("IsExecutableInPath(%v) == %v; want %v", tt.filename, got, tt.want)
			}
		})
	}
}

func TestBinaryMode(t *testing.T) {
	testFile := "/tmp/test"

	_, err := os.Create(testFile) // 0644 (after apply user mask)
	if err != nil {
		t.Fatal("failed to create test file:", testFile)
	}

	tests := []struct {
		name     string
		filename string
		want     string
		wantErr  error
	}{
		{"0644", testFile, "110100100", nil},
		{"invalid filename", "foo.txt", "", ErrFileNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryMode(tt.filename)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("BinaryMode(%v), want error %v; got nil", tt.filename, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("BinaryMode(%v), want error %v; got %v", tt.filename, tt.wantErr, err)
				}
			}

			if got != tt.want {
				t.Errorf("BinaryMode(%v) == %q; want %q", tt.filename, got, tt.want)
			}
		})
	}
}

func TestGetStrings(t *testing.T) {
	var lines string

	for i := 0; i < 10; i++ {
		line, err := stringx.RandomString(10)
		if err != nil {
			t.Fatal("failed to generate random string")
		}
		line += "\n"
		lines += line
	}

	testFile := "/tmp/test"

	f, err := os.Create(testFile) // 0644 (after apply user mask)
	if err != nil {
		t.Fatal("failed to create test file:", testFile)
	}

	f.Write([]byte(lines))
	err = f.Close()
	if err != nil {
		t.Fatal("failed to close test file")
	}

	emptyFile := "/tmp/empty"

	f2, err := os.Create(emptyFile)
	if err != nil {
		t.Fatal("failed to create empty file:", emptyFile)
	}
	err = f2.Close()
	if err != nil {
		t.Fatal("failed to close test file")
	}

	tests := []struct {
		name       string
		filename   string
		ignoreCase bool
		want       int // the length of string slice
		wantErr    error
	}{
		{"valid filename", "/tmp/test", false, 10, nil},
		{"valid filename and ignore case", "/tmp/test", true, 10, nil},
		{"invalid filename", "/tmp/invalid", true, 0, ErrOpenFile},
		{"empty file", "/tmp/empty", true, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSlice, err := GetStrings(tt.filename, tt.ignoreCase)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("GetStrings(%v), want error %v; got nil", tt.filename, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetStrings(%v), want error %v; got %v", tt.filename, tt.wantErr, err)
				}
			}

			got := len(gotSlice)

			if got != tt.want {
				t.Errorf("GetStrings(%v) == %v; want %v", tt.filename, got, tt.want)
			}
		})
	}
}
