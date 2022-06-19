package geometry

import (
	"math"
	"testing"
)

func TestSlope(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA CoordinateXY
		coordinateB CoordinateXY
		want        float64
	}{
		{"horizontal line", CoordinateXY{0, 0}, CoordinateXY{10, 0}, 0},
		{"vertical line", CoordinateXY{0, 0}, CoordinateXY{0, 10}, math.Inf(1)},
		{"positive valid", CoordinateXY{1, 2}, CoordinateXY{5, 7}, 1.25},
		{"negative valid", CoordinateXY{1, 2}, CoordinateXY{3, -8}, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Slope(tt.coordinateA, tt.coordinateB)
			if got != tt.want {
				t.Errorf("want %v; got %v", tt.want, got)
			}
		})
	}
}

func TestMidpoint(t *testing.T) {
	// testing..
}
