package search

import "testing"

func TestBinarySearchInt(t *testing.T) {
	tests := []struct {
		name string
		data int
		list []int
		want bool
	}{
		{"valid search", 4, []int{1, 2, 3, 4, 5, 6}, true},
		{"invalid search", 4, []int{1, 2, 5, 6, 8}, false},
		{"left boundary search", 0, []int{0, 1, 2, 3, 4, 5}, true},
		{"right boundary search", 10, []int{1, 2, 3, 5, 6, 9, 10}, true},
		{"high number", 100, []int{1, 2, 3, 4, 5, 6}, false},
		{"low number", -1, []int{9, 10, 11, 13, 19, 21}, false},
		{"negative number", -3, []int{-9, -8, -3}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearchInt(tt.data, tt.list)

			if got != tt.want {
				t.Errorf("BinarySearchInt() == %v; want %v", got, tt.want)
			}
		})
	}
}
