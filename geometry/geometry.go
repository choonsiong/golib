package geometry

import "math"

type CoordinateXY struct {
	X float64
	Y float64
}

type CoordinateXYZ struct {
	X float64
	Y float64
	Z float64
}

func Slope(a CoordinateXY, b CoordinateXY) float64 {
	return (b.Y - a.Y) / (b.X - a.X)
}

func Midpoint(a CoordinateXY, b CoordinateXY) (float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2

	return x, y
}

func Distance3D(a CoordinateXYZ, b CoordinateXYZ) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2) + math.Pow(a.Z-b.Z, 2))
}

func Midpoint3D(a CoordinateXYZ, b CoordinateXYZ) (float64, float64, float64) {
	x := (a.X + b.X) / 2
	y := (a.Y + b.Y) / 2
	z := (a.Z + b.Z) / 2

	return x, y, z
}
