// Package format implements various formatting functions.
package format

import (
	"math"
	"strconv"
)

// PadZeroToNineWithZero returns a string with its argument int padded with '0' in front.
// For example: '0' => '00', '1' => '01'...
func PadZeroToNineWithZero(i int) string {
	i = int(math.Abs(float64(i)))

	if i > 10 {
		return strconv.Itoa(i)
	}

	return "0" + strconv.Itoa(i)
}
