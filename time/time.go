// Package time provides helpers to work with time.
package time

import (
	"fmt"
	"github.com/choonsiong/golib/v2/time/timezone"
	"time"
)

// GetCalculateTime returns new time with added duration to current time.
func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, fmt.Errorf("GetCalculateTime(%v, %v): %w", currentTime, d, ErrInvalidDuration)
	}
	return currentTime.Add(duration), nil
}

// GetTimeInTimezone returns the time in given timezone.
func GetTimeInTimezone(t time.Time, tz string) (time.Time, error) {
	if !timezone.HasTimezone(tz) {
		return time.Time{}, ErrInvalidTimezone
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		return time.Time{}, ErrInvalidLocation
	}

	return t.In(loc), nil
}

// GetTimeNow returns the current local time.
func GetTimeNow() time.Time {
	return time.Now()
}

// GetTimeNowInLocation returns the current time in location.
func GetTimeNowInLocation(loc string) (time.Time, error) {
	var location *time.Location
	var err error

	if loc == "" {
		return time.Time{}, fmt.Errorf("GetTimeNowInLocation(%q): %w", loc, ErrInvalidLocation)
		//location, err = time.LoadLocation("UTC")
		//if err != nil {
		//	return time.Time{}, ErrInvalidLocation
		//}
	} else {
		location, err = time.LoadLocation(loc)
		if err != nil {
			return time.Time{}, fmt.Errorf("GetTimeNowInLocation(%q): %w", loc, ErrInvalidLocation)
		}
	}

	return time.Now().In(location), nil
}

// NormalizeHourInTimezone returns the hour hr in the given timezone tz.
// For example:
// UTC 0:00 in +8 = NormalizeHourInTimezone(0, 8) => 8:00
// UTC 20:00 in +8 = NormalizeHourInTimezone(20, 8) => 4:00
// UTC 20:00 in -22 = NormalizeHourInTimezone(20, -22) => 22:00
func NormalizeHourInTimezone(hr int, tz int) (int, error) {
	if hr < 0 || hr > 24 {
		return -1, fmt.Errorf("NormalizeHourInTimezone(%d, %d): %w", hr, tz, ErrInvalidHour)
	}

	if tz < -24 || tz > 24 {
		return -1, fmt.Errorf("NormalizeHourInTimezone(%d, %d): %w", hr, tz, ErrInvalidTimezone)
	}

	newHour := hr + tz

	if newHour >= 24 {
		return newHour - 24, nil
	} else if newHour <= 0 {
		return newHour + 24, nil
	}

	return newHour, nil
}
