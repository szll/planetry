package main

import "math"

const TWO_PI = math.Pi * 2

// This color is used in case an object is too small to display
var defaultColor = Color{Red: 127, Green: 127, Blue: 127, Alpha: 255}

func DrawCircle(renderer Renderer, x, y, radius int, color Color) {
	// Object is too small, draw in default color
	if radius == 0 {
		renderer.SetDrawColor(defaultColor.Red, defaultColor.Green, defaultColor.Blue, defaultColor.Alpha)
		renderer.DrawPoint(int(x), int(y))
		return
	}

	renderer.SetDrawColor(color.Red, color.Green, color.Blue, color.Alpha)

	angleInc := 1.0 / float64(radius)
	for angle := 0.0; angle <= TWO_PI; angle += angleInc {
		xpos := float64(x) + float64(radius)*math.Cos(angle)
		ypos := float64(y) + float64(radius)*math.Sin(angle)
		renderer.DrawPoint(int(xpos), int(ypos))
	}
}
