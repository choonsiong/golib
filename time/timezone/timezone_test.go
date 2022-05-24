package timezone

import (
	"errors"
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"reflect"
	"testing"
)

func TestNewTimezone(t *testing.T) {
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

func TestTimezone_TimezoneToString(t *testing.T) {
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
			result, err := tz.TimezoneToString(tt.code)
			if tt.wantError != nil {
				if err == nil {
					t.Errorf("want error: %q but got nil", tt.wantError)
				}

				if !errors.Is(err, tt.wantError) {
					t.Errorf("want error: %q; got %q", tt.wantError, err)
				}
			}

			if tt.want != result {
				t.Errorf("want %q; got %q", tt.want, result)
			}
		})
	}
}
