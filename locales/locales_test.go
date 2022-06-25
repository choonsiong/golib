package locales

import (
	"errors"
	"github.com/choonsiong/golib/logger/jsonlog"
	"os"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	want := New(logger)

	got := &Locale{logger}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v; got %v", want, got)
	}

	got = New(nil)
	if got != nil {
		t.Errorf("want nil; got %v", got)
	}
}

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
			got, err := loc.LocaleToString(tt.code)

			if tt.wantError != nil {
				if err == nil {
					t.Errorf("want error %q; got nil", tt.wantError)
				}
				if !errors.Is(err, tt.wantError) {
					t.Errorf("want error %q; got %q", err, tt.wantError)
				}
			}

			if got != tt.want {
				t.Errorf("LocaleToString(%q) == %q; want %q", tt.code, got, tt.want)
			}
		})
	}
}
