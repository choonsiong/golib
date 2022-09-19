// Package form provides helpers to work with form validation.
package form

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct.
func New(data url.Values) *Form {
	return &Form{
		data,
		errors{},
	}
}

// Valid returns true if no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if form field is in post and not empty.
func (f *Form) Has(field string) bool {
	v := f.Get(field)
	if v == "" {
		return false
	}
	return true
}

// MinLength checks for string minimum length.
func (f *Form) MinLength(field string, length int) bool {
	v := f.Get(field)
	if len(v) < length {
		f.Errors.Add(field, fmt.Sprintf("This field has a minimum length of %d characters", length))
		return false
	}
	return true
}

// MaxLength checks for string maximum length.
func (f *Form) MaxLength(field string, length int) bool {
	v := f.Get(field)
	if len(v) > length {
		f.Errors.Add(field, fmt.Sprintf("This field has a maximum length of %d characters", length))
		return false
	}
	return true
}

// Required checks for required fields.
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		v := f.Get(field)
		if strings.TrimSpace(v) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// IsEmail checks for valid email address.
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
