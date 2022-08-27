// Package stringx provides extra functions to work with string.
package stringx

import (
	"crypto/rand"
	"log"
	"strings"
)

// CapitalizeEachWord returns string s with each word capitalized.
func CapitalizeEachWord(s string) string {
	if s == "" {
		return ""
	}

	if len([]byte(s)) != len([]rune(s)) {
		return s
	}

	ss := strings.Split(strings.ToLower(s), " ")
	var result string

	for _, ch := range ss {
		bs := []byte(ch)

		// Ignore all non-English characters
		if bs[0] < 97 || bs[0] > 122 {
			return s
		}

		bs[0] -= 32
		result += string(bs)
		result += " "
	}

	return strings.TrimSpace(result)
}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// RandomString returns a string of random characters of length,
// using randomStringSource as the source for the string.
func RandomString(length int) string {
	s, r := make([]rune, length), []rune(randomStringSource)
	for i := range s { // i is index
		p, err := rand.Prime(rand.Reader, len(r))
		if err != nil {
			log.Fatal(err)
		}

		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
