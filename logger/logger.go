// Package logger provides the interface to logging facilities.
package logger

// Logger is the interface that wraps the basic logging methods.
type Logger interface {
	PrintDebug(message string, properties map[string]string)
	PrintError(err error, properties map[string]string)
	PrintFatal(err error, properties map[string]string)
	PrintInfo(message string, properties map[string]string)
	PrintWarning(message string, properties map[string]string)
	Write(message []byte) (n int, err error)
}
