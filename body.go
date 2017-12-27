package main

import (
	"errors"
)

type Body struct {
	Name     string
	Mass     float64
	Radius   float64
	Position *Point3D
	Velocity *Vector3D
}

// GetAttraction calculates the grafitational force between two bodies
func (b *Body) GetAttraction(o *Body) (*Vector3D, error) {
	v := b.Position.VectorBetween(o.Position)
	l := v.Length() // Also distance between bodies

	if l == 0 {
		return nil, errors.New("Dividing by zero will result in a black hole")
	}

	fg := G * b.Mass * o.Mass / (l * l)

	return v.NormalizeAndFactorize(fg), nil
}

func (b *Body) Collides(o *Body) bool {
	return false
}
