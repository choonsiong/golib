package format

import "strconv"

// StringZeroToNine returns a single digit number padded with '0' in front.
// '0' => '00', '1' => '01' etc.
func StringZeroToNine(i int) string {
	if i >= 0 && i < 10 {
		return "0" + strconv.Itoa(i)
	}

	return strconv.Itoa(i)
}
