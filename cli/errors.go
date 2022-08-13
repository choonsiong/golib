package cli

import "errors"

var (
	ErrEmptyFilename         = errors.New("cli: empty filename")
	ErrInsufficientArguments = errors.New("cli: insufficient command-line arguments")
	ErrInvalidFilename       = errors.New("cli: invalid filename")
	ErrTooManyArguments      = errors.New("cli: too many arguments")
)
