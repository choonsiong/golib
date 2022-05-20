package jsonlog

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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := c.level.String()

			if strings.Compare(got, c.want) != 0 {
				t.Errorf("Level.String() == %q, want %q", got, c.want)
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := LogLevel(c.level)

			if c.wantErr != nil {
				if err == nil {
					t.Errorf("LogLevel(%q), want %q, but got nil", c.level, c.wantErr)
				}
				if !errors.Is(err, c.wantErr) {
					t.Errorf("LogLevel(%q), want %q, but got %q", c.level, c.wantErr, err)
				}
			}

			if got != c.want {
				t.Errorf("LogLevel(%q) == %q, want %q", c.level, got, c.want)
			}
		})
	}
}
