package slicex

import (
	"reflect"
)

// Compare returns true if both slices have identical elements,
// else returns false.
func Compare[T any](s1 []T, s2 []T) bool {
	return reflect.DeepEqual(s1, s2)
}

// Contains returns true if element e is found in slice s else returns false.
// The type T can be either int, string or bool only.
func Contains[T any](e T, s []T) bool {
	r1 := reflect.ValueOf(&e).Elem()

	for _, v := range s {
		r2 := reflect.ValueOf(&v).Elem()
		switch r2.Type().Name() {
		case "int":
			if r1.Int() == r2.Int() {
				return true
			}
		case "string":
			if r1.String() == r2.String() {
				return true
			}
		case "bool":
			if r1.Bool() == r2.Bool() {
				return true
			}
		}
	}

	return false
}

// DeleteElementAtIndex returns a new slice with element at index deleted.
func DeleteElementAtIndex[T any](i int, s []T) []T {
	if len(s) == 0 {
		return s
	}

	if i > (len(s) - 1) {
		return s
	}

	return append(s[:i], s[i+1:]...)
}
