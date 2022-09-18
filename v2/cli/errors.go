package cli

import "errors"

var (
	ErrEmptyFilename         = errors.New("cli: empty filename")
	ErrInsufficientArguments = errors.New("cli: insufficient command-line arguments")
	ErrInvalidFilename       = errors.New("cli: invalid filename")
	ErrParseFloat            = errors.New("cli: failed to parse float")
	ErrReadString            = errors.New("cli: failed to read string")
	ErrTooManyArguments      = errors.New("cli: too many arguments")
)
