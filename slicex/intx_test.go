package slicex

import "testing"

func TestInts_String(t *testing.T) {
	tests := []struct {
		name  string
		input Ints
		want  string
	}{
		{"valid input", []int{1, 2, 3, 4}, "[0:1,1:2,2:3,3:4]"},
		{"empty slice", []int{}, "[]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.String()
			if got != tt.want {
				t.Errorf("String() == %v; want %v", got, tt.want)
			}
		})
	}
}

func TestInts_Total(t *testing.T) {
	tests := []struct {
		name  string
		input Ints
		want  int
	}{
		{"[1,2,3,4]", []int{1, 2, 3, 4}, 10},
		{"empty slice", []int{}, 0},
		{"[-1,-2,-3,-4]", []int{-1, -2, -3, -4}, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.Total()
			if got != tt.want {
				t.Errorf("Total() == %v; want %v", got, tt.want)
			}
		})
	}
}
