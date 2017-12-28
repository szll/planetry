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

func loadColor(obj *jason.Object) (*Color, error) {
	redColorInt, err := obj.GetInt64("red")
	if err != nil {
		return nil, err
	}

	greenColorInt, err := obj.GetInt64("green")
	if err != nil {
		return nil, err
	}

	blueColorInt, err := obj.GetInt64("blue")
	if err != nil {
		return nil, err
	}

	alphaColorInt, err := obj.GetInt64("alpha")
	if err != nil {
		return nil, err
	}

	return &Color{
		Red:   uint8(redColorInt),
		Green: uint8(greenColorInt),
		Blue:  uint8(blueColorInt),
		Alpha: uint8(alphaColorInt),
	}, nil
}

func loadPoint3D(obj *jason.Object) (*Point3D, error) {
	x, err := obj.GetFloat64("x")
	if err != nil {
		return nil, err
	}

	y, err := obj.GetFloat64("y")
	if err != nil {
		return nil, err
	}

	z, err := obj.GetFloat64("z")
	if err != nil {
		return nil, err
	}

	return &Point3D{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func loadVector3D(obj *jason.Object) (*Vector3D, error) {
	x, err := obj.GetFloat64("x")
	if err != nil {
		return nil, err
	}

	y, err := obj.GetFloat64("y")
	if err != nil {
		return nil, err
	}

	z, err := obj.GetFloat64("z")
	if err != nil {
		return nil, err
	}

	return &Vector3D{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func loadBody(obj *jason.Object) (*Body, error) {
	name, err := obj.GetString("name")
	if err != nil {
		return nil, err
	}

	mass, err := obj.GetFloat64("mass")
	if err != nil {
		return nil, err
	}

	radius, err := obj.GetFloat64("radius")
	if err != nil {
		return nil, err
	}

	pObj, err := obj.GetObject("position")
	if err != nil {
		return nil, err
	}

	position, err := loadPoint3D(pObj)
	if err != nil {
		return nil, err
	}

	position.X = position.X * AU
	position.Y = position.Y * AU
	position.Z = position.Z * AU

	vObj, err := obj.GetObject("velocity")
	if err != nil {
		return nil, err
	}

	velocity, err := loadVector3D(vObj)
	if err != nil {
		return nil, err
	}

	velocity.X = velocity.X * 1000
	velocity.Y = velocity.Y * 1000
	velocity.Z = velocity.Z * 1000

	return &Body{
		Name:     name,
		Mass:     mass,
		Radius:   radius,
		Position: position,
		Velocity: velocity,
	}, nil
}

func loadDrawableBody(obj *jason.Object) (*DrawableBody, error) {
	body, err := loadBody(obj)
	if err != nil {
		return nil, err
	}

	cObj, err := obj.GetObject("color")
	if err != nil {
		return nil, err
	}

	color, err := loadColor(cObj)
	if err != nil {
		return nil, err
	}

	return &DrawableBody{
		PhysicalBody: body,
		Path:         PointQueue{},
		Color:        color,
	}, nil
}

func loadScene(path string) (*Scene, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	v, err := jason.NewObjectFromBytes(file)
	if err != nil {
		return nil, err
	}

	bgcObj, err := v.GetObject("backgroundColor")
	if err != nil {
		return nil, err
	}

	bgColor, err := loadColor(bgcObj)
	if err != nil {
		return nil, err
	}

	bodies, err := v.GetObjectArray("bodies")
	if err != nil {
		return nil, err
	}

	var dbs []*DrawableBody
	for _, bodyObj := range bodies {
		db, err := loadDrawableBody(bodyObj)
		if err != nil {
			return nil, err
		}

		dbs = append(dbs, db)
	}

	return &Scene{
		Bodies:          dbs,
		ForcesOfBodies:  map[*DrawableBody]Vector3D{},
		Camera:          nil,
		BackgroundColor: bgColor,
		zoom:            10,
		destroyed:       false,
		simulations:     0,
		paused:          false,
	}, nil
}

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
