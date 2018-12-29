package main

import "math"

const TWO_PI = math.Pi * 2

func DrawCircle(renderer Renderer, x, y, radius int, color Color) {
	renderer.SetDrawColor(color.Red, color.Green, color.Blue, color.Alpha)
	
	// Object is too small, draw in default color
	if radius == 0 {
		renderer.DrawPoint(x, y)
		return
	}

	angleInc := 1.0 / float64(radius)
	for angle := 0.0; angle <= TWO_PI; angle += angleInc {
		xpos := float64(x) + float64(radius)*math.Cos(angle)
		ypos := float64(y) + float64(radius)*math.Sin(angle)
		renderer.DrawPoint(int(xpos), int(ypos))
	}
}
