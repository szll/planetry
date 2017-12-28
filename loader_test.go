package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/antonholmquist/jason"
)

func TestLoadColor(t *testing.T) {
	color := `
		{
			"red": 1,
			"green": 2,
			"blue": 3,
			"alpha": 4
		}
	`

	cObj, err := jason.NewObjectFromBytes([]byte(color))
	if err != nil {
		t.Error(err)
	}

	c, err := loadColor(cObj)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, c.Red, uint8(1), "red should be 1")
	assert.Equal(t, c.Green, uint8(2), "green should be 2")
	assert.Equal(t, c.Blue, uint8(3), "blue should be 3")
	assert.Equal(t, c.Alpha, uint8(4), "alpha should be 4")
}

func TestLoadColorRedError(t *testing.T) {
	color := `
		{
			"green": 2,
			"blue": 3,
			"alpha": 4
		}
	`

	cObj, err := jason.NewObjectFromBytes([]byte(color))
	if err != nil {
		t.Error(err)
	}

	_, err = loadColor(cObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadColorGreenError(t *testing.T) {
	color := `
		{
			"red": 1,
			"blue": 3,
			"alpha": 4
		}
	`

	cObj, err := jason.NewObjectFromBytes([]byte(color))
	if err != nil {
		t.Error(err)
	}

	_, err = loadColor(cObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadColorBlueError(t *testing.T) {
	color := `
		{
			"red": 1,
			"green": 2,
			"alpha": 4
		}
	`

	cObj, err := jason.NewObjectFromBytes([]byte(color))
	if err != nil {
		t.Error(err)
	}

	_, err = loadColor(cObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadColorAlphaError(t *testing.T) {
	color := `
		{
			"red": 1,
			"green": 2,
			"blue": 3
		}
	`

	cObj, err := jason.NewObjectFromBytes([]byte(color))
	if err != nil {
		t.Error(err)
	}

	_, err = loadColor(cObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}
