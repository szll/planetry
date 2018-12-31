![planetry logo](./docs/logo-01.png "planetry logo")

# Planetry

[![Build Status](https://travis-ci.org/szll/planetry.svg?branch=master)](https://travis-ci.org/szll/planetry)
[![Coverage Status](https://coveralls.io/repos/github/szll/planetry/badge.svg?branch=master)](https://coveralls.io/github/szll/planetry?branch=master)

**ATTENTION:** this project is totally **WORK IN PROGRESS**

Planetry is a simple app for gravitational simulations of objects in space. You can specify the objects or bodys by your own. A scripting system allows you to check for specific events.

For example: you simulate the solar system with the sun and all its planets and add a second star that passes through the solar system. A simple script, written in Lua, can check on each simulation step, when or if Earth moves outside of the habitable zone. If the event occurs it can pause the simulation and display additional information.

![planetry](./docs/scene.gif "planetry")

## Installation

### External Dependencies

This project needs [SDL2](https://www.libsdl.org) and [Lua 5.1](https://www.lua.org/manual/5.1/) as requirements which have to be installed first.

#### Mac

First install the requirements via [homebrew](https://brew.sh): `brew install dep sdl2 lua@5.1`

#### Linux (tested on Ubuntu)

Just type: `./install-dependencies-linux.sh`. This may take some time as SDL gets build from source within the script.

#### Windows

Todo ... I would be more than happy to get help here :D

### Golang dependencies and building

The golang dependencies are managed via [dep](https://github.com/golang/dep). To get the dependencies type: `dep ensure`

To build the executable type: `go build`

## Usage

After the executable was build, you can start the simulation by typing:

```
./planetry
```

> **NOTE**: Planetry loads all the files in `testdata/`. Later you can create projects and load these, but currently if you want to change something in the simulation, just alter the files in `testdata/`.

Feel free to add more Lua scripts or edit `testdata/system.json`.

On runtime, you can use:
 - `y` to zoom out,
 - `x` to zoom in and
 - `SPACE` to pause/unpause

### Describing the simulated scene

See [docs/SCENE.md](docs/SCENE.md).

### Lua scripting

See [docs/SCRIPTING.md](docs/SCRIPTING.md).

## TODO

As you may noticed: this project is in a really early stage of development. Here's the stuff that has to be done next:

- [x] Attach camera to object
- [x] Draw bodies in their real size; not just pixels
- [x] Scripts per scene description
- [ ] Create project by command
- [ ] Live time / event script support like on pause, on start up, ...
- [ ] Draw name on pause
- [ ] Collision
- [ ] Alter speed
- [ ] Runtime memory functions for scripting
- [ ] 3D rendering

## Know Issues

- Window appears but no objects were drawn. After switching to MacOS Mojave I had to update the go dependencies as well as SDL on the system.
