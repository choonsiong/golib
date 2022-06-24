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
			gotX, gotY := Midpoint(tt.coordinateA, tt.coordinateB)
			if tt.midpointX != gotX {
				t.Errorf("want midpoint x %v; got %v", tt.midpointX, gotX)
			}

			if tt.midpointY != gotY {
				t.Errorf("want midpoint y %v; got %v", tt.midpointY, gotY)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := tt.coordinateA.Midpoint(tt.coordinateB)
			if tt.midpointX != gotX {
				t.Errorf("want midpoint x %v; got %v", tt.midpointX, gotX)
			}

			if tt.midpointY != gotY {
				t.Errorf("want midpoint y %v; got %v", tt.midpointY, gotY)
			}
		})
	}
}

func TestDistance3D(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA Coordinate3D
		coordinateB Coordinate3D
		want        float64
	}{
		{"zero", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, 0}, 0},
		{"positive z", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, 10}, 10},
		{"negative z", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, -10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distance3D(tt.coordinateA, tt.coordinateB)
			if tt.want != got {
				t.Errorf("want %v; got %v", tt.want, got)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.coordinateA.Distance(tt.coordinateB)
			if tt.want != got {
				t.Errorf("want %v; got %v", tt.want, got)
			}
		})
	}
}

func TestMidpoint3D(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA Coordinate3D
		coordinateB Coordinate3D
		wantX       float64
		wantY       float64
		wantZ       float64
	}{
		{"zero", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, 0}, 0, 0, 0},
		{"positive x", Coordinate3D{0, 0, 0}, Coordinate3D{10, 0, 0}, 5, 0, 0},
		{"negative x", Coordinate3D{0, 0, 0}, Coordinate3D{-10, 0, 0}, -5, 0, 0},
		{"positive y", Coordinate3D{0, 0, 0}, Coordinate3D{0, 10, 0}, 0, 5, 0},
		{"negative y", Coordinate3D{0, 0, 0}, Coordinate3D{0, -10, 0}, 0, -5, 0},
		{"positive z", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, 10}, 0, 0, 5},
		{"negative z", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, -10}, 0, 0, -5},
		{"all positive", Coordinate3D{0, 0, 0}, Coordinate3D{5, 5, 10}, 2.5, 2.5, 5},
		{"all negative", Coordinate3D{0, 0, 0}, Coordinate3D{-5, -5, -10}, -2.5, -2.5, -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotZ := Midpoint3D(tt.coordinateA, tt.coordinateB)
			if tt.wantX != gotX {
				t.Errorf("want %v; got %v", tt.wantX, gotX)
			}
			if tt.wantY != gotY {
				t.Errorf("want %v; got %v", tt.wantY, gotY)
			}
			if tt.wantZ != gotZ {
				t.Errorf("want %v; got %v", tt.wantZ, gotZ)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotZ := tt.coordinateA.Midpoint(tt.coordinateB)
			if tt.wantX != gotX {
				t.Errorf("want %v; got %v", tt.wantX, gotX)
			}
			if tt.wantY != gotY {
				t.Errorf("want %v; got %v", tt.wantY, gotY)
			}
			if tt.wantZ != gotZ {
				t.Errorf("want %v; got %v", tt.wantZ, gotZ)
			}
		})
	}
}

func TestEndPoint3D(t *testing.T) {
	tests := []struct {
		name        string
		coordinateA Coordinate3D
		midpoint    Coordinate3D
		wantX       float64
		wantY       float64
		wantZ       float64
	}{
		{"zero", Coordinate3D{0, 0, 0}, Coordinate3D{0, 0, 0}, 0, 0, 0},
		{"positive", Coordinate3D{0, 0, 0}, Coordinate3D{5, 5, 5}, 10, 10, 10},
		{"negative", Coordinate3D{0, 0, 0}, Coordinate3D{-5, -5, -5}, -10, -10, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotZ := EndPoint3D(tt.coordinateA, tt.midpoint)
			if tt.wantX != gotX {
				t.Errorf("want %v; got %v", tt.wantX, gotX)
			}
			if tt.wantY != gotY {
				t.Errorf("want %v; got %v", tt.wantY, gotY)
			}
			if tt.wantZ != gotZ {
				t.Errorf("want %v; got %v", tt.wantZ, gotZ)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, gotZ := tt.coordinateA.EndPointB(tt.midpoint)
			if tt.wantX != gotX {
				t.Errorf("want %v; got %v", tt.wantX, gotX)
			}
			if tt.wantY != gotY {
				t.Errorf("want %v; got %v", tt.wantY, gotY)
			}
			if tt.wantZ != gotZ {
				t.Errorf("want %v; got %v", tt.wantZ, gotZ)
			}
		})
	}
}
