// Package stringx provides extra helpers to work with string.
package stringx

import (
	"crypto/rand"
	"fmt"
	"regexp"
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

// RandomPassword returns a random password for the given length.
func RandomPassword(length int) (string, error) {
	return RandomString(length)
}

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// RandomString returns a string of random characters of length, using
// randomStringSource as the source for the string.
func RandomString(length int) (string, error) {
	if length < 0 {
		return "", fmt.Errorf("%w: %v", ErrInvalidInput, length)
	}

	s, r := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		p, err := rand.Prime(rand.Reader, len(r))
		if err != nil {
			return "", fmt.Errorf("%w: %v", ErrGenerateRandomString, err)
		}

		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s), nil
}

// RandomStringIgnoreError is a convenient method for RandomString.
func RandomStringIgnoreError(length int) string {
	s, _ := RandomString(length)
	return s
}

// Slugify returns a string with all non-english letters and non-digits with
// '-', for example, 'hello world' -> 'hello-world'.
func Slugify(s string) (string, error) {
	if s == "" {
		return "", ErrEmptyString
	}

	// Match all non-english and non-digits characters
	var re = regexp.MustCompile(`[^a-z\d]+`)

	slug := strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")

	if len(slug) == 0 {
		return "", ErrEmptySlug
	}

	return slug, nil
}
