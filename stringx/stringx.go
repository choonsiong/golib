// Package stringx provides extra functions to work with string.
package stringx

import (
	"strings"
)

// CapitalizeEachWord returns s with each word capitalized.
func CapitalizeEachWord(s string) string {
	if s == "" {
		return ""
	}

	ss := strings.Split(strings.ToLower(s), " ")
	var str string

	for _, w := range ss {
		bs := []byte(w)

		// Ignore all non-English characters
		if bs[0] < 97 || bs[0] > 122 {
			return s
		}

		bs[0] -= 32
		str += string(bs)
		str += " "
	}

	return strings.TrimSpace(str)
}
