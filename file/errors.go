package file

import "errors"

var (
	ErrFileNotFound    = errors.New("file: file not found")
	ErrInvalidFilename = errors.New("file: invalid filename")
	ErrInvalidTriplet  = errors.New("file: invalid triplet")
)
