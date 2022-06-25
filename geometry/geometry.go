// Package geometry provides helpers to work with geometry.
package geometry

import "math"

// Coordinate2D represents a coordinate in two-dimensional space.
type Coordinate2D struct {
	X float64
	Y float64
}

// Coordinate3D represents a coordinate in three-dimensional space.
type Coordinate3D struct {
	X float64
	Y float64
	Z float64
}

// Slope returns the slope of endpoints A and B.
func (a Coordinate2D) Slope(b Coordinate2D) float64 {
	return Slope(a, b)
}

// Slope returns the slope of endpoints A and B.
func Slope(a Coordinate2D, b Coordinate2D) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}

// Midpoint returns the midpoint of endpoints A and B.
func (a Coordinate2D) Midpoint(b Coordinate2D) (float64, float64) {
	return Midpoint(a, b)
}

// Midpoint returns the midpoint of endpoints A and B.
func Midpoint(a Coordinate2D, b Coordinate2D) (float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2

	return x, y
}

// Distance returns the distance of endpoints A and B in three-dimensional space.
func (a Coordinate3D) Distance(b Coordinate3D) float64 {
	return Distance3D(a, b)
}

// Distance3D returns the distance of endpoints A and B in three-dimensional space.
func Distance3D(a Coordinate3D, b Coordinate3D) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

// Midpoint returns the midpoint of endpoints A and B in three-dimensional space.
func (a Coordinate3D) Midpoint(b Coordinate3D) (float64, float64, float64) {
	return Midpoint3D(a, b)
}

// Midpoint3D returns the midpoint of endpoints A and B in three-dimensional space.
func Midpoint3D(a Coordinate3D, b Coordinate3D) (float64, float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2
	z := (a.Z + b.Z) / 2

	return x, y, z
}

// EndPointB returns the endpoint B with given endpoint A and midpoint M.
func (a Coordinate3D) EndPointB(m Coordinate3D) (float64, float64, float64) {
	return EndPoint3D(a, m)
}

// EndPoint3D returns the endpoint B with given endpoint A and midpoint M.
func EndPoint3D(a Coordinate3D, m Coordinate3D) (float64, float64, float64) {
	x := (m.X * 2) - a.X
	y := (m.Y * 2) - a.Y
	z := (m.Z * 2) - a.Z

	return x, y, z
}