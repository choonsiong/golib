// Package compare provides comparison helpers.
package compare

// Float64Maps return true if both maps have identical elements.
func Float64Maps(m1 map[float64]float64, m2 map[float64]float64) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}

// StringMaps return true if both maps have identical elements.
func StringMaps(m1 map[string]string, m2 map[string]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}

	return true
}
