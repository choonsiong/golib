// Package cli provides helpers to work with command-line.
package cli

import (
	"errors"
)

var (
	ErrInsufficientArguments = errors.New("cli: insufficient command-line arguments")
	ErrInvalidArgument       = errors.New("cli: invalid command-line arguments")
)

// Filename returns the filename from command-line arguments slice s,
// s should be pass in as os.Args.
func Filename(s []string) (string, error) {
	if len(s) != 2 {
		return "", ErrInsufficientArguments
	}

	return s[1], nil
}
