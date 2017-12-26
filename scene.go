package main

import (
	"github.com/stevedonovan/luar"
	"github.com/veandco/go-sdl2/sdl"
)

const SIMULATION_STEP = float64(24 * 60 * 60)
const MAX_TRACING_POINTS = 50

type DrawableBody struct {
	PhysicalBody *Body
	Path         PointQueue
	Color        Color
}

type Scene struct {
	Bodies          []*DrawableBody
	ForcesOfBodies  map[*DrawableBody]Vector3D
	Camera          *Camera
	BackgroundColor Color
	zoom            int16
	destroyed       bool
	simulations     int64
	paused          bool
}

func (s *Scene) Simulate(delta float64) {
	if s.paused {
		return
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
				panic(err)
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
}

func (s *Scene) Draw(renderer *sdl.Renderer) {
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
			x := int(point.X*scale) - s.Camera.x
			y := int(point.Y*scale) - s.Camera.y
			if s.Camera.IsVisible(x, y) {
				renderer.DrawPoint(x, y)
			}
		}
	}
}

func (s *Scene) GetScale() float64 {
	return float64(s.zoom) / AU
}

func (s *Scene) Zoom(amount int16) {
	s.zoom = s.zoom + amount
	if s.zoom <= 0 {
		s.zoom = 1
	}
	if s.zoom >= 200 {
		s.zoom = 1
	}
}

func (s *Scene) destroy() {
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

func (s *Scene) GetBodyByName(name string) *Body {
	for _, dBody := range s.Bodies {
		if dBody.PhysicalBody.Name == name {
			return dBody.PhysicalBody
		}
	}
	return nil
}

func (s *Scene) getVMMethodes() luar.Map {
	return luar.Map{
		"AU":            AU,
		"getBodyByName": s.GetBodyByName,
		"getSteps":      s.GetSimulations,
		"setPaused":     s.SetPaused,
	}
}

// void draw_circle(SDL_Point center, int radius, SDL_Color color)
// {
//     SDL_SetRenderDrawColor(renderer, color.r, color.g, color.b, color.a);
//     for (int w = 0; w < radius * 2; w++)
//     {
//         for (int h = 0; h < radius * 2; h++)
//         {
//             int dx = radius - w; // horizontal offset
//             int dy = radius - h; // vertical offset
//             if ((dx*dx + dy*dy) <= (radius * radius))
//             {
//                 SDL_RenderDrawPoint(renderer, center.x + dx, center.y + dy);
//             }
//         }
//     }
// }

// TODO: AddBodyToScene()
// TODO: RemoveBodyFromScene()
// TODO: Collision
