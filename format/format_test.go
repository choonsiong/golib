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
