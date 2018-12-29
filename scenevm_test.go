package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetBodyById(t *testing.T) {
	s := createTestingScene()
	b := s.GetBodyById("n1id")

	assert.Exactly(t, b, s.Bodies[0].PhysicalBody, "b should be s.Bodies[0].PhysicalBody")

	assert.Nil(t, s.GetBodyById("x1id"))
}

func TestGetBodyByName(t *testing.T) {
	s := createTestingScene()
	b := s.GetBodyByName("n1")

	assert.Exactly(t, b, s.Bodies[0].PhysicalBody, "b should be s.Bodies[0].PhysicalBody")

	assert.Nil(t, s.GetBodyByName("x1"))
}

func TestCreatePoint3D(t *testing.T) {
	p := CreatePoint3D(1, 2, 3)
	assert.Equal(t, p, &Point3D{X: 1, Y: 2, Z: 3}, "points should be equal")
}

func TestCreateVector3D(t *testing.T) {
	p := CreateVector3D(1, 2, 3)
	assert.Equal(t, p, &Vector3D{X: 1, Y: 2, Z: 3}, "vectors should be equal")
}

func TestCreateBody(t *testing.T) {
	p := &Point3D{X: 1, Y: 2, Z: 3}
	v := &Vector3D{X: 1, Y: 2, Z: 3}

	b := CreateBody("name", 1, 1, p, v)
	b2 := &Body{
		Name:     "name",
		Mass:     1,
		Radius:   1,
		Position: p,
		Velocity: v,
	}

	assert.Equal(t, b, b2, "bodies should be equal")
}

func TestAddBodyToScene(t *testing.T) {
	s := createTestingScene()

	p := &Point3D{X: 1, Y: 2, Z: 3}
	v := &Vector3D{X: 1, Y: 2, Z: 3}
	b := CreateBody("name", 1, 1, p, v)

	s.AddBodyToScene(b, 10, 10, 10, 10)

	b2 := s.GetBodyByName("name")

	assert.Exactly(t, b, b2, "bodies should be exactly the same")
}

func TestRemoveBodyFromScene(t *testing.T) {
	s := createTestingScene()
	s.RemoveBodyByName("n1")

	assert.Nil(t, s.GetBodyByName("n1"), "getBodyByName should return nil")
}

func TestRemoveBodyFromSceneById(t *testing.T) {
	s := createTestingScene()
	s.RemoveBodyById("n1id")

	assert.Nil(t, s.GetBodyById("n1id"), "getBodyById should return nil")
}


func TestGetVMMethodes(t *testing.T) {
	s := createTestingScene()
	m := s.getVMMethodes()

	assert.Equal(t, len(m), 11, "getVMMethodes should return map containing 10 entries")
}