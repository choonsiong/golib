package time

import "errors"

var (
	ErrInvalidDuration = errors.New("time: invalid duration")
	ErrInvalidHour     = errors.New("time: invalid hour")
	ErrInvalidLocation = errors.New("time: invalid location")
	ErrInvalidTimezone = errors.New("time: invalid timezone")
)
