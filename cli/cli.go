// Package cli provides helpers to work with command-line.
package cli

import (
	"errors"
)

var (
	ErrInsufficientArguments = errors.New("cli: insufficient command-line arguments")
	ErrInvalidFilename       = errors.New("cli: invalid filename")
)

// Filename returns the value of args in index 1 as the filename.
// args should be pass in as os.Args.
func Filename(args []string) (string, error) {
	if len(args) != 2 {
		return "", ErrInsufficientArguments
	}

	if args[1] == "" {
		return "", ErrInvalidFilename
	}

	return args[1], nil
}
