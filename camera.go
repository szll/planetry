package main

import "errors"

type Camera struct {
	x            int
	y            int
	windowWidth  int
	windowHeight int
	zoom         float64
}

func NewCamera(windowWidth, windowHeight int) (*Camera, error) {
	if windowWidth < 1 || windowHeight < 1 {
		return nil, errors.New("windowWidth or windowHeight cannot be < 1")
	}
	return &Camera{windowWidth / 2, windowHeight / 2, windowWidth, windowHeight, 8}, nil
}

func (c *Camera) SetToPosition(x, y int) {
	c.x = c.windowWidth / 2 + x
	c.y = c.windowHeight / 2 + y
}

func (c *Camera) IsVisible(x, y int) bool {
	lowWidth := x < c.windowWidth / 2 - c.x
	hightWidth := x > c.windowWidth / 2 + c.x
	lowHeight := y < c.windowHeight / 2 - c.y
	highWidth := y > c.windowHeight / 2 + c.y
	return !(lowWidth || hightWidth || lowHeight || highWidth)
}

func (c *Camera) ZoomIn() {
	c.zoom *= 2
}

func (c *Camera) ZoomOut() {
	c.zoom /= 2
	if c.zoom < 1 {
		c.zoom = 1
	}
}