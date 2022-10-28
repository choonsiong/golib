package mathx

import "math"

// CompareFloats compares f1 and f2 using the given epsilon.
func CompareFloats(f1 float64, f2 float64, e float64) bool {
	if math.Abs(f1-f2) > e {
		return false
	}

	return true
}
