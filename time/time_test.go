package time

import (
	"errors"
	"testing"
)

func TestNormalizeHourInTimezone(t *testing.T) {
	cases := []struct {
		name    string
		hr      int
		tz      int
		want    int
		wantErr error
	}{
		{"positive timezone", 16, 8, 0, nil},
		{"negative timezone", 0, -8, 16, nil},
		{"zero hour and zero timezone", 0, 0, 24, nil},
		{"invalid hour", 30, 8, -1, ErrInvalidHour},
		{"invalid timezone", 24, 30, -1, ErrInvalidTimezone},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := NormalizeHourInTimezone(c.hr, c.tz)

			if c.wantErr != nil {
				if err == nil {
					t.Errorf("NormalizeHourInTimezone(%q, %q), want %q, but got nil", c.hr, c.tz, c.wantErr)
				}

				if !errors.Is(err, c.wantErr) {
					t.Errorf("NormalizeHourInTimezone(%q, %q), want %q, but got %q", c.hr, c.tz, c.wantErr, got)
				}
			}

			if got != c.want {
				t.Errorf("NormalizeHourInTimezone(%q, %q) == %q, want %q", c.hr, c.tz, got, c.want)
			}
		})
	}
}
