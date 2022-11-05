package logger

import (
	"fmt"
	"testing"
)

type TestJsonLog struct{}

func (j *TestJsonLog) PrintDebug(message string, properties map[string]string) {
	j.print(message)
}

func (j *TestJsonLog) PrintError(err error, properties map[string]string) {
	if err != nil {
		j.print(err.Error())
	}
}

func (j *TestJsonLog) PrintFatal(err error, properties map[string]string) {
	if err != nil {
		j.print(err.Error())
	}
}

func (j *TestJsonLog) PrintWarning(message string, properties map[string]string) {
	j.print(message)
}

func (j *TestJsonLog) PrintInfo(message string, properties map[string]string) {
	j.print(message)
}

func (j *TestJsonLog) Write(message []byte) (n int, err error) {
	return len(message), nil
}

func (j *TestJsonLog) print(message string) {
	fmt.Println(message)
}

func TestLogger_JsonLog(t *testing.T) {
	f := func(l Logger) {
		l.PrintDebug("debug", nil)
	}

	f(&TestJsonLog{})

	f = func(l Logger) {
		l.PrintError(nil, nil)
	}

	f(&TestJsonLog{})

	f = func(l Logger) {
		l.PrintInfo("info", nil)
	}

	f(&TestJsonLog{})

	f = func(l Logger) {
		l.PrintFatal(nil, nil)
	}

	f(&TestJsonLog{})

	word := "hello"

	g := func(l Logger) (int, error) {
		return l.Write([]byte(word))
	}

	got, err := g(&TestJsonLog{})
	if err != nil {
		t.Fatal(err)
	}

	if got != len(word) {
		t.Errorf("Logger.Write(%v) == %d; want %d", word, got, len(word))
	}
}
