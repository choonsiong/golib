package locales

import (
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"testing"
)

func TestLocaleToString(t *testing.T) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	loc := New(logger)

	tests := []struct {
		name      string
		code      string
		want      string
		wantError error
	}{
		{
			name:      "en_US",
			code:      "en_US",
			want:      "English (United States)",
			wantError: nil,
		},
		{
			name:      "en_GB",
			code:      "en_GB",
			want:      "English (United Kingdom)",
			wantError: nil,
		},
		{
			name:      "fr_CH",
			code:      "fr_CH",
			want:      "French (Switzerland)",
			wantError: nil,
		},
		{
			name:      "ms_MY",
			code:      "ms_MY",
			want:      "Malay (Malaysia)",
			wantError: nil,
		},
		{
			name:      "ta_MY",
			code:      "ta_MY",
			want:      "Tamil (Malaysia)",
			wantError: nil,
		},
		{
			name:      "zh_MY",
			code:      "zh_MY",
			want:      "Chinese (Malaysia)",
			wantError: nil,
		},
		{
			name:      "Invalid Code",
			code:      "en_ZZ",
			want:      "",
			wantError: ErrNoMatchingLocale,
		},
		{
			name:      "Empty Code",
			code:      "",
			want:      "",
			wantError: ErrEmptyInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := loc.LocaleToString(tt.code)

			if err != tt.wantError {
				t.Errorf("want error: '%v'; got '%v'", tt.wantError, err)
			}

			if result != tt.want {
				t.Errorf("want '%v'; got '%v'", tt.want, result)
			}
		})
	}
}
