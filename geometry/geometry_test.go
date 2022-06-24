package geometry

import (
	"math"
	"testing"
)

func TestSlope(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA Coordinate2D
		coordinateB Coordinate2D
		want        float64
	}{
		{"horizontal line", Coordinate2D{0, 0}, Coordinate2D{10, 0}, 0},
		{"vertical line", Coordinate2D{0, 0}, Coordinate2D{0, 10}, math.Inf(1)},
		{"positive slope", Coordinate2D{0, 0}, Coordinate2D{10, 10}, 1},
		{"negative slope", Coordinate2D{0, 0}, Coordinate2D{-10, 10}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Slope(tt.coordinateA, tt.coordinateB)
			if got != tt.want {
				t.Errorf("want %v; got %v", tt.want, got)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coordinateA.Slope(tt.coordinateB)
			if got != tt.want {
				t.Errorf("want %v; got %v", tt.want, got)
			}
		})
	}
}

func TestMidpoint(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA Coordinate2D
		coordinateB Coordinate2D
		midpointX   float64
		midpointY   float64
	}{
		{"positive coordinates", Coordinate2D{0, 0}, Coordinate2D{10, 10}, 5, 5},
		{"negative coordinates", Coordinate2D{0, 0}, Coordinate2D{-10, -10}, -5, -5},
		{"vertical line", Coordinate2D{0, 10}, Coordinate2D{0, 0}, 0, 5},
		{"horizontal line", Coordinate2D{0, 0}, Coordinate2D{10, 0}, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantX, wantY := Midpoint(tt.coordinateA, tt.coordinateB)
			if tt.midpointX != wantX {
				t.Errorf("want midpoint x %v; got %v", tt.midpointX, wantX)
			}

			if tt.midpointY != wantY {
				t.Errorf("want midpoint y %v; got %v", tt.midpointY, wantY)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantX, wantY := tt.coordinateA.Midpoint(tt.coordinateB)
			if tt.midpointX != wantX {
				t.Errorf("want midpoint x %v; got %v", tt.midpointX, wantX)
			}

			if tt.midpointY != wantY {
				t.Errorf("want midpoint y %v; got %v", tt.midpointY, wantY)
			}
		})
	}
}
