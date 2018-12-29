package main

import (
	"github.com/stevedonovan/luar"
)

// Following functions are available in lua scope

func (s *Scene) GetBodyById(id string) *Body {
	for _, dBody := range s.Bodies {
		if dBody.PhysicalBody.ID == id {
			return dBody.PhysicalBody
		}
	}
	return nil
}

func (s *Scene) GetBodyByName(name string) *Body {
	for _, dBody := range s.Bodies {
		if dBody.PhysicalBody.Name == name {
			return dBody.PhysicalBody
		}
	}
	return nil
}

func CreatePoint3D(x, y, z float64) *Point3D {
	return &Point3D{X: x, Y: y, Z: z}
}

func CreateVector3D(x, y, z float64) *Vector3D {
	return &Vector3D{X: x, Y: y, Z: z}
}

func CreateBody(name string, mass, radius float64, position *Point3D, velocity *Vector3D) *Body {
	return &Body{
		Name:     name,
		Mass:     mass,
		Radius:   radius,
		Position: position,
		Velocity: velocity,
	}
}

func (s *Scene) AddBodyToScene(body *Body, red, green, blue, alpha uint8) {
	s.Bodies = append(s.Bodies, &DrawableBody{
		PhysicalBody: body,
		Path:         PointQueue{},
		Color:        &Color{Red: red, Green: green, Blue: blue, Alpha: alpha},
	})
}

func (s *Scene) RemoveBodyById(id string) {
	index := -1
	for i, dBody := range s.Bodies {
		if dBody.PhysicalBody.ID == id {
			index = i
			break
		}
	}

	if index >= 0 {
		s.Bodies = append(s.Bodies[:index], s.Bodies[index+1:]...)
	}
}

func (s *Scene) RemoveBodyByName(name string) {
	index := -1
	for i, dBody := range s.Bodies {
		if dBody.PhysicalBody.Name == name {
			index = i
			break
		}
	}

	if index >= 0 {
		s.Bodies = append(s.Bodies[:index], s.Bodies[index+1:]...)
	}
}

func (s *Scene) getVMMethodes() luar.Map {
	return luar.Map{
		"AU":               AU,
		"getBodyById":      s.GetBodyById,
		"getBodyByName":    s.GetBodyByName,		// DEPRECATED
		"getSteps":         s.GetSimulations,
		"setPaused":        s.SetPaused,
		"createPoint3D":    CreatePoint3D,
		"createVector3D":   CreateVector3D,
		"createBody":       CreateBody,
		"addBodyToScene":   s.AddBodyToScene,
		"removeBodyByName": s.RemoveBodyByName, // DEPRECATED
		"removeBodyById":   s.RemoveBodyById,
	}
}