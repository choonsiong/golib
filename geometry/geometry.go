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

// Distance returns the distance between coordinate a and coordinate b.
func (a Coordinate2D) Distance(b Coordinate2D) float64 {
	return Distance(a, b)
}

// Slope returns the slope of coordinate a and coordinate b.
func (a Coordinate2D) Slope(b Coordinate2D) float64 {
	return Slope(a, b)
}

// Midpoint returns the midpoint of coordinate a and coordinate b.
func (a Coordinate2D) Midpoint(b Coordinate2D) (float64, float64) {
	return Midpoint(a, b)
}

// Distance returns the distance of coordinates a and b in three-dimensional space.
func (a Coordinate3D) Distance(b Coordinate3D) float64 {
	return Distance3D(a, b)
}

// Midpoint returns the midpoint of coordinates a and b in three-dimensional space.
func (a Coordinate3D) Midpoint(b Coordinate3D) (float64, float64, float64) {
	return Midpoint3D(a, b)
}

// EndPointB returns the coordinate b with given coordinate a and midpoint m.
func (a Coordinate3D) EndPointB(m Coordinate3D) (float64, float64, float64) {
	return EndPoint3D(a, m)
}

// Distance returns the distance between coordinate a and coordinate b.
func Distance(a Coordinate2D, b Coordinate2D) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

// Distance3D returns the distance of coordinates a and b in three-dimensional space.
func Distance3D(a Coordinate3D, b Coordinate3D) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

// EndPoint3D returns the coordinate b with given coordinate a and midpoint m.
func EndPoint3D(a Coordinate3D, m Coordinate3D) (float64, float64, float64) {
	x := (m.X * 2) - a.X
	y := (m.Y * 2) - a.Y
	z := (m.Z * 2) - a.Z

	return x, y, z
}

// Midpoint returns the midpoint of coordinate a and coordinate b.
func Midpoint(a Coordinate2D, b Coordinate2D) (float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2

	return x, y
}

// Midpoint3D returns the midpoint of coordinates a and b in three-dimensional space.
func Midpoint3D(a Coordinate3D, b Coordinate3D) (float64, float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2
	z := (a.Z + b.Z) / 2

	return x, y, z
}

// Slope returns the slope of coordinate a and coordinate b.
func Slope(a Coordinate2D, b Coordinate2D) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}
