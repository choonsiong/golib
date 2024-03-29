package stringx

import "errors"

var (
	ErrEmptySlug            = errors.New("stringx: empty slug")
	ErrEmptyString          = errors.New("stringx: empty string")
	ErrGenerateRandomString = errors.New("stringx: failed to generate random string")
	ErrInvalidLength        = errors.New("stringx: invalid length")
	ErrInvalidInput         = errors.New("stringx: invalid input")
)
