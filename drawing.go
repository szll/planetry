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

// DrawGrid needs information about the window as well as the zoom/scale
// It should only draw max five grid lines horizontal and vertical
func DrawGrid(renderer Renderer, scale float64, windowWidth, windowHeight int) (int) {
	renderer.SetDrawColor(33, 33, 33, 255)

	halfWidth := windowWidth / 2
	halfHeight := windowHeight / 2
	distance := int(AU * scale)

	if distance == 0 { return 0 }

	// Calculate count of lines (this has to be done only once for a set scale)
	lines := 0
	lastHit := 0
	lastCount := 0
	for i := 1; i <= 1000; i += 2 {
		count := halfWidth / (distance * i)

		if count > 2 && count <= 4 { // Min and max count of lines
			lastHit = i
			lastCount = count
		}

		// Early exit
		if count == 0 && i > 1 {
			distance = distance * lastHit
			if distance == 0 {
				break
			}

			lines = lastCount
			break
		}
	}

	// Draw grid
	for i := 0; i < windowHeight; i++ {
		renderer.DrawPoint(halfWidth, i)

		for j := 1; j <= lines; j++ {
			renderer.DrawPoint(halfWidth + distance * j, i)
			renderer.DrawPoint(halfWidth - distance * j, i)
		}
	}
	for i := 0; i < windowWidth; i++ {
		renderer.DrawPoint(i, halfHeight)

		for j := 1; j <= lines; j++ {
			renderer.DrawPoint(i, halfHeight + distance * j)
			renderer.DrawPoint(i, halfHeight - distance * j)
		}
	}

	return lastHit
}
