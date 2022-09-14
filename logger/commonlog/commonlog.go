// Package commonlog implements common logging.
package commonlog

import (
	"github.com/choonsiong/golib/logger"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// CommonLog is custom logger type. It holds the output destination that the log
// entries will be written to, the minimum severity level that log entries will
// be written for, plus a mutex for coordinating the writes.
type CommonLog struct {
	out      io.Writer
	minLevel logger.Level
	mu       sync.Mutex
}

// New returns a new CommonLog instance which writes log entries at or above a
// minimum severity level to a specific output destination.
func New(out io.Writer, minLevel logger.Level) *CommonLog {
	return &CommonLog{
		out:      out,
		minLevel: minLevel,
	}
}

// PrintDebug is a helper method to write DEBUG level log entries.
func (l *CommonLog) PrintDebug(message string, properties map[string]string) {
	l.print(logger.LevelDebug, message, properties)
}

// PrintInfo is a helper method to write INFO level log entries.
func (l *CommonLog) PrintInfo(message string, properties map[string]string) {
	l.print(logger.LevelInfo, message, properties)
}

// PrintError is a helper method to write ERROR level log entries.
func (l *CommonLog) PrintError(err error, properties map[string]string) {
	l.print(logger.LevelError, err.Error(), properties)
}

// PrintFatal is a helper method to write FATAL level log entries.
func (l *CommonLog) PrintFatal(err error, properties map[string]string) {
	l.print(logger.LevelFatal, err.Error(), properties)
	os.Exit(1) // terminate application also
}

// Print is a private method for writing the log entry.
func (l *CommonLog) print(level logger.Level, message string, properties map[string]string) (int, error) {
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

	// Lock the mutex so that no two writes to the output destination can
	// happen concurrently.
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write the log entry (JSON encoded) followed by a newline.
	return l.out.Write([]byte(line))
}

// Write writes a log entry at the ERROR level with no additional properties.
// It is implemented to satisfy the io.Writer interface.
func (l *CommonLog) Write(message []byte) (n int, err error) {
	return l.print(logger.LevelError, string(message), nil)
}
