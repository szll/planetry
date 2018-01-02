package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestQueue(t *testing.T) {
	p := Point3D{X: 1, Y: 1, Z: 1}
	q := PointQueue{}

	q.Push(p)

	assert.Equal(t, len(q), 1, "queue should have one element")

	p2 := q.Pop()

	assert.Exactly(t, p, p2, "queue should have popped the inserted point")
	assert.Equal(t, len(q), 0, "length of queue should be 0")
}
