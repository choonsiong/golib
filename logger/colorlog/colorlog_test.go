package colorlog

import (
	"bytes"
	"fmt"
	"github.com/choonsiong/golib/v2/logger"
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	got := New(nil, logger.LevelDebug)
	want := &ColorLog{nil, logger.LevelDebug, sync.Mutex{}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("CommonLog.New() == %v; want %v", got, want)
	}
}

func TestColorLog_PrintDebug(t *testing.T) {
	buffer := new(bytes.Buffer)
	l := New(buffer, logger.LevelDebug)

	l.PrintDebug("test debug", map[string]string{"key": "value"})

	now := time.Now().UTC().Format(time.RFC3339)
	want := now + ` DEBUG test debug` + "\n\t" + `key: value` + "\n"
	got := buffer.String()

	fmt.Println(got)

	if strings.Compare(got, want) != 0 {
		t.Errorf("ColorLog.PrintDebug() == %q; want %q", got, want)
	}
}
