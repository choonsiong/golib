package time

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestNormalizeHourInTimezone(t *testing.T) {
	tests := []struct {
		name    string
		hr      int
		tz      int
		want    int
		wantErr error
	}{
		{"positive timezone", 16, 8, 0, nil},
		{"negative timezone", 0, -8, 16, nil},
		{"zero hour and zero timezone", 0, 0, 24, nil},
		{"one hour and zero timezone", 1, 0, 1, nil},
		{"invalid hour", 30, 8, -1, ErrInvalidHour},
		{"invalid timezone", 24, 30, -1, ErrInvalidTimezone},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeHourInTimezone(tt.hr, tt.tz)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("NormalizeHourInTimezone(%v, %v) == nil; want %v", tt.hr, tt.tz, tt.wantErr)
				}

				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NormalizeHourInTimezone(%v, %v) == %v; want %v", tt.hr, tt.tz, err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("NormalizeHourInTimezone(%v, %v) == %v, want %v", tt.hr, tt.tz, got, tt.want)
			}
		})
	}
}

func TestGetTimeNow(t *testing.T) {
	want := time.Now().Format("2006-01-02 15:04:05")
	got := GetTimeNow().Format("2006-01-02 15:04:05")

	if got != want {
		t.Errorf("want %v; but got %v", want, got)
	}
}

func TestGetTimeNowInLocation(t *testing.T) {
	tests := []struct {
		name     string
		location string
		want     time.Time
		wantErr  error
	}{
		{"UTC", "UTC", time.Now().UTC(), nil},
		{"Asia/Tokyo", "Asia/Tokyo", time.Now().UTC().Add(time.Hour * 9), nil},
		{"Europe/London", "Europe/Helsinki", time.Now().UTC().Add(time.Hour * 3), nil},
		{"US/Pacific", "US/Pacific", time.Now().UTC().Add(time.Hour * -7), nil},
		{"Empty location", "", time.Time{}, nil},
		{"Invalid location", "Asia/Foobar", time.Time{}, ErrInvalidLocation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTimeNowInLocation(tt.location)
			gotStr := got.Format("2006-01-02 15:04:05")

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("GetTimeNowInLocation(%q) == nil; want %v", tt.location, tt.wantErr)
				}

				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetTimeNowInLocation(%q) == %v; want %v", tt.location, err, tt.wantErr)
				}
			}

			wantStr := tt.want.Format("2006-01-02 15:04:05")

			if strings.Compare(gotStr, wantStr) != 0 {
				t.Errorf("GetTimeNowInLocation(%q) == %q, want %q", tt.location, gotStr, wantStr)
			}
		})
	}
}

func TestGetCalculateTime(t *testing.T) {
	currentTime := time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)
	tests := []struct {
		name        string
		currentTime time.Time
		duration    string
		want        time.Time
		wantErr     error
	}{
		{"add 1 day", currentTime, "24h", time.Date(2022, 12, 10, 0, 0, 0, 0, time.Local), nil},
		{"minus 1 day", currentTime, "-24h", time.Date(2022, 12, 8, 0, 0, 0, 0, time.Local), nil},
		{"invalid duration unit", currentTime, "1y", time.Time{}, ErrInvalidDuration},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCalculateTime(tt.currentTime, tt.duration)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("GetCalculateTime(%v, %v) == nil; want %v", tt.currentTime, tt.duration, tt.wantErr)
				}

				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetCalculateTime(%v, %v) == %v; want %v", tt.currentTime, tt.duration, err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("GetCalculateTime(%v, %v) == %v, want %v", tt.currentTime, tt.duration, got, tt.want)
			}
		})
	}
}
