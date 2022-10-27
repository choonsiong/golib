package slicex

import (
	"strconv"
	"strings"
)

type Ints []int

// String returns a nicely formatted output to use in fmt functions.
func (ints Ints) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	for idx, int := range ints {
		sb.WriteString(strconv.Itoa(idx) + ":")
		sb.WriteString(strconv.Itoa(int))
		if idx != (len(ints) - 1) {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")

	return sb.String()
}

// Total returns the total value of the slice of int.
func (ints Ints) Total() (total int) {
	for _, i := range ints {
		total += i
	}

	return
}
