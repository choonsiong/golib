package geometry

import "math"

// CoordinateXY represents a coordinate in two-dimensional space.
type CoordinateXY struct {
	X float64
	Y float64
}

// CoordinateXYZ represents a coordinate in three-dimensional space.
type CoordinateXYZ struct {
	X float64
	Y float64
	Z float64
}

// Slope returns the slope of endpoints A and B.
func Slope(a CoordinateXY, b CoordinateXY) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}

// Midpoint returns the midpoint of endpoints A and B.
func Midpoint(a CoordinateXY, b CoordinateXY) (float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2

	return x, y
}

// Distance3D returns the distance of endpoints A and B in three-dimensional space.
func Distance3D(a CoordinateXYZ, b CoordinateXYZ) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

// Midpoint3D returns the midpoint of endpoints A and B in three-dimensional space.
func Midpoint3D(a CoordinateXYZ, b CoordinateXYZ) (float64, float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2
	z := (a.Z + b.Z) / 2

	return x, y, z
}

// CoordinateForMidPoint3D returns the endpoint B with given endpoint A and midpoint M.
func CoordinateForMidPoint3D(a CoordinateXYZ, m CoordinateXYZ) (float64, float64, float64) {
	x := (m.X * 2) - a.X
	y := (m.Y * 2) - a.Y
	z := (m.Z * 2) - a.Z

	return x, y, z
}
