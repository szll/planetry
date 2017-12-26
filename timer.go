package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Timer struct {
	timeMS      uint32
	oldTimeMS   uint32
	frameTimeMS uint32
}

func (t *Timer) start() {
	t.timeMS = sdl.GetTicks()
}

func (t *Timer) getTime() (float64, uint32) {
	t.oldTimeMS = t.timeMS
	t.timeMS = sdl.GetTicks()
	t.frameTimeMS = t.timeMS - t.oldTimeMS
	return float64(t.frameTimeMS) / 1000.0, t.frameTimeMS
}
