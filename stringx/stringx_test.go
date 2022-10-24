package stringx

import (
	"errors"
	"golang.org/x/text/language"
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

func TestRandomPassword(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		want    int
		wantErr error
	}{
		{"length -1", -1, 0, ErrInvalidLength},
		{"length 0", 0, 0, nil},
		{"length 10", 10, 10, nil},
		{"length 20", 20, 20, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := RandomPassword(tt.length)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("RandomPassword(%v), want error %v; got nil", tt.length, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("RandomPassword(%v), want error %v; got %v", tt.length, tt.wantErr, err)
				}
			}

			got := len(s)
			if got != tt.want {
				t.Errorf("RandomPassword(%d)'s length == %d; want %d", tt.length, got, tt.want)
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
		{"length -1", -1, 0, ErrInvalidLength},
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

func TestRandomStringIgnoreError(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   int
	}{
		{"length -1", -1, 0},
		{"length 0", 0, 0},
		{"length 10", 10, 10},
		{"length 20", 20, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := RandomStringIgnoreError(tt.length)

			got := len(s)

			if got != tt.want {
				t.Errorf("RandomStringIgnoreError(%d)'s length == %d; want %d", tt.length, got, tt.want)
			}
		})
	}
}

func TestSlugify(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{"valid string", "now is the time", "now-is-the-time", nil},
		{"empty string", "", "", ErrEmptyString},
		{"complex string", "h@llo#$%%^WORLD", "h-llo-world", nil},
		{"non-english string", "ありがとうございました", "", ErrEmptySlug},
		{"mixed string", "helloありがとうございましたworld", "hello-world", nil},
	}

	for _, tt := range tests {
		got, err := Slugify(tt.input)

		if tt.wantErr != nil {
			if err == nil {
				t.Errorf("Slugify(%v), want error %v; got nil", tt.input, tt.wantErr)
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Slugify(%v), want error %v; got %v", tt.input, tt.wantErr, err)
			}
		}

		if got != tt.want {
			t.Errorf("Slugify(%q) == %q; want %q", tt.input, got, tt.want)
		}
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
		{" Foo Foo_Bar", " Foo Foo_Bar", language.English, "FooFooBar"},
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
		{"Foo Foo_Bar", "Foo Foo_Bar", language.English, "fooFooBar"},
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
				t.Errorf("CamelCaseToUnderscore(%q) == %q; want %q", tt.word, got, tt.want)
			}
		})
	}
}

func TestTrimExtraWhiteSpacesInOut(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"  foo bar  ", "  foo bar  ", "foo bar"},
		{"       foo     bar    ", "       foo     bar    ", "foo bar"},
		{"foo     bar", "foo     bar", "foo bar"},
		{"foo     bar          alice smith ", "foo     bar          alice smith ", "foo bar alice smith"},
		{"foobar", "foobar", "foobar"},
		{"f o o b a r", "f o o b a r", "f o o b a r"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrimExtraWhiteSpacesInOut(tt.input)

			if got != tt.want {
				t.Errorf("TrimExtraWhiteSpacesInOut(%q) == %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
