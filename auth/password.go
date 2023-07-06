// Package auth provides types and functions for working with
// password verification.
package auth

import "unicode"

type Password struct {
	MinPasswordLength    int
	WantUppercase        bool
	WantLowercase        bool
	WantNumber           bool
	WantSpecialCharacter bool
}

// CheckPassword returns true if pw satisfied the password requirements defined
// in the Password struct else returns false.
func (p *Password) CheckPassword(pw string) bool {
	runes := []rune(pw)
	if len(runes) < p.MinPasswordLength {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range runes {
		if p.WantUppercase && unicode.IsUpper(v) {
			hasUpper = true
		}
		if p.WantLowercase && unicode.IsLower(v) {
			hasLower = true
		}
		if p.WantNumber && unicode.IsNumber(v) {
			hasNumber = true
		}
		if p.WantSpecialCharacter && (unicode.IsPunct(v) || unicode.IsSymbol(v)) {
			hasSymbol = true
		}
	}

	if !p.WantUppercase {
		hasUpper = true
	}

	if !p.WantLowercase {
		hasLower = true
	}

	if !p.WantNumber {
		hasNumber = true
	}

	if !p.WantSpecialCharacter {
		hasSymbol = true
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
