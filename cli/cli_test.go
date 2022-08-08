package cli

import (
	"errors"
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
		{"with filename", []string{"command", "foo.txt"}, "foo.txt", nil},
		{"empty filename", []string{"command", ""}, "", ErrInvalidFilename},
		{"insufficient argument", []string{"command"}, "", ErrInsufficientArguments},
		{"too many arguments", []string{"command", "foo.txt", "bar.txt"}, "", ErrTooManyArguments},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filename(tt.args)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Filename(%v), want error %v; got nil", tt.args, tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("Filename(%v), want error %v; got %v", tt.args, tt.wantErr, err)
				}
			}
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("Filename(%v) == %q; want %q", tt.args, got, tt.want)
			}
		})
	}
}
