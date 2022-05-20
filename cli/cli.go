package cli

import (
	"errors"
)

var (
	ErrInvalidArgument = errors.New("invalid command-line arguments")
)

// NeedFilename returns the filename from s, s should be pass in as os.Args.
func NeedFilename(s []string) (string, error) {
	if len(s) != 2 {
		return "", ErrInvalidArgument
	}

	return s[1], nil
}
