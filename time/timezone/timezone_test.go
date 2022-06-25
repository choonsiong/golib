package timezone

import (
	"errors"
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	tz := New(logger)

	want := &Timezone{
		Logger:    logger,
		Timezones: tz.Timezones,
	}

	if !reflect.DeepEqual(want, tz) {
		t.Errorf("want %v; got %v", want, tz)
	}
}

func TestTimezoneToString(t *testing.T) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	tz := New(logger)

	tests := []struct {
		name      string
		code      string
		want      string
		wantError error
	}{
		{
			name:      "valid timezone",
			code:      "Asia/Tokyo",
			want:      "Tokyo",
			wantError: nil,
		},
		{
			name:      "invalid timezone",
			code:      "Asia/Petaling_Jaya",
			want:      "",
			wantError: ErrNoMatchingTimezone,
		},
		{
			name:      "empty timezone",
			code:      "",
			want:      "",
			wantError: ErrTimezoneIsEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tz.TimezoneToString(tt.code)
			if tt.wantError != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantError)
				}

				if !errors.Is(err, tt.wantError) {
					t.Errorf("want error %q; got %q", tt.wantError, err)
				}
			}

			if tt.want != got {
				t.Errorf("Timezone.TimezoneToString(%q) == %q; want %q", tt.code, got, tt.want)
			}
		})
	}
}
