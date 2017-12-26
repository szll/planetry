package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/antonholmquist/jason"
)

// loadScene creates a new scene from a json file
func loadScene(path string) (*Scene, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	v, err := jason.NewObjectFromBytes(file)
	if err != nil {
		return nil, err
	}

	jBackgroundColor, err := v.GetObject("backgroundColor")
	if err != nil {
		return nil, err
	}

	redBackgroundColorInt, err := jBackgroundColor.GetInt64("red")
	if err != nil {
		return nil, err
	}

	greenBackgroundColorInt, err := jBackgroundColor.GetInt64("green")
	if err != nil {
		return nil, err
	}

	blueBackgroundColorInt, err := jBackgroundColor.GetInt64("blue")
	if err != nil {
		return nil, err
	}

	alphaBackgroundColorInt, err := jBackgroundColor.GetInt64("alpha")
	if err != nil {
		return nil, err
	}

	jBodies, err := v.GetObjectArray("bodies")
	if err != nil {
		return nil, err
	}

	var dbs []*DrawableBody
	for _, jBody := range jBodies {
		name, err := jBody.GetString("name")
		if err != nil {
			return nil, err
		}

		mass, err := jBody.GetFloat64("mass")
		if err != nil {
			return nil, err
		}

		radius, err := jBody.GetFloat64("radius")
		if err != nil {
			return nil, err
		}

		jPosition, err := jBody.GetObject("position")
		if err != nil {
			return nil, err
		}

		pX, err := jPosition.GetFloat64("x")
		if err != nil {
			return nil, err
		}

		pY, err := jPosition.GetFloat64("y")
		if err != nil {
			return nil, err
		}

		pZ, err := jPosition.GetFloat64("z")
		if err != nil {
			return nil, err
		}

		jVelocity, err := jBody.GetObject("velocity")
		if err != nil {
			return nil, err
		}

		vX, err := jVelocity.GetFloat64("x")
		if err != nil {
			return nil, err
		}

		vY, err := jVelocity.GetFloat64("y")
		if err != nil {
			return nil, err
		}

		vZ, err := jVelocity.GetFloat64("z")
		if err != nil {
			return nil, err
		}

		jColor, err := jBody.GetObject("color")
		if err != nil {
			return nil, err
		}

		redInt, err := jColor.GetInt64("red")
		if err != nil {
			return nil, err
		}

		greenInt, err := jColor.GetInt64("green")
		if err != nil {
			return nil, err
		}

		blueInt, err := jColor.GetInt64("blue")
		if err != nil {
			return nil, err
		}

		alphaInt, err := jColor.GetInt64("alpha")
		if err != nil {
			return nil, err
		}

		db := &DrawableBody{
			PhysicalBody: &Body{
				Name:   name,
				Mass:   mass,
				Radius: radius,
				Position: &Point3D{
					X: pX * AU,
					Y: pY * AU,
					Z: pZ * AU,
				},
				Velocity: &Vector3D{
					X: vX * 1000,
					Y: vY * 1000,
					Z: vZ * 1000,
				},
			},
			Path: PointQueue{},
			Color: Color{
				Red:   uint8(redInt),
				Green: uint8(greenInt),
				Blue:  uint8(blueInt),
				Alpha: uint8(alphaInt),
			},
		}

		dbs = append(dbs, db)
	}

	return &Scene{
		Bodies:         dbs,
		ForcesOfBodies: map[*DrawableBody]Vector3D{},
		BackgroundColor: Color{
			Red:   uint8(redBackgroundColorInt),
			Green: uint8(greenBackgroundColorInt),
			Blue:  uint8(blueBackgroundColorInt),
			Alpha: uint8(alphaBackgroundColorInt),
		},
		Camera:    nil,
		zoom:      10, // TODO: initial state from json
		destroyed: false,
	}, nil
}

// loadAllLuaFiles loads all Lua script files in a directory
func loadAllLuaFiles(dir string) ([]string, error) {
	fi, err := os.Stat(dir)
	if err != nil {
		return []string{}, err
	}

	if !fi.Mode().IsDir() {
		return []string{}, errors.New(fmt.Sprintf("%s is not a directory", dir))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	luaFiles := []string{}
	for _, file := range files {
		fileName := file.Name()
		if !file.IsDir() && strings.Contains(fileName, ".lua") {
			content, err := ioutil.ReadFile(path.Join(dir, fileName))
			if err != nil {
				return []string{}, err
			}

			luaFiles = append(luaFiles, string(content))
		}
	}

	return luaFiles, nil
}
