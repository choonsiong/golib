package form

import "testing"

func TestErrors_Get(t *testing.T) {
	e := errors{}

	got := e.Get("test")
	if got != "" {
		t.Errorf("got %q; want empty string", got)
	}

	e.Add("test", "testing")
	got = e.Get("test")
	if got == "" {
		t.Errorf("got empty string; want %q", "testing")
	}
}
