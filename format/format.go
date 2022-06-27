// Package format implements various formatting functions.
package format

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math"
	"strconv"
	"strings"
	"unicode"
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

// UnderscoreToUpperCamelCase returns underscore string s in language t
// in upper camel case.
// For example: foo_bar -> FooBar
func UnderscoreToUpperCamelCase(s string, t language.Tag) string {
	s = strings.Replace(s, "_", " ", -1)

	titleCase := cases.Title(t)
	s = titleCase.String(s)

	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase returns underscore string s in language t
// in lower camel case.
// For example: foo_bar -> fooBar
func UnderscoreToLowerCamelCase(s string, t language.Tag) string {
	s = UnderscoreToUpperCamelCase(s, t)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

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
					if idx+1 == i {
						// Two consecutive uppercase letter found,
						// i.e. FOobar, fOObar
						return s
					} else {
						idx = i
					}
				}
			}
		}
	}
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
