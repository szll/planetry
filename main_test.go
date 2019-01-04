package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlterSpeed(t *testing.T) {
	alterSpeed(true)
	assert.Equal(t, uint32(0x5), speed)
	assert.Equal(t, uint32(0x49), loopsPerSecond)
	assert.Equal(t, uint64(0x3), nthLoop)

	alterSpeed(false)
	assert.Equal(t, uint32(0x4), speed)
	assert.Equal(t, uint32(0x5b), loopsPerSecond)
	assert.Equal(t, uint64(0x4), nthLoop)
}
