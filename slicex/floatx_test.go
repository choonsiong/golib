package slicex

import "testing"

func TestFloats_String(t *testing.T) {
	tests := []struct {
		name  string
		input Floats
		want  string
	}{
		{"valid input", []float64{1.1, 2.2, 3.3, 4.4}, "[0:1.1,1:2.2,2:3.3,3:4.4]"},
		{"empty slice", []float64{}, "[]"},
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

func TestFloats_Total(t *testing.T) {
	tests := []struct {
		name  string
		input Floats
		want  float64
	}{
		{"[1.0,2.0,3.0,4.0]", []float64{1.0, 2.0, 3.0, 4.0}, 10.0},
		{"empty slice", []float64{}, 0},
		{"[-1.0,-2.0,-3.0,-4.0]", []float64{-1.0, -2.0, -3.0, -4.0}, -10.0},
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
