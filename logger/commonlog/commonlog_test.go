package commonlog

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
	want := &CommonLog{nil, logger.LevelDebug, sync.Mutex{}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CommonLog.New() == %v; want %v", got, want)
	}
}

func TestCommonLog_PrintDebug(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelDebug)

	logger.PrintDebug("test debug", map[string]string{"key": "value"})

	now := time.Now().UTC().Format(time.RFC3339)
	want := now + ` DEBUG test debug` + "\n" + `key:value` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("CommonLog.PrintDebug() == %q; want %q", got, want)
	}
}

func TestCommonLog_PrintInfo(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelInfo)

	logger.PrintInfo("test info", map[string]string{"key": "value"})

	now := time.Now().UTC().Format(time.RFC3339)
	want := now + ` INFO test info` + "\n" + `key:value` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("CommonLog.PrintInfo() == %q; want %q", got, want)
	}
}

func TestCommonLog_PrintError(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelError)

	logger.PrintError(errors.New("test error"), nil)

	got := buffer.String()
	want := `test error`

	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.PrintError(); %q not found", want)
	}

	want = `ERROR`
	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.PrintError(); %q not found", want)
	}
}

func TestCommonLog_PrintFatal(t *testing.T) {
	t.Skip("not implement")
}

func TestCommonLog_Write(t *testing.T) {
	buffer := new(bytes.Buffer)
	logger := New(buffer, logger.LevelError)

	logger.Write([]byte("test write"))

	got := buffer.String()
	want := `test write`

	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.Write(); %q not found", want)
	}

	want = `ERROR`
	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.Write(); %q not found", want)
	}
}
