package main

import "testing"
import "time"
import "github.com/stretchr/testify/assert"

func TestTimer(t *testing.T) {
	timer := Timer{}

	timer.start()

	time.Sleep(time.Second)

	f, u := timer.getTime()

	assert.True(t, f > 0, "f should be > 0")
	assert.True(t, u > 0, "u should be > 0")
}
