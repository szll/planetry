package main

import "math"

type Vector3D struct {
	X float64
	Y float64
	Z float64
}

func (v *Vector3D) NormalizeAndFactorize(factor float64) *Vector3D {
	length := v.Length()

	return &Vector3D{
		(v.X / length) * factor,
		(v.Y / length) * factor,
		(v.Z / length) * factor,
	}
}

func (v *Vector3D) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3D) Add(other *Vector3D) {
	v.X = v.X + other.X
	v.Y = v.Y + other.Y
	v.Z = v.Z + other.Z
}
