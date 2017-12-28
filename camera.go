package main

import "errors"

type Camera struct {
	x            int
	y            int
	windowWidth  int
	windowHeight int
}

func NewCamera(windowWidth, windowHeight int) (*Camera, error) {
	if windowWidth < 1 || windowHeight < 1 {
		return nil, errors.New("windowWidth or windowHeight cannot be < 1")
	}
	return &Camera{windowWidth / 2, windowHeight / 2, windowWidth, windowHeight}, nil
}

func (c *Camera) SetToPosition(x int, y int) {
	c.x = x + c.windowWidth/2
	c.y = y + c.windowHeight/2
}

func (c *Camera) IsVisible(x, y int) bool {
	return !(x < 0 || x > c.windowWidth || y < 0 || y > c.windowHeight)
}
