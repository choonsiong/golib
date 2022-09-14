package logger

import (
	"errors"
	"strings"
	"testing"
)

func TestLevel_String(t *testing.T) {
	cases := []struct {
		name  string
		level Level
		want  string
	}{
		{"debug", LevelDebug, "DEBUG"},
		{"info", LevelInfo, "INFO"},
		{"error", LevelError, "ERROR"},
		{"fatal", LevelFatal, "FATAL"},
		{"off", LevelOff, "OFF"},
		{"invalid", LevelInvalid, "INVALID"},
		{"unknown", Level(42), "UNKNOWN LEVEL"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.level.String()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("Level.String() == %q; want %q", got, tt.want)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	cases := []struct {
		name    string
		level   string
		want    Level
		wantErr error
	}{
		{"debug", "debug", LevelDebug, nil},
		{"info", "info", LevelInfo, nil},
		{"error", "error", LevelError, nil},
		{"fatal", "fatal", LevelFatal, nil},
		{"off", "off", LevelOff, nil},
		{"invalid", "invalid", LevelInvalid, nil},
		{"unknown", "unknown", LevelInvalid, ErrUnknownLogLevel},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogLevel(tt.level)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("LogLevel(%q) == nil; want %q", tt.level, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("LogLevel(%q) == %q; want %q", tt.level, err, tt.wantErr)
				}
			}

			if got != tt.want {
				t.Errorf("LogLevel(%q) == %q; want %q", tt.level, got, tt.want)
			}
		})
	}
}
