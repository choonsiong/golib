package jsonlog

import (
	"fmt"
	"strings"
)

// Level represents the severity level for a log entry.
type Level int8

const (
	LevelDebug   Level = iota // 0
	LevelInfo                 // 1
	LevelError                // 2
	LevelFatal                // 3
	LevelOff     = 99
	LevelInvalid = -1
)

// String returns a human-friendly string for the severity level.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelOff:
		return "OFF"
	case LevelInvalid:
		return "INVALID"
	default:
		return "UNKNOWN LEVEL"
	}
}

// LogLevel matches the level string to a Level type.
func LogLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "debug":
		return LevelDebug, nil
	case "info":
		return LevelInfo, nil
	case "error":
		return LevelError, nil
	case "fatal":
		return LevelFatal, nil
	case "off":
		return LevelOff, nil
	default:
		return LevelInvalid, fmt.Errorf("invalid log level")
	}
}
