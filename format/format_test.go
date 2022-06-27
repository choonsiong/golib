package format

import (
	"golang.org/x/text/language"
	"testing"
)

func TestPadZeroToNineWithZero(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{"one", 1, "01"},
		{"zero", 0, "00"},
		{"negative input", -1, "01"},
		{"eleven", 11, "11"},
		{"hundred", 100, "100"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadZeroToNineWithZero(tt.in)

			if got != tt.want {
				t.Errorf("PadIntWithZero(%q) == %q; want %q", tt.in, got, tt.want)
			}
		})

	}
}

func TestLeftPaddingWithSize(t *testing.T) {
	tests := []struct {
		name      string
		length    int
		source    string
		character string
		want      string
	}{
		{"valid", 10, "hello", "*", "*****hello"},
		{"invalid", 5, "helloworld", "*", "helloworld"},
		{"same", 5, "hello", "*", "hello"},
		{"negative", -10, "hello", "*", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LeftPaddingWithSize(tt.length, tt.source, tt.character)

			if got != tt.want {
				t.Errorf("LeftPaddingWithSize(%d, %q, %q) == %q; want %q", tt.length, tt.source, tt.character, got, tt.want)
			}
		})
	}
}

func TestRightPaddingWithSize(t *testing.T) {
	tests := []struct {
		name      string
		length    int
		source    string
		character string
		want      string
	}{
		{"valid", 10, "hello", "*", "hello*****"},
		{"invalid", 5, "helloworld", "*", "helloworld"},
		{"same", 5, "hello", "*", "hello"},
		{"negative", -10, "hello", "*", "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RightPaddingWithSize(tt.length, tt.source, tt.character)

			if got != tt.want {
				t.Errorf("RightPaddingWithSize(%d, %q, %q) == %q; want %q", tt.length, tt.source, tt.character, got, tt.want)
			}
		})
	}
}

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		language language.Tag
		want     string
	}{
		{"foobar", "foobar", language.English, "Foobar"},
		{"foo_bar", "foo_bar", language.English, "FooBar"},
		{"FOOBAR", "FOOBAR", language.English, "Foobar"},
		{"FOO_BAR", "FOO_BAR", language.English, "FooBar"},
		{"_foo_bar", "_foo_bar", language.English, "FooBar"},
		{"foo_bar_", "foo_bar_", language.English, "FooBar"},
		{"foo__bar", "foo__bar", language.English, "FooBar"},
		{"foo bar", "foo bar", language.English, "FooBar"},
		{"Foo Bar", "Foo Bar", language.English, "FooBar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnderscoreToUpperCamelCase(tt.word, tt.language)

			if got != tt.want {
				t.Errorf("UnderscoreToUpperCamelCase(%q, %v) == %q; want %q", tt.word, tt.language, got, tt.want)
			}
		})
	}
}

func TestUnderscoreToLowerCamelCase(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		language language.Tag
		want     string
	}{
		{"foobar", "foobar", language.English, "foobar"},
		{"foo_bar", "foo_bar", language.English, "fooBar"},
		{"FOOBAR", "FOOBAR", language.English, "foobar"},
		{"FOO_BAR", "FOO_BAR", language.English, "fooBar"},
		{"_foo_bar", "_foo_bar", language.English, "fooBar"},
		{"foo_bar_", "foo_bar_", language.English, "fooBar"},
		{"foo__bar", "foo__bar", language.English, "fooBar"},
		{"foo bar", "foo bar", language.English, "fooBar"},
		{"Foo Bar", "Foo Bar", language.English, "fooBar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UnderscoreToLowerCamelCase(tt.word, tt.language)

			if got != tt.want {
				t.Errorf("UnderscoreToLowerCamelCase((%q, %v) == %q; want %q", tt.word, tt.language, got, tt.want)
			}
		})
	}
}

func TestCamelCaseToUnderscore(t *testing.T) {
	tests := []struct {
		name string
		word string
		want string
	}{
		{"foobar", "foobar", "foobar"},
		{"FooBar", "FooBar", "foo_bar"},
		{"fooBar", "fooBar", "foo_bar"},
		{"Foobar", "Foobar", "foobar"},
		{"foo_bar", "foo_bar", "foo_bar"},
		{"foo__bar", "foo__bar", "foo__bar"},
		{"Foo_Bar", "Foo_Bar", "Foo_Bar"},
		{"FOOBAR", "FOOBAR", "FOOBAR"},
		{"FOObar", "FOObar", "FOObar"},
		{"FOO_BAR", "FOO_BAR", "FOO_BAR"},
		{"_foo_bar", "_foo_bar", "_foo_bar"},
		{"foo_bar_", "foo_bar_", "foo_bar_"},
		{"foo bar", "foo bar", "foo bar"},
		{"Foo Bar", "Foo Bar", "Foo Bar"},
		{"fOObar", "fOObar", "fOObar"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CamelCaseToUnderscore(tt.word)

			if got != tt.want {
				t.Errorf("CamelCaseToUnderscore((%q) == %q; want %q", tt.word, got, tt.want)
			}
		})
	}
}
