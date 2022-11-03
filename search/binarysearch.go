package search

// BinarySearchInt return true if data is in the list else returns false.
func BinarySearchInt(data int, list []int) bool {
	low, high := 0, len(list)-1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] < data {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(list) || list[low] != data {
		return false
	}

	return true
}
