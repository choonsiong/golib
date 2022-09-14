package jsonlog

import (
	"bytes"
	"errors"
	"github.com/choonsiong/golib/logger"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	got := New(nil, logger.LevelDebug)
	want := &JSONLog{nil, logger.LevelDebug, sync.Mutex{}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("JSONLog.New() == %v; want %v", got, want)
	}
}

func TestJSONLog_PrintDebug(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelDebug)

	logger.PrintDebug("test debug", nil)

	now := time.Now().UTC().Format(time.RFC3339)
	want := `{"level":"DEBUG","time":"` + now + `","message":"test debug"}` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("JSONLog.PrintDebug() == %q; want %q", got, want)
	}
}

func TestJSONLog_PrintInfo(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelInfo)

	logger.PrintInfo("test info", nil)

	now := time.Now().UTC().Format(time.RFC3339)
	want := `{"level":"INFO","time":"` + now + `","message":"test info"}` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("JSONLog.PrintInfo() == %q; want %q", got, want)
	}
}

func TestJSONLog_PrintError(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelError)

	logger.PrintError(errors.New("test error"), nil)

	got := buffer.String()
	want := "\"message\":\"test error\""

	if !strings.Contains(got, want) {
		t.Errorf("JSONLog.PrintError(); %q not found", want)
	}

	want = "\"level\":\"ERROR\""
	if !strings.Contains(got, want) {
		t.Errorf("JSONLog.PrintError(); %q not found", want)
	}
}

func TestJSONLog_PrintFatal(t *testing.T) {
	t.Skip("not implement")
}

func TestJSONLog_Write(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelError)

	logger.Write([]byte("test write"))

	got := buffer.String()
	want := "\"message\":\"test write\""

	if !strings.Contains(got, want) {
		t.Errorf("JSONLog.Write(); %q not found", want)
	}

	want = "\"level\":\"ERROR\""
	if !strings.Contains(got, want) {
		t.Errorf("JSONLog.Write(); %q not found", want)
	}
}
