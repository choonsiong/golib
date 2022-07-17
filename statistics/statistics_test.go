package statistics

import (
	"errors"
	"testing"
)

func TestAverage(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
		want    float64
		wantErr error
	}{
		{"one number", []float64{1.0}, 1.0, nil},
		{"empty slice", []float64{}, 0.0, ErrEmptySlice},
		{"many numbers", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, 3.0, nil},
		{"with negative numbers", []float64{-1.0, 2.0, -3.0, 4.0}, 0.5, nil},
		{"all negative numbers", []float64{-1.0, -2.0, -3.0, -4.0, -5.0}, -3.0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Average(tt.numbers...)
			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("Average(%v), want error %v; got nil", tt.numbers, tt.wantErr)
				}
				if !errors.Is(tt.wantErr, err) {
					t.Errorf("Average(%v), want error %v; got %v", tt.numbers, tt.wantErr, err)
				}
			}
			if got != tt.want {
				t.Errorf("Average(%v) == %v; want %v", tt.numbers, got, tt.want)
			}
		})
	}
}

func TestDeviationFromMean(t *testing.T) {
	tests := []struct {
		name    string
		numbers []float64
		want    map[float64]float64
	}{
		{"one number", []float64{1.0}, map[float64]float64{1.0: 0.0}},
		{"many numbers", []float64{1.0, 2.0, 3.0, 4.0, 5.0}, map[float64]float64{1.0: -2.0, 2.0: -1.0, 3.0: 0.0, 4.0: 1.0, 5.0: 2.0}},
		{"empty slice", []float64{}, nil},
		{"negative numbers", []float64{-1.0, -2.0, -3.0, -4.0, -5.0}, map[float64]float64{-1.0: 2.0, -2.0: 1.0, -3.0: 0.0, -4.0: -1.0, -5.0: -2.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeviationFromMean(tt.numbers...)
			if !isFloat64MapsSame(got, tt.want) {
				t.Errorf("DeviationFromMean(%v) == %v; want %v", tt.numbers, got, tt.want)
			}
		})
	}
}

func isFloat64MapsSame(m1 map[float64]float64, m2 map[float64]float64) bool {
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
