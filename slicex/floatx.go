package slicex

import (
	"strconv"
	"strings"
)

type Floats []float64

// String returns a nicely formatted output to use in fmt functions.
func (fs Floats) String() string {
	var sb strings.Builder

	sb.WriteString("[")

	for i, f := range fs {
		sb.WriteString(strconv.Itoa(i) + ":")
		sb.WriteString(strconv.FormatFloat(f, 'f', -1, 64))
		if i != (len(fs) - 1) {
			sb.WriteString(",")
		}
	}

	sb.WriteString("]")

	return sb.String()
}

// Total returns the total value of the slice of float64.
func (fs Floats) Total() (total float64) {
	for _, i := range fs {
		total += i
	}

	return
}
