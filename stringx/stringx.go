package stringx

import (
	"strings"
)

// CapitalizeEachWord returns s with first character in upper case.
func CapitalizeEachWord(s string) string {
	if s == "" {
		return ""
	}

	ss := strings.Split(strings.ToLower(s), " ")
	var str string

	for _, w := range ss {
		bs := []byte(w)

		// Ignore all non-English characters
		if bs[0] > 127 {
			return s
		}

		bs[0] -= 32
		str += string(bs)
		str += " "
	}

	return strings.TrimSpace(str)
}
