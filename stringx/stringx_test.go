// Package stringx provides extra functions to work with string.
package stringx

import (
	"strings"
	"testing"
)

func TestToCamelCase(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToCamelCase(tt.input)
			if strings.Compare(got, tt.want) != 0 {
				t.Errorf("ToCamelCase(%q) == %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
