package time

import (
	"errors"
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
					t.Errorf("NormalizeHourInTimezone(%q, %q) == nil; want %q", tt.hr, tt.tz, tt.wantErr)
				}

				if !errors.Is(err, tt.wantErr) {
					t.Errorf("NormalizeHourInTimezone(%q, %q) == %q; want %q", tt.hr, tt.tz, err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("NormalizeHourInTimezone(%q, %q) == %q, want %q", tt.hr, tt.tz, got, tt.want)
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
					t.Errorf("GetCalculateTime(%q, %q) == nil; want %q", tt.currentTime, tt.duration, tt.wantErr)
				}

				if !errors.Is(err, tt.wantErr) {
					t.Errorf("GetCalculateTime(%q, %q) == %q; want %q", tt.currentTime, tt.duration, err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("GetCalculateTime(%q, %q) == %q, want %q", tt.currentTime, tt.duration, got, tt.want)
			}
		})
	}
}
