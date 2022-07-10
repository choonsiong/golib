// Package cli provides helpers to work with command-line.
package cli

import (
	"errors"
)

var (
	ErrInsufficientArguments = errors.New("cli: insufficient command-line arguments")
	ErrInvalidArguments      = errors.New("cli: invalid command-line arguments")
)

// Filename returns the value of args in index 1 as the filename.
// args should be pass in as os.Args.
func Filename(args []string) (string, error) {
	if len(args) < 2 {
		return "", ErrInsufficientArguments
	}
	if len(args) != 2 {
		return "", ErrInvalidArguments
	}

	return args[1], nil
}
