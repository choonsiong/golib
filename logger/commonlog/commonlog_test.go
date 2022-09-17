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
	l := New(buffer, logger.LevelDebug)

	l.PrintDebug("test debug", map[string]string{"key": "value"})

	now := time.Now().UTC().Format(time.RFC3339)
	want := now + ` DEBUG test debug` + "\n\t" + `key: value` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("CommonLog.PrintDebug() == %q; want %q", got, want)
	}
}

func TestCommonLog_PrintInfo(t *testing.T) {
	buffer := new(bytes.Buffer)
	l := New(buffer, logger.LevelInfo)

	l.PrintInfo("test info", map[string]string{"key": "value"})

	now := time.Now().UTC().Format(time.RFC3339)
	want := now + ` INFO test info` + "\n\t" + `key: value` + "\n"
	got := buffer.String()

	if strings.Compare(got, want) != 0 {
		t.Errorf("CommonLog.PrintInfo() == %q; want %q", got, want)
	}
}

func TestCommonLog_PrintError(t *testing.T) {
	buffer := new(bytes.Buffer)
	l := New(buffer, logger.LevelError)

	l.PrintError(errors.New("test error"), nil)

	got := buffer.String()
	want := `test error`

	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.PrintError() == %q; want %q", got, want)
	}

	want = `ERROR`
	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.PrintError() == %q; want %q", got, want)
	}
}

func TestCommonLog_PrintFatal(t *testing.T) {
	t.Skip("not implement")
}

func TestCommonLog_Write(t *testing.T) {
	buffer := new(bytes.Buffer)
	l := New(buffer, logger.LevelError)

	_, err := l.Write([]byte("test write"))
	if err != nil {
		t.Fatal(err)
	}

	got := buffer.String()
	want := `test write`

	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.Write() == %q;  want %q", got, want)
	}

	want = `ERROR`
	if !strings.Contains(got, want) {
		t.Errorf("CommonLog.Write() == %q;  want %q", got, want)
	}
}

func TestCommonLog_print(t *testing.T) {
	buffer := new(bytes.Buffer)
	l := New(buffer, logger.LevelDebug)
	l.minLevel = 1
	want := 0

	got, err := l.print(-1, "", nil)
	if err != nil {
		t.Errorf("CommonLog.print() == %v; want nil", err)
	}

	if got != want {
		t.Errorf("CommonLog.print(%v, %v, %v) == %d; want %d", -1, "", "", got, want)
	}
}
