// Package colorlog implements color logging.
package colorlog

import (
	"github.com/choonsiong/golib/v2/logger"
	"github.com/fatih/color"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// ColorLog is custom logger type. It holds the output destination that the log
// entries will be written to, the minimum severity level that log entries will
// be written for, plus a mutex for coordinating the writes.
type ColorLog struct {
	out      io.Writer
	minLevel logger.Level
	mu       sync.Mutex
}

var (
	debugColor   = color.New(color.FgBlue).SprintfFunc()
	errorColor   = color.New(color.FgRed).SprintfFunc()
	fatalColor   = color.New(color.FgHiRed).SprintfFunc()
	infoColor    = color.New(color.FgGreen).SprintfFunc()
	warningColor = color.New(color.FgYellow).SprintfFunc()
)

// New returns a new ColorLog instance which writes log entries at or above a
// minimum severity level to a specific output destination.
func New(out io.Writer, minLevel logger.Level) *ColorLog {
	return &ColorLog{
		out:      out,
		minLevel: minLevel,
	}
}

// PrintDebug is a helper method to write DEBUG level log entries.
func (l *ColorLog) PrintDebug(message string, properties map[string]string) {
	l.print(logger.LevelDebug, message, properties)
}

// PrintInfo is a helper method to write INFO level log entries.
func (l *ColorLog) PrintInfo(message string, properties map[string]string) {
	l.print(logger.LevelInfo, message, properties)
}

// PrintWarning is a helper method to write WARNING level log entries.
func (l *ColorLog) PrintWarning(message string, properties map[string]string) {
	l.print(logger.LevelWarning, message, properties)
}

// PrintError is a helper method to write ERROR level log entries.
func (l *ColorLog) PrintError(err error, properties map[string]string) {
	l.print(logger.LevelError, err.Error(), properties)
}

// PrintFatal is a helper method to write FATAL level log entries.
func (l *ColorLog) PrintFatal(err error, properties map[string]string) {
	l.print(logger.LevelFatal, err.Error(), properties)
	os.Exit(1) // terminate application also
}

// Print is a private method for writing the log entry.
func (l *ColorLog) print(level logger.Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	// An anonymous struct holding the data for the log entry.
	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	if level >= logger.LevelError {
		aux.Trace = string(debug.Stack())
	}

	// To hold the actual log entry text.
	var line string
	line += aux.Time
	line += " "
	line += aux.Level
	line += " "
	line += aux.Message
	if len(aux.Properties) != 0 {
		for k, v := range aux.Properties {
			line += "\n\t"
			line += k
			line += ": "
			line += v
		}
	}
	if aux.Trace != "" {
		line += "\n"
		line += aux.Trace
	}
	line += "\n"

	var colorLine string

	switch level {
	case logger.LevelDebug:
		colorLine = debugColor(line)
	case logger.LevelError:
		colorLine = errorColor(line)
	case logger.LevelFatal:
		colorLine = fatalColor(line)
	case logger.LevelInfo:
		colorLine = infoColor(line)
	case logger.LevelWarning:
		colorLine = warningColor(line)
	}

	// Lock the mutex so that no two writes to the output destination can
	// happen concurrently.
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write the log entry (JSON encoded) followed by a newline.
	//return l.out.Write([]byte(line))
	return l.out.Write([]byte(colorLine))
}

// Write writes a log entry at the ERROR level with no additional properties.
// It is implemented to satisfy the io.Writer interface.
func (l *ColorLog) Write(message []byte) (n int, err error) {
	return l.print(logger.LevelError, string(message), nil)
}
