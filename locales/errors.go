package locales

import "errors"

var (
	ErrEmptyInput       = errors.New("locales: input is empty")
	ErrNoMatchingLocale = errors.New("locales: no matching locale found")
)
