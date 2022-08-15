package cli

import (
	"errors"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFilename(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr error
	}{
		{"valid filename", []string{"command", "foo.txt"}, "foo.txt", nil},
		{"empty filename", []string{"command", ""}, "", ErrEmptyFilename},
		{"insufficient arguments", []string{"command"}, "", ErrInsufficientArguments},
		{"too many arguments", []string{"command", "foo.txt", "bar.txt"}, "", ErrTooManyArguments},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filename(tt.args)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Filename(%v), want error %v; got nil", tt.args, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Filename(%v), want error %v; got %v", tt.args, tt.wantErr, err)
				}
			}
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("Filename(%v) == %q; want %q", tt.args, got, tt.want)
			}
		})
	}
}

func TestGetFloat(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr error
	}{
		{"valid float64", "3.141\n", 3.141, nil},
		{"invalid float64", "hello\n", 0, ErrParseFloat},
		{"empty string", "\n", 0, ErrParseFloat},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, err := os.Pipe()
			if err != nil {
				log.Fatal(err)
			}
			origStdin := os.Stdin
			os.Stdin = r

			_, err = w.Write([]byte(tt.input))
			if err != nil {
				w.Close()
				os.Stdin = origStdin
				log.Fatal(err)
			}

			got, err := GetFloat()
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("GetFloat(), want error %v; got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetFloat(), want error %v; got %v", tt.wantErr, err)
				}
			}
			if got != tt.want {
				t.Errorf("GetFloat() == %v; want %v", got, tt.want)
			}

			w.Close()
			os.Stdin = origStdin
		})
	}
}
