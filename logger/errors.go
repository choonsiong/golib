package logger

import "errors"

var (
	ErrUnknownLogLevel = errors.New("logger: unknown log level")
)
