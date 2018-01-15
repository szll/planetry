package main

import (
	"github.com/stevedonovan/luar"
)

const SIMULATION_STEP = float64(24 * 60 * 60)
const MAX_TRACING_POINTS = 50

type Renderer interface {
	SetDrawColor(r, g, b, a uint8) error
	Clear() error
	DrawPoint(x, y int) error
}

type DrawableBody struct {
	PhysicalBody *Body
	Path         PointQueue
	Color        *Color
}

type Scene struct {
	Bodies          []*DrawableBody
	ForcesOfBodies  map[*DrawableBody]Vector3D
	Camera          *Camera
	BackgroundColor *Color
	destroyed       bool
	simulations     int64
	paused          bool
}

func (s *Scene) Simulate(delta float64) error {
	if s.paused || s.destroyed {
		return nil
	}

	step := SIMULATION_STEP

	for _, drawableBody := range s.Bodies {
		totalForce := Vector3D{X: 0.0, Y: 0.0, Z: 0.0}

		for _, other := range s.Bodies {
			if other == drawableBody {
				continue
			}

			force, err := drawableBody.PhysicalBody.GetAttraction(other.PhysicalBody)
			if err != nil {
				return err
			}

			totalForce.Add(force)
		}

		s.ForcesOfBodies[drawableBody] = totalForce
	}

	for _, drawableBody := range s.Bodies {
		force := s.ForcesOfBodies[drawableBody]
		// F * t = m * v
		drawableBody.PhysicalBody.Velocity.X = drawableBody.PhysicalBody.Velocity.X + force.X/drawableBody.PhysicalBody.Mass*step
		drawableBody.PhysicalBody.Velocity.Y = drawableBody.PhysicalBody.Velocity.Y + force.Y/drawableBody.PhysicalBody.Mass*step
		drawableBody.PhysicalBody.Velocity.Z = drawableBody.PhysicalBody.Velocity.Z + force.Z/drawableBody.PhysicalBody.Mass*step
		// s = v * t
		drawableBody.PhysicalBody.Position.X = drawableBody.PhysicalBody.Position.X + drawableBody.PhysicalBody.Velocity.X*step
		drawableBody.PhysicalBody.Position.Y = drawableBody.PhysicalBody.Position.Y + drawableBody.PhysicalBody.Velocity.Y*step
		drawableBody.PhysicalBody.Position.Z = drawableBody.PhysicalBody.Position.Z + drawableBody.PhysicalBody.Velocity.Z*step

		// Add tracing point
		drawableBody.Path.Push(Point3D{
			drawableBody.PhysicalBody.Position.X,
			drawableBody.PhysicalBody.Position.Y,
			drawableBody.PhysicalBody.Position.Z,
		})

		// Delete old tracing points
		currenPathLength := len(drawableBody.Path)
		if currenPathLength > MAX_TRACING_POINTS {
			toBeDeleted := currenPathLength - MAX_TRACING_POINTS
			for i := 0; i < toBeDeleted; i++ {
				drawableBody.Path.Pop()
			}
		}
	}

	s.simulations = s.simulations + 1
	return nil
}

func (s *Scene) Draw(renderer Renderer) {
	renderer.SetDrawColor(
		s.BackgroundColor.Red,
		s.BackgroundColor.Green,
		s.BackgroundColor.Blue,
		s.BackgroundColor.Alpha)
	renderer.Clear()

	scale := s.GetScale()

	for _, drawableBody := range s.Bodies {
		renderer.SetDrawColor(
			drawableBody.Color.Red,
			drawableBody.Color.Green,
			drawableBody.Color.Blue,
			drawableBody.Color.Alpha,
		)
		for _, point := range drawableBody.Path {
			x := int(point.X*scale) + s.Camera.x
			y := int(point.Y*scale) + s.Camera.y
			if s.Camera.IsVisible(x, y) {
				renderer.DrawPoint(x, y)
			}
		}
	}
}

func (s *Scene) GetScale() float64 {
	return float64(s.Camera.zoom) / AU
}

func (s *Scene) Destroy() {
	// TODO: free resources / nil refrences
	s.destroyed = true
}

func (s *Scene) IsDestroyed() bool {
	return s.destroyed
}

func (s *Scene) GetSimulations() int64 {
	return s.simulations
}

func (s *Scene) IsPaused() bool {
	return s.paused
}

func (s *Scene) SetPaused(paused bool) {
	s.paused = paused
}

// Following functions are available in lua scope

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
		"getBodyByName":    s.GetBodyByName,
		"getSteps":         s.GetSimulations,
		"setPaused":        s.SetPaused,
		"createPoint3D":    CreatePoint3D,
		"createVector3D":   CreateVector3D,
		"createBody":       CreateBody,
		"addBodyToScene":   s.AddBodyToScene,
		"removeBodyByName": s.RemoveBodyByName,
	}
}
