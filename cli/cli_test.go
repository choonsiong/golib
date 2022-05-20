package cli

import (
	"errors"
	"strings"
	"testing"
)

func TestNeedFilename(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		want    string
		wantErr error
	}{
		{"valid argument", []string{"command", "foo.txt"}, "foo.txt", nil},
		{"invalid argument", []string{"command"}, "", ErrInvalidArgument},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := NeedFilename(c.args)

			if c.wantErr != nil {
				if err == nil {
					t.Errorf("NeedFilename(), want %q, but got nil", c.wantErr)
				}
				if !errors.Is(err, c.wantErr) {
					t.Errorf("NeedFilename(), want %q, but got %q", c.wantErr, err)
				}
			}

			if strings.Compare(got, c.want) != 0 {
				t.Errorf("NeedFilename() == %q, want %q", got, c.want)
			}
		})
	}
}
