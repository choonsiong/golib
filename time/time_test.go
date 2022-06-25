package time

import (
	"errors"
	"testing"
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
