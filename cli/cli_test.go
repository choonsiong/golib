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
		{"without filename", []string{"command"}, "", ErrInsufficientArguments},
		{"empty filename", []string{"command", ""}, "", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filename(tt.args)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("want error %q; got %q", tt.wantErr, err)
				}
			}

			if strings.Compare(tt.want, got) != 0 {
				t.Errorf("want %q; got %q", tt.want, got)
			}
		})
	}
}
