package main

import "testing"

func TestDrawCircle(t *testing.T) {
	// This should not fail; TODO: I know it's poor testing at this point ...
	DrawCircle(&MockRenderer{}, 10, 10, 5, Color{Red: 1, Green: 2, Blue: 3, Alpha: 4})
}

func TestDrawGrid(t *testing.T) {
	// This should not fail; TODO: I know it's poor testing at this point ...
	DrawGrid(&MockRenderer{}, 1.0, 10, 10)
	DrawGrid(&MockRenderer{}, 0, 10, 10)
}
