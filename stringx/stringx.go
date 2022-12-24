// Package stringx provides helpers to work with string.
package stringx

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"unicode"
)

// CamelCaseToUnderscore returns camel case string s in underscore style.
// For example: FooBar -> foo_bar
func CamelCaseToUnderscore(s string) string {
	if strings.Contains(s, "_") || strings.Contains(s, " ") {
		return s
	}

	if len(s) > 1 {
		var idx, count int
		for i, r := range s {
			if unicode.IsUpper(r) {
				count++
				if i == 0 {
					idx = i
					continue
				} else {
					// If idx + 1 is equal to i, means the previous index
					// and the current index are both uppercase letters,
					// e.g. FOobar, f00bar
					if idx+1 == i {
						return s
					} else {
						idx = i
					}
				}
			}
		}
	}

	var builder strings.Builder

	for i, r := range s {
		if i == 0 {
			builder.WriteRune(unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			builder.WriteRune('_')
		}

		builder.WriteRune(unicode.ToLower(r))
	}

	return builder.String()
}

// CapitalizeEachWord returns string s with each word capitalized.
// This function works on ASCII letters only.
func CapitalizeEachWord(s string) string {
	// Below line has same effect, but it also handles those words
	// start with a non-ASCII letter, e.g. \hello -> \Hello which
	// might not what we want.
	//return strings.TrimSpace(strings.Title(s))

	if s == "" {
		return ""
	}

	// Handle multi-bytes characters
	if len([]byte(s)) != len([]rune(s)) {
		return s
	}

	words := strings.Split(strings.ToLower(s), " ")

	// Using strings.Builder is more efficient than concatenating
	// regular string values.
	var result strings.Builder

	for _, word := range words {
		bs := []byte(word)

		// ASCII letters only
		if bs[0] < 97 || bs[0] > 122 {
			return s
		}

		bs[0] -= 32

		result.Write(bs)
		result.WriteRune(' ')
	}

	return strings.TrimSpace(result.String())
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
		return "", fmt.Errorf("%w: %v", ErrInvalidLength, length)
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
// '-'. For example, 'hello world' -> 'hello-world'.
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

// UnderscoreToLowerCamelCase returns the underscore string s in language t
// in lower camel case.
// For example: foo_bar -> fooBar
func UnderscoreToLowerCamelCase(s string, t language.Tag) string {
	s = UnderscoreToUpperCamelCase(s, t)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// UnderscoreToUpperCamelCase returns the underscore string s in language t
// in all upper camel case.
// For example: foo_bar -> FooBar
func UnderscoreToUpperCamelCase(s string, t language.Tag) string {
	r := strings.ReplaceAll(s, "_", " ")

	titleCase := cases.Title(t)
	r = titleCase.String(r)

	return strings.ReplaceAll(r, " ", "")
}

// TrimExtraWhiteSpacesInOut returns a string with extra whitespaces
// surrounded each words removed.
// For example: " foo     bar  alice    smith" -> "foo bar alice smith"
func TrimExtraWhiteSpacesInOut(s string) string {
	return strings.Join(strings.Fields(s), " ")

	//str := ""
	//count := 0
	//inWhitespace := false
	//
	//for _, ch := range strings.TrimSpace(s) {
	//	// If we encounter the whitespace character first time
	//	if ch == ' ' && count == 0 {
	//		// Set inWhitespace to true to indicate we are in the
	//		// whitespace now
	//		inWhitespace = true
	//		count += 1
	//	} else if ch == ' ' && count > 0 {
	//		// We are still in the whitespace, so continue to next
	//		// character
	//		if inWhitespace {
	//			continue
	//		}
	//	} else {
	//		// Out of whitespace
	//		inWhitespace = false
	//		count = 0
	//	}
	//
	//	str += string(ch)
	//}
	//
	//return str
}
