package logger

import (
	"strings"
)

// Level represents the severity level for a log entry.
type Level int8

const (
	LevelDebug   Level = iota // 0
	LevelInfo                 // 1
	LevelWarning              // 2
	LevelError                // 3
	LevelFatal                // 4
	LevelUnknown = -1
)

// String returns a descriptive string for the severity level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogLevel matches the level string to a Level value.
func LogLevel(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warning":
		return LevelWarning
	case "error":
		return LevelError
	case "fatal":
		return LevelFatal
	default:
		return LevelUnknown
	}
}
