package logger

import (
	"strings"
	"testing"
)

func TestLevel_String(t *testing.T) {
	tests := []struct {
		name  string
		level Level
		want  string
	}{
		{"debug", LevelDebug, "DEBUG"},
		{"info", LevelInfo, "INFO"},
		{"error", LevelError, "ERROR"},
		{"fatal", LevelFatal, "FATAL"},
		{"unknown", Level(42), "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.level.String()

			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("Level.String() == %q; want %q", got, tt.want)
			}
		})
	}
}

func TestLogLevel(t *testing.T) {
	tests := []struct {
		name  string
		level string
		want  Level
	}{
		{"debug", "debug", LevelDebug},
		{"info", "info", LevelInfo},
		{"error", "error", LevelError},
		{"fatal", "fatal", LevelFatal},
		{"unknown", "unknown", LevelUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LogLevel(tt.level)

			if got != tt.want {
				t.Errorf("LogLevel(%q) == %q; want %q", tt.level, got, tt.want)
			}
		})
	}
}
