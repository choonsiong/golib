package file

import (
	"errors"
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
		{"invalid filename", "abc", false, ErrInvalidFilename},
		{"empty filename", "", false, ErrInvalidFilename},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsExecutableInPath(tt.filename)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if tt.want != got {
				t.Errorf("want %v; got %v", tt.want, got)
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
		{"invalid filename", "foo.txt", "", ErrInvalidFilename},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryMode(tt.filename)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if tt.want != got {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}

func Test_convertToBinary(t *testing.T) {
	tests := []struct {
		name        string
		permissions string
		want        string
		wantErr     error
	}{
		{"-rwxrwxrwx", "-rwxrwxrwx", "111111111", nil},
		{"-abcrwxrwx", "-abcrwxrwx", "", ErrInvalidTriplet},
		{"-rwxabcrwx", "-rwxabcrwx", "", ErrInvalidTriplet},
		{"-rwxrwxabc", "-rwxrwxabc", "", ErrInvalidTriplet},
		{"empty", "", "", ErrInvalidTriplet},
		{"all invalid", "aaaaaaaaaa", "", ErrInvalidTriplet},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToBinary(tt.permissions)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if tt.want != got {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}

func Test_tripletToBinary(t *testing.T) {
	tests := []struct {
		name    string
		triplet string
		want    string
		wantErr error
	}{
		{"rwx", "rwx", "111", nil},
		{"-wx", "-wx", "011", nil},
		{"--x", "--x", "001", nil},
		{"---", "---", "000", nil},
		{"r-x", "r-x", "101", nil},
		{"r--", "r--", "100", nil},
		{"rw-", "rw-", "110", nil},
		{"-w-", "-w-", "010", nil},
		{"empty", "", "", ErrInvalidTriplet},
		{"invalid triplet", "abc", "", ErrInvalidTriplet},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tripletToBinary(tt.triplet)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if tt.want != got {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}
