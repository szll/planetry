package main

import "math"

func DrawCircle(renderer Renderer, x, y, radius int, color Color) {
	twoPi := math.Pi * 2

	renderer.SetDrawColor(color.Red, color.Green, color.Blue, color.Alpha)

	angleInc := 1.0 / float64(radius)
	for angle := 0.0; angle <= twoPi; angle += angleInc {
		xpos := float64(x) + float64(radius)*math.Cos(angle)
		ypos := float64(y) + float64(radius)*math.Sin(angle)
		renderer.DrawPoint(int(xpos), int(ypos))
	}
}
