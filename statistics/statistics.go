package statistics

// Average returns the mean of numbers.
func Average(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0.0, ErrEmptySlice
	}

	total := 0.0

	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers)), nil
}

// DeviationFromMean returns the deviation from mean for numbers.
func DeviationFromMean(numbers ...float64) map[float64]float64 {
	result := map[float64]float64{}

	average, _ := Average(numbers...)

	for _, num := range numbers {
		deviation := num - average
		result[num] = deviation
	}

	return result
}
