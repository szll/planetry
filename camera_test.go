package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCamera(t *testing.T) {
	windowWidth := 10
	windowHeight := 10
	c, err := NewCamera(windowWidth, windowHeight)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, c.windowWidth, windowWidth, "c.windowWidth should be equal to windowWidth")
	assert.Equal(t, c.windowHeight, windowHeight, "c.windowHeight should be equal to windowHeight")
	assert.Equal(t, c.x, windowWidth/2, "x should be half windowWidth")
	assert.Equal(t, c.y, windowHeight/2, "y should be half windowWidth")
}

func TestNewCameraError(t *testing.T) {
	windowWidth := -1
	windowHeight := 10
	_, err := NewCamera(windowWidth, windowHeight)

	assert.Equal(t, err.Error(), "windowWidth or windowHeight cannot be < 1", "should return the right error")
}

func TestSetToPosition(t *testing.T) {
	windowWidth := 10
	windowHeight := 10
	c, err := NewCamera(windowWidth, windowHeight)
	if err != nil {
		t.Error(err)
	}

	c.SetToPosition(10, 10)

	assert.Equal(t, c.x, windowWidth/2+10, "x should be 15")
	assert.Equal(t, c.y, windowHeight/2+10, "y should be 15")
}

func TestIsVisible(t *testing.T) {
	windowWidth := 10
	windowHeight := 10

	x1 := 0
	y1 := 0
	x2 := 16
	y2 := -21

	c, err := NewCamera(windowWidth, windowHeight)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, c.IsVisible(x1, y1), true, "(x1, y1) should be visible")
	assert.Equal(t, c.IsVisible(x2, y2), false, "(x2, y2) should be invisible")
}

func TestZoom(t *testing.T) {
	windowWidth := 10
	windowHeight := 10

	c, err := NewCamera(windowWidth, windowHeight)
	if err != nil {
		t.Error(err)
	}

	c.Zoom(-1000)
	assert.Equal(t, c.zoom, int16(1), "zoom should be 0")

	c.Zoom(1000)
	assert.Equal(t, c.zoom, int16(200), "zoom should be 200")

	c.Zoom(-1)
	assert.Equal(t, c.zoom, int16(199), "zoom should be 199")
}
