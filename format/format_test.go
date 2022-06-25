package format

import "testing"

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
