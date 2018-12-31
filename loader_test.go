package main

import (
	"testing"

	"github.com/antonholmquist/jason"
	"github.com/stretchr/testify/assert"
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

func TestLoadColorError(t *testing.T) {
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

func TestLoadPoint3D(t *testing.T) {
	point := `
		{
			"x": 1.0,
			"y": 2.0,
			"z": 3.0
		}
	`

	pObj, err := jason.NewObjectFromBytes([]byte(point))
	if err != nil {
		t.Error(err)
	}

	p, err := loadPoint3D(pObj)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, p.X, 1.0, "X should be 1.0")
	assert.Equal(t, p.Y, 2.0, "Y should be 2.0")
	assert.Equal(t, p.Z, 3.0, "Z should be 3.0")
}

func TestLoadPointError(t *testing.T) {
	point := `
		{
			"x": "hans",
			"y": 2,
			"z": 3
		}
	`

	pObj, err := jason.NewObjectFromBytes([]byte(point))
	if err != nil {
		t.Error(err)
	}

	_, err = loadPoint3D(pObj)
	assert.Equal(t, err.Error(), "not a number", "should return error 'not a number'")
}

func TestLoadVector3D(t *testing.T) {
	point := `
		{
			"x": 1.0,
			"y": 2.0,
			"z": 3.0
		}
	`

	pObj, err := jason.NewObjectFromBytes([]byte(point))
	if err != nil {
		t.Error(err)
	}

	v, err := loadVector3D(pObj)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, v.X, 1.0, "X should be 1.0")
	assert.Equal(t, v.Y, 2.0, "Y should be 2.0")
	assert.Equal(t, v.Z, 3.0, "Z should be 3.0")
}

func TestLoadVector3DError(t *testing.T) {
	point := `
		{
			"x": "hans",
			"y": 2,
			"z": 3
		}
	`

	pObj, err := jason.NewObjectFromBytes([]byte(point))
	if err != nil {
		t.Error(err)
	}

	_, err = loadVector3D(pObj)
	assert.Equal(t, err.Error(), "not a number", "should return error 'not a number'")
}

func TestLoadBody(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`
	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	b, err := loadBody(bObj)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, b.ID, "id", "Name should be 'id'")
	assert.Equal(t, b.Name, "name", "Name should be 'name'")
	assert.Equal(t, b.Mass, 1.0, "Mass should be 1")
	assert.Equal(t, b.Radius, 1.0, "Radius should be 1")

	assert.Equal(t, b.Position.X, 1.0*AU, "X should be 1.0 * AU")
	assert.Equal(t, b.Position.Y, 2.0*AU, "Y should be 2.0 * AU")
	assert.Equal(t, b.Position.Z, 3.0*AU, "Z should be 3.0 * AU")

	assert.Equal(t, b.Velocity.X, 4000.0, "X should be 4000")
	assert.Equal(t, b.Velocity.Y, 5000.0, "Y should be 5000")
	assert.Equal(t, b.Velocity.Z, 6000.0, "Z should be 6000")
}

func TestLoadBodyErrorId(t *testing.T) {
	body := `
		{
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorName(t *testing.T) {
	body := `
		{
			"id": "sun",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorMass(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorRadius(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorPosition(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorPoint3D(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorVelocity(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadBodyErrorVector3D(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"y": 5.0,
				"z": 6.0
			}
		}
	`

	bObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	_, err = loadBody(bObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadDrawableBody(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			},
			"color": {
				"red": 1,
				"green": 2,
				"blue": 3,
				"alpha": 4
			}
		}
	`
	dObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	// Check that error is nil, the rest was tested before ...
	d, err := loadDrawableBody(dObj)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, d)
}

func TestLoadDrawableBodyErrorBody(t *testing.T) {
	body := `
		{
			"id": "id",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			},
			"color": {
				"red": 1,
				"green": 2,
				"blue": 3,
				"alpha": 4
			}
		}
	`
	dObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	// Check that error is nil, the rest was tested before ...
	_, err = loadDrawableBody(dObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadDrawableBodyErrorColor(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			}
		}
	`
	dObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	// Check that error is nil, the rest was tested before ...
	_, err = loadDrawableBody(dObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadDrawableBodyErrorLoadColor(t *testing.T) {
	body := `
		{
			"id": "id",
			"name": "name",
			"mass": 1.0,
			"radius": 1.0,
			"position": {
				"x": 1.0,
				"y": 2.0,
				"z": 3.0
			},
			"velocity": {
				"x": 4.0,
				"y": 5.0,
				"z": 6.0
			},
			"color": {
				"green": 2,
				"blue": 3,
				"alpha": 4
			}
		}
	`
	dObj, err := jason.NewObjectFromBytes([]byte(body))
	if err != nil {
		t.Error(err)
	}

	// Check that error is nil, the rest was tested before ...
	_, err = loadDrawableBody(dObj)
	assert.Equal(t, err.Error(), "key not found", "should return error 'key not found'")
}

func TestLoadScene(t *testing.T) {
	s, err := loadScene("./testdata/system-second-star.json")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, s)
}

func TestLoadSceneWithoutTarget(t *testing.T) {
	s, err := loadScene("./testdata/system-second-star-without-target.json")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, s)
}

func TestLoadSceneErrorReadFile(t *testing.T) {
	_, err := loadScene("./testdata/system2.json")
	assert.Equal(t, err.Error(), "open ./testdata/system2.json: no such file or directory")
}

func TestLoadSceneErrorNewObjectFromBytes(t *testing.T) {
	_, err := loadScene("./testdata/script-test.lua")
	assert.Equal(t, err.Error(), "invalid character 'h' looking for beginning of value")
}

func TestLoadLuaFiles(t *testing.T) {
	fileContents, err := loadLuaFiles("./testdata", []string{"script-test.lua"})
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, len(fileContents), 1, "there should be one file loaded")
}
