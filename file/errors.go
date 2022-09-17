package file

import "errors"

var (
	ErrCloseFile       = errors.New("file: file close error")
	ErrCreateDir       = errors.New("file: failed to create directory")
	ErrFileNotFound    = errors.New("file: file not found")
	ErrInvalidFilename = errors.New("file: invalid filename")
	ErrInvalidTriplet  = errors.New("file: invalid triplet")
	ErrOpenFile        = errors.New("file: failed to open file")
	ErrScanFile        = errors.New("file: file scan error")
)
