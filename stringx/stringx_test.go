package stringx

import (
	"errors"
	"strings"
	"testing"
)

func TestCapitalizeEachWord(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"one word", "hello", "Hello"},
		{"two words", "hello world", "Hello World"},
		{"special characters", "i'm foo bar", "I'm Foo Bar"},
		{"empty", "", ""},
		{"chinese characters", "你好吗", "你好吗"},
		{"japanese characters", "ありがとうございました", "ありがとうございました"},
		{"special character prefix", "/hello world", "/hello world"},
		{"start with a", "a", "A"},
		{"start with A", "A", "A"},
		{"start with z", "z", "Z"},
		{"start with Z", "Z", "Z"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CapitalizeEachWord(tt.input)
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("CapitalizeEachWord(%q) == %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestRandomString(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		want    int
		wantErr error
	}{
		{"length -1", -1, 0, ErrInvalidInput},
		{"length 0", 0, 0, nil},
		{"length 10", 10, 10, nil},
		{"length 20", 20, 20, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := RandomString(tt.length)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("RandomString(%v), want error %v; got nil", tt.length, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("RandomString(%v), want error %v; got %v", tt.length, tt.wantErr, err)
				}
			}

			got := len(s)
			if got != tt.want {
				t.Errorf("RandomString(%d)'s length == %d; want %d", tt.length, got, tt.want)
			}
		})
	}
}
