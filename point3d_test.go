package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorBetween(t *testing.T) {
	p := Point3D{X: 1, Y: 2, Z: 3}
	p2 := Point3D{X: 1, Y: 2, Z: 3}
	v := p.VectorBetween(&p2)

	assert.Equal(t, v.X, 0.0, "X should be 0.0")
	assert.Equal(t, v.Y, 0.0, "Y should be 0.0")
	assert.Equal(t, v.Z, 0.0, "Z should be 0.0")
}
