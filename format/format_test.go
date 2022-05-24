package format

import "testing"

func TestPadZeroToNineWithZero(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := PadZeroToNineWithZero(c.in)

			if got != c.want {
				t.Errorf("PadIntWithZero(%q) == %q, want %q", c.in, got, c.want)
			}
		})

	}
}

func TestLeftPaddingWithSize(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := LeftPaddingWithSize(c.length, c.source, c.character)

			if got != c.want {
				t.Errorf("LeftPaddingWithSize(%d, %q, %q) == %q, want %q", c.length, c.source, c.character, got, c.want)
			}
		})
	}
}

func TestRightPaddingWithSize(t *testing.T) {
	cases := []struct {
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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := RightPaddingWithSize(c.length, c.source, c.character)

			if got != c.want {
				t.Errorf("RightPaddingWithSize(%d, %q, %q) == %q, want %q", c.length, c.source, c.character, got, c.want)
			}
		})
	}
}
