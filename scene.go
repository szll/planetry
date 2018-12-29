package main

import (
	// "fmt"
)

const SIMULATION_STEP = float64(24 * 60 * 60) // one earth day
const MAX_TRACING_POINTS = 1

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
	TargetId				string
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
		x := int(drawableBody.PhysicalBody.Position.X * scale)
		y := int(drawableBody.PhysicalBody.Position.Y * scale)

		if s.TargetId != "" && drawableBody.PhysicalBody.ID == s.TargetId {
			s.Camera.SetToPosition(-x, -y)
		}

		// Tracing path
		// for _, point := range drawableBody.Path {
		// 	x := int(point.X*scale) + s.Camera.x
		// 	y := int(point.Y*scale) + s.Camera.y
		// 	if s.Camera.IsVisible(x, y) {
		// 		DrawCircle(renderer, x, y, 1, *drawableBody.Color)
		// 	}
		// }

		// X and Y realtive to camera
		drawingRadius := int((drawableBody.PhysicalBody.Radius / AU) * scale)

		// if drawableBody.PhysicalBody.ID == "sun" {
		// 	fmt.Println(drawingRadius, drawableBody.PhysicalBody.Radius / AU, scale)
		// }

		DrawCircle(
			renderer,
			x + s.Camera.x,
			y + s.Camera.y,
			drawingRadius,
			*drawableBody.Color,
		)
	}
}

func (s *Scene) GetScale() float64 {
	return float64(s.Camera.zoom) / AU
}

func (s *Scene) Destroy() {
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
