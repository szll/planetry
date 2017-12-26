package main

type Point3D struct {
	X float64
	Y float64
	Z float64
}

func (p *Point3D) VectorBetween(o *Point3D) *Vector3D {
	return &Vector3D{
		o.X - p.X,
		o.Y - p.Y,
		o.Z - p.Z,
	}
}
