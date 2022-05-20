package cli

import (
	"errors"
	"strings"
	"testing"
)

func TestNeedFilename_ValidArgument(t *testing.T) {
	args := []string{"command", "foo.txt"}

	want := "foo.txt"

	filename, err := NeedFilename(args)
	if err != nil {
		if !errors.Is(err, ErrInvalidArgument) {
			t.Fatalf("NeedFileName(), want %q, but got %q", ErrInvalidArgument, err)
		}
	}

	if strings.Compare(want, filename) != 0 {
		t.Errorf("NeedFilename(), want %q, but got %q", want, filename)
	}
}
