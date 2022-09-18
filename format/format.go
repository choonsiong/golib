// Package format implements various formatting functions.
package format

import (
	"math"
	"strconv"
)

// PadZeroToNineWithZero pad i with '0' in front.
// For example: '0' => '00', '1' => '01' ... '9' => '09'
func PadZeroToNineWithZero(i int) string {
	i = int(math.Abs(float64(i)))
	if i > 10 {
		return strconv.Itoa(i)
	}
	return "0" + strconv.Itoa(i)
}

// LeftPaddingWithSize pads the string s with enough of character ch
// to the length l.
// For example: LeftPaddingWithSize(10, "hello", "*") will produce
// the string "*****hello"
func LeftPaddingWithSize(l int, s string, ch string) string {
	if l <= len(s) {
		return s
	}
	str := ""
	for i := 0; i < l; i++ {
		if i < (l - len(s)) {
			str += ch
		}
	}
	str += s
	return str
}

// RightPaddingWithSize pads the string s with enough of character ch
// to the length l.
// For example: RightPaddingWithSize(10, "hello", "*") will produce
// the string "hello*****"
func RightPaddingWithSize(l int, s string, ch string) string {
	if l <= len(s) {
		return s
	}
	str := s
	for i := len(s); i < l; i++ {
		str += ch
	}
	return str
}
