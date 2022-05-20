package jsonlog

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	got := New(nil, LevelDebug)
	want := &Logger{nil, LevelDebug, sync.Mutex{}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Logger.New() == %v, want %v", got, want)
	}
}

func TestLogger_PrintDebug(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, LevelDebug)

	logger.PrintDebug("test debug", nil)

	now := time.Now().UTC().Format(time.RFC3339)
	want := `{"level":"DEBUG","time":"` + now + `","message":"test debug"}` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("Logger.PrintDebug() == %q, want %q", got, want)
	}
}

func TestLogger_PrintInfo(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, LevelInfo)

	logger.PrintInfo("test info", nil)

	now := time.Now().UTC().Format(time.RFC3339)
	want := `{"level":"INFO","time":"` + now + `","message":"test info"}` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("Logger.PrintInfo() == %q, want %q", got, want)
	}
}

func TestLogger_PrintError(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, LevelError)

	logger.PrintError(errors.New("test error"), nil)

	got := buffer.String()
	want := "\"message\":\"test error\""

	if !strings.Contains(got, want) {
		t.Errorf("Logger.PrintError(), %q not found", want)
	}

	want = "\"level\":\"ERROR\""
	if !strings.Contains(got, want) {
		t.Errorf("Logger.PrintError(), %q not found", want)
	}
}

func TestLogger_PrintFatal(t *testing.T) {
	t.Skip("not implement")
}

func TestLogger_Write(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, LevelError)

	logger.Write([]byte("test write"))

	got := buffer.String()
	want := "\"message\":\"test write\""

	if !strings.Contains(got, want) {
		t.Errorf("Logger.Write(), %q not found", want)
	}

	want = "\"level\":\"ERROR\""
	if !strings.Contains(got, want) {
		t.Errorf("Logger.Write(), %q not found", want)
	}
}
