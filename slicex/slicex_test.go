package slicex

import "testing"

func TestCompare_int(t *testing.T) {

	tests := []struct {
		name string
		s1   []int
		s2   []int
		want bool
	}{
		{"empty slice", []int{}, []int{}, true},
		{"valid slice", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"invalid slice", []int{1, 2}, []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare(tt.s1, tt.s2)

			if got != tt.want {
				t.Errorf("Compare(%v, %v) == %v; want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestCompare_string(t *testing.T) {

	tests := []struct {
		name string
		s1   []string
		s2   []string
		want bool
	}{
		{"empty slice", []string{}, []string{}, true},
		{"valid slice", []string{"a", "b", "c"}, []string{"a", "b", "c"}, true},
		{"invalid slice", []string{"a", "b"}, []string{"a", "b", "c"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compare(tt.s1, tt.s2)

			if got != tt.want {
				t.Errorf("Compare(%v, %v) == %v; want %v", tt.s1, tt.s2, got, tt.want)
			}
		})
	}
}

func TestContains_bool(t *testing.T) {

	tests := []struct {
		name    string
		element bool
		input   []bool
		want    bool
	}{
		{"empty slice", true, []bool{}, false},
		{"bool valid", false, []bool{true, true, true, false}, true},
		{"bool invalid", false, []bool{true, true, true, true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.element, tt.input)

			if got != tt.want {
				t.Errorf("Contains(%v, %v) == %v; want %v", tt.element, tt.input, got, tt.want)
			}
		})
	}
}

func TestContains_float(t *testing.T) {
	tests := []struct {
		name    string
		element float64
		input   []float64
		want    bool
	}{
		{"empty slice", 1.0, []float64{}, false},
		{"float64 valid", 2.0, []float64{1.0, 2.0, 3.0, 4.0}, true},
		{"float64 invalid", 5.0, []float64{1.0, 2.0, 3.0, 4.0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.element, tt.input)

			if got != tt.want {
				t.Errorf("Contains(%v, %v) == %v; want %v", tt.element, tt.input, got, tt.want)
			}
		})
	}
}

func TestContains_int(t *testing.T) {
	tests := []struct {
		name    string
		element int
		input   []int
		want    bool
	}{
		{"empty slice", 1, []int{}, false},
		{"int valid", 2, []int{1, 2, 3, 4}, true},
		{"int invalid", 5, []int{1, 2, 3, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.element, tt.input)

			if got != tt.want {
				t.Errorf("Contains(%v, %v) == %v; want %v", tt.element, tt.input, got, tt.want)
			}
		})
	}
}

func TestContains_string(t *testing.T) {

	tests := []struct {
		name    string
		element string
		input   []string
		want    bool
	}{
		{"empty slice", "", []string{}, false},
		{"string valid", "b", []string{"a", "b", "c", "d"}, true},
		{"string invalid", "e", []string{"a", "b", "c", "d"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.element, tt.input)

			if got != tt.want {
				t.Errorf("Contains(%v, %v) == %v; want %v", tt.element, tt.input, got, tt.want)
			}
		})
	}
}

func TestDeleteElementAtIndex_int(t *testing.T) {

	tests := []struct {
		name    string
		index   int
		element int
		input   []int
		want    []int
	}{
		{"int slice", 2, 3, []int{1, 2, 3, 4}, []int{1, 2, 4}},
		{"empty slice", 2, -1, []int{}, []int{}},
		{"invalid index", 4, -1, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"first index", 0, 1, []int{1, 2, 3, 4}, []int{2, 3, 4}},
		{"last index", 3, 4, []int{1, 2, 3, 4}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteElementAtIndex(tt.index, tt.input)

			if Contains(tt.element, got) {
				t.Errorf("DeleteElementAtIndex(%d, %v) == %v; want %v", tt.index, tt.input, got, tt.want)
			}
		})
	}
}

func TestDeleteElementAtIndex_string(t *testing.T) {

	tests := []struct {
		name    string
		index   int
		element string
		input   []string
		want    []string
	}{
		{"string slice", 2, "c", []string{"a", "b", "c", "d"}, []string{"a", "b", "d"}},
		{"empty slice", 2, "", []string{}, []string{}},
		{"invalid index", 4, "", []string{"a", "b", "c", "d"}, []string{"a", "b", "c", "d"}},
		{"first index", 0, "a", []string{"a", "b", "c", "d"}, []string{"b", "c", "d"}},
		{"last index", 3, "d", []string{"a", "b", "c", "d"}, []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteElementAtIndex(tt.index, tt.input)

			if Contains(tt.element, got) {
				t.Errorf("DeleteElementAtIndex(%d, %v) == %v; want %v", tt.index, tt.input, got, tt.want)
			}
		})
	}
}
