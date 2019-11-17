//+build !test

package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const WINDOW_TITLE = "Planetry"
const WINDOW_WIDTH = 800
const WINDOW_HEIGHT = 600

var divisor uint32 = 4
var loopsPerSecond uint32 = 365 / divisor
var nthLoop uint64 = uint64(loopsPerSecond) / 20

func alterSpeed(increase bool) {
	switch increase {
	case true:
		divisor -= 1
		if divisor < 1 {
			divisor = 1
		}
	case false:
		divisor += 1
		if divisor > 365 {
			divisor = 365
		}
	}

	loopsPerSecond = 365 / divisor

	if (loopsPerSecond) < 20 {
		loopsPerSecond = 20
		if increase {
			divisor -= 1
		} else {
			divisor += 1
		}
	}

	nthLoop = uint64(loopsPerSecond) / 20
}

func setUpScene(sceneFilePath string, cameraPosX, cameraPosY, windowWidth, windowHeight int) (*Scene, error) {
	scene, err := loadScene(sceneFilePath)
	if err != nil {
		return nil, err
	}

	camera, err := NewCamera(windowWidth, windowHeight)
	if err != nil {
		return nil, err
	}

	scene.Camera = camera
	return scene, nil
}

func setUpScriptingVM(scriptDir string, scene *Scene) (*Vm, error) {
	scriptingVM := NewVM(scene)

	_, err := scriptingVM.LoadScripts("./testdata")
	if err != nil {
		return nil, err
	}

	return scriptingVM, nil
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(
		WINDOW_TITLE,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		sdl.WINDOW_OPENGL&sdl.WINDOW_RESIZABLE,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	scene, err := setUpScene("./testdata/system.json", 0, 0, WINDOW_WIDTH, WINDOW_HEIGHT)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load scene: %s\n", err)
		return 3
	}

	scriptingVM, err := setUpScriptingVM("./testdata", scene)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start scripting vm: %s\n", err)
		return 4
	}
	defer scriptingVM.Destroy()

	// Loop state
	var event sdl.Event
	running := true
	paused := false
	delta := 0.0
	ticks := uint32(0)

	timer := Timer{}
	timer.start()

	var loop uint64 = 0 // These are also the simulated days
	for running {
		paused = scene.IsPaused()

		// Events
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_SPACE:
					if t.Repeat == 0 {
						scene.SetPaused(!paused)
					}
				case sdl.K_x:
					scene.Camera.ZoomIn()
				case sdl.K_y:
					scene.Camera.ZoomOut()
				case sdl.K_a:
					alterSpeed(false)
				case sdl.K_s:
					alterSpeed(true)
				}
			}
		}

		// Break if destroyed
		if scene.IsDestroyed() {
			running = false
		}

		scene.Simulate(delta)
		if !paused {
			scriptingVM.CallScripts()
		}

		// Draw only every nth loop to save expensive drawing time
		if loop%nthLoop == 0 {
			scene.Draw(renderer, WINDOW_WIDTH, WINDOW_HEIGHT)
			renderer.Present()
		}

		// Sleep the remaining loop time
		delta, ticks = timer.getTime()
		if ticks < 1000/loopsPerSecond {
			sdl.Delay((1000 / loopsPerSecond) - ticks)
		}

		loop = loop + 1
	}

	return 0
}

func main() {
	os.Exit(run())
}
