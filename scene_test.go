package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestingScene() *Scene {
	db1 := DrawableBody{
		PhysicalBody: &Body{
			Name:   "n1",
			Mass:   1,
			Radius: 1,
			Position: &Point3D{
				X: 1,
				Y: 0,
				Z: 0,
			},
			Velocity: &Vector3D{
				X: 1,
				Y: 0,
				Z: 0,
			},
		},
		Path: PointQueue{},
		Color: &Color{
			Red:   1,
			Green: 2,
			Blue:  3,
			Alpha: 4,
		},
	}

	db2 := DrawableBody{
		PhysicalBody: &Body{
			Name:   "n2",
			Mass:   1,
			Radius: 1,
			Position: &Point3D{
				X: -1,
				Y: 0,
				Z: 0,
			},
			Velocity: &Vector3D{
				X: 1,
				Y: 0,
				Z: 0,
			},
		},
		Path: PointQueue{},
		Color: &Color{
			Red:   1,
			Green: 2,
			Blue:  3,
			Alpha: 4,
		},
	}

	c, _ := NewCamera(10, 10)

	return &Scene{
		Bodies:          []*DrawableBody{&db1, &db2},
		ForcesOfBodies:  map[*DrawableBody]Vector3D{},
		Camera:          c,
		BackgroundColor: &Color{},
		zoom:            10,
		destroyed:       false,
		simulations:     0,
		paused:          false,
	}
}

type MockRenderer struct{}

func (m *MockRenderer) SetDrawColor(r, g, b, a uint8) error {
	return nil
}

func (m *MockRenderer) Clear() error {
	return nil
}

func (m *MockRenderer) DrawPoint(x, y int) error {
	return nil
}

func TestSimulateScene(t *testing.T) {
	s := createTestingScene()
	s.Simulate(1)

	// TODO: check the values
	// Check for simulations
	assert.Equal(t, s.simulations, int64(1), "s.simulations should be 1")

	// Check for tracing path (one point each)
	assert.Equal(t, len(s.Bodies[0].Path), 1, "len(s.Bodies[0].Path) should be 1")
	assert.Equal(t, len(s.Bodies[1].Path), 1, "len(s.Bodies[1].Path) should be 1")

	// Check for forces of bodies (two entries)
	assert.Equal(t, len(s.ForcesOfBodies), 2, "len(s.ForcesOfBodies) should be 2")
}

func TestSimulateScenePaused(t *testing.T) {
	s := Scene{}
	s.paused = true

	// That should not crash, because simulation returns early
	assert.Nil(t, s.Simulate(1), "should return nil")
}

func TestSimulateSceneErrorGetAttraction(t *testing.T) {
	s := createTestingScene()
	s.Bodies[1].PhysicalBody.Position.X = 1

	err := s.Simulate(1)

	assert.EqualError(t, err, "Dividing by zero will result in a black hole", "should return devision by zero error")
}

func TestSimulateSceneRemovePointsFromPath(t *testing.T) {
	s := createTestingScene()
	for i := 0; i < 100; i++ {
		s.Bodies[0].Path.Push(Point3D{})
	}

	s.Simulate(1)

	assert.Equal(t, len(s.Bodies[0].Path), 50, "path queue should have 50 entries")
}

func TestDraw(t *testing.T) {
	s := createTestingScene()
	s.Simulate(1)

	// This should not fail; TODO: I know it's poor testing at this point ...
	s.Draw(&MockRenderer{})
}

func TestZoom(t *testing.T) {
	s := createTestingScene()
	s.Zoom(-1000)
	assert.Equal(t, s.zoom, int16(1), "zoom shoud be 0")

	s.Zoom(1000)
	assert.Equal(t, s.zoom, int16(200), "zoom shoud be 200")

	s.Zoom(-1)
	assert.Equal(t, s.zoom, int16(199), "zoom shoud be 199")
}

func TestDestroy(t *testing.T) {
	// TODO: implement
}

func TestIsDestroyed(t *testing.T) {
	// TODO: implement
}

func TestGetSimulations(t *testing.T) {
	// TODO: implement
}

func TestIsPaused(t *testing.T) {
	// TODO: implement
}

func TestSetPaused(t *testing.T) {
	// TODO: implement
}

func TestGetBodyByName(t *testing.T) {
	// TODO: implement
}

func TestGsetVMMethodes(t *testing.T) {
	// TODO: implement
}
