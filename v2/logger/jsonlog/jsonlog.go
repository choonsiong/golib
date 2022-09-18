// Package jsonlog implements json format logging.
package jsonlog

import (
	"encoding/json"
	"github.com/choonsiong/golib/v2/logger"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// JSONLog is custom logger type. It holds the output destination that the log
// entries will be written to, the minimum severity level that log entries will
// be written for, plus a mutex for coordinating the writes.
type JSONLog struct {
	out      io.Writer
	minLevel logger.Level
	mu       sync.Mutex
}

// New returns a new JSONLog instance which writes log entries at or above a
// minimum severity level to a specific output destination.
func New(out io.Writer, minLevel logger.Level) *JSONLog {
	return &JSONLog{
		out:      out,
		minLevel: minLevel,
	}
}

// PrintDebug is a helper method to write DEBUG level log entries.
func (l *JSONLog) PrintDebug(message string, properties map[string]string) {
	l.print(logger.LevelDebug, message, properties)
}

// PrintInfo is a helper method to write INFO level log entries.
func (l *JSONLog) PrintInfo(message string, properties map[string]string) {
	l.print(logger.LevelInfo, message, properties)
}

// PrintError is a helper method to write ERROR level log entries.
func (l *JSONLog) PrintError(err error, properties map[string]string) {
	l.print(logger.LevelError, err.Error(), properties)
}

// PrintFatal is a helper method to write FATAL level log entries.
func (l *JSONLog) PrintFatal(err error, properties map[string]string) {
	l.print(logger.LevelFatal, err.Error(), properties)
	os.Exit(1) // terminate application also
}

// Print is a private method for writing the log entry.
func (l *JSONLog) print(level logger.Level, message string, properties map[string]string) (int, error) {
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
		// Print it to console (it is a bit difficult to read the trace in
		// JSON format :p
		//log.Println(string(debug.Stack()))
	}

	// To hold the actual log entry text.
	var line []byte

	line, err := json.Marshal(aux)
	if err != nil {
		line = []byte(logger.LevelError.String() + ": unable to marshal log messages: " + err.Error())
	}

	// Lock the mutex so that no two writes to the output destination can
	// happen concurrently.
	l.mu.Lock()
	defer l.mu.Unlock()

	// Write the log entry (JSON encoded) followed by a newline.
	return l.out.Write(append(line, '\n'))
}

// Write writes a log entry at the ERROR level with no additional properties.
// It is implemented to satisfy the io.Writer interface.
func (l *JSONLog) Write(message []byte) (n int, err error) {
	return l.print(logger.LevelError, string(message), nil)
}
