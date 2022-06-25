package cli

import (
	"errors"
	"strings"
	"testing"
)

func TestFilename(t *testing.T) {
	cases := []struct {
		name    string
		args    []string
		want    string
		wantErr error
	}{
		{"with filename", []string{"command", "foo.txt"}, "foo.txt", nil},
		{"without filename", []string{"command"}, "", ErrInsufficientArguments},
		{"empty filename", []string{"command", ""}, "", nil},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Filename(c.args)

			if c.wantErr != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", c.wantErr)
				}
				if !errors.Is(err, c.wantErr) {
					t.Errorf("want error %q; got %q", c.wantErr, err)
				}
			}

			if strings.Compare(got, c.want) != 0 {
				t.Errorf("want %q; got %q", c.want, got)
			}
		})
	}
}
