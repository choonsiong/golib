package stringx

import "errors"

var (
	ErrGenerateRandomString = errors.New("stringx: failed to generate random string")
	ErrInvalidInput         = errors.New("stringx: invalid input")
)
