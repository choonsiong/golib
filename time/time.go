// Package time implements various functions to work with time.
package time

import "errors"

var (
	ErrInvalidHour     = errors.New("time: invalid hour")
	ErrInvalidTimezone = errors.New("time: invalid timezone")
)

// NormalizeHourInTimezone returns the hour hr in the given timezone tz.
// For example:
// UTC 0:00 in +8 = NormalizeHourInTimezone(0, 8) => 8:00
// UTC 20:00 in +8 = NormalizeHourInTimezone(20, 8) => 4:00
// UTC 20:00 in -22 = NormalizeHourInTimezone(20, -22) => 22:00
func NormalizeHourInTimezone(hr int, tz int) (int, error) {
	if hr < 0 || hr > 24 {
		return -1, ErrInvalidHour
	}

	if tz < -24 || tz > 24 {
		return -1, ErrInvalidTimezone
	}

	newHour := hr + tz

	if newHour >= 24 {
		return newHour - 24, nil
	} else if newHour <= 0 {
		return newHour + 24, nil
	}

	return newHour, nil
}
