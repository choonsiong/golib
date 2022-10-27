package slicex

import (
	"strconv"
	"strings"
)

type Ints []int

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

func (ints Ints) Total() (total int) {
	for _, i := range ints {
		total += i
	}

	return
}
