package format

import "testing"

func TestNumberZero(t *testing.T) {
	n := 0
	want := "00"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}

func TestNumberOne(t *testing.T) {
	n := 1
	want := "01"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}

func TestNumberTen(t *testing.T) {
	n := 10
	want := "10"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}

func TestNumberHundred(t *testing.T) {
	n := 100
	want := "100"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}

func TestNegativeOne(t *testing.T) {
	n := -1
	want := "-1"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}

func TestNegativeTen(t *testing.T) {
	n := -10
	want := "-10"
	result := PadWithZero(n)

	if result != want {
		t.Fatalf(`PadWithZero(%d) = %q, want "%v"`, n, result, want)
	}
}
