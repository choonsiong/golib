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

func TestPositiveTimezone(t *testing.T) {
	hr := 16
	tz := 8
	want := 0
	result := NormalizeHour(hr, tz)

	if result != want {
		t.Fatalf(`NormalizeHour(%d, %d) = %v, want "%v"`, hr, tz, result, want)
	}
}

func TestPositiveRandomTimezone(t *testing.T) {
	hr := 10
	tz := 8
	want := 18
	result := NormalizeHour(hr, tz)

	if result != want {
		t.Fatalf(`NormalizeHour(%d, %d) = %v, want "%v"`, hr, tz, result, want)
	}
}

func TestNegativeTimezone(t *testing.T) {
	hr := 0
	tz := -8
	want := 16
	result := NormalizeHour(hr, tz)

	if result != want {
		t.Fatalf(`NormalizeHour(%d, %d) = %v, want "%v"`, hr, tz, result, want)
	}
}

func TestNegativeRandomTimezone(t *testing.T) {
	hr := 5
	tz := -11
	want := 18
	result := NormalizeHour(hr, tz)

	if result != want {
		t.Fatalf(`NormalizeHour(%d, %d) = %v, want "%v"`, hr, tz, result, want)
	}
}