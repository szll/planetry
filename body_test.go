package main

import (
	"testing"
)

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func TestGetAttraction(t *testing.T) {
	var b1 = &Body{
		Name:     "b1",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 0, Y: 0, Z: 0},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	var b2 = &Body{
		Name:     "b2",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 1, Y: 1, Z: 1},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	gForce, err := b1.GetAttraction(b2)
	if err != nil {
		t.Error(err)
	}

	expectedForce := 1.284465784882312e-11

	if !floatEquals(expectedForce, gForce.X) || !floatEquals(expectedForce, gForce.Y) || !floatEquals(expectedForce, gForce.Z) {
		t.Errorf("Each component should have value %f, returned vector has %+v", expectedForce, gForce)
	}
}

func TestGetAttractionError(t *testing.T) {
	var b1 = &Body{
		Name:     "b1",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 0, Y: 0, Z: 0},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	var b2 = &Body{
		Name:     "b2",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 0, Y: 0, Z: 0},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	_, err := b1.GetAttraction(b2)
	if err == nil {
		t.Error("There should have been an error")
	}
	if err.Error() != "Dividing by zero will result in a black hole" {
		t.Error("The error message should be: 'Dividing by zero will result in a black hole'")
	}
}

func TestCollides(t *testing.T) {
	var b1 = &Body{
		Name:     "b1",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 0, Y: 0, Z: 0},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	var b2 = &Body{
		Name:     "b2",
		Mass:     1,
		Radius:   1,
		Position: &Point3D{X: 1, Y: 1, Z: 1},
		Velocity: &Vector3D{X: 0, Y: 0, Z: 0},
	}

	collides := b1.Collides(b2)
	if collides {
		t.Error("Should have returned false")
	}
}
