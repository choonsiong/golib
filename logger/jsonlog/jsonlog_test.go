package jsonlog

import (
	"reflect"
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	got := New(nil, LevelDebug)
	want := &Logger{nil, LevelDebug, sync.Mutex{}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Logger.New() == %v, want %v", got, want)
	}
}
