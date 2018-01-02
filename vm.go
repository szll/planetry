package main

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/aarzilli/golua/lua"
	"github.com/stevedonovan/luar"
)

var re *regexp.Regexp

func init() {
	re, _ = regexp.Compile("function ([a-zA-Z0-9]*)\\(\\)")
}

type Vm struct {
	luaVM   *lua.State
	scene   *Scene
	scripts map[string]*luar.LuaObject
}

// NewVM creates a new Lua VM for a given scene
func NewVM(scene *Scene) *Vm {
	if scene == nil {
		return nil
	}

	vm := Vm{
		luaVM:   luar.Init(),
		scene:   scene,
		scripts: map[string]*luar.LuaObject{},
	}

	luar.Register(vm.luaVM, "", scene.getVMMethodes())

	return &vm
}

// Destroy releases all the resources of the VM and destroys the VM
func (vm *Vm) Destroy() {
	for scriptName := range vm.scripts {
		vm.scripts[scriptName].Close()
	}

	vm.luaVM.Close()
	vm.luaVM = nil
}

// LoadScripts loads all lua scripts inside a given directory
func (vm *Vm) LoadScripts(dir string) (int, error) {
	luaFiles, err := loadAllLuaFiles(dir)
	if err != nil {
		return 0, err
	}

	i := 0
	for _, fileContent := range luaFiles {
		err = vm.loadScript(fileContent)
		if err != nil {
			return i, err
		}

		i = i + 1
	}

	return i, nil
}

// loadScript loads a script (Lua function as string) into the Lua VM and saves a handle for it
func (vm *Vm) loadScript(script string) error {
	submatches := re.FindAllStringSubmatch(script, 1)

	if len(submatches) != 1 {
		return errors.New(fmt.Sprintf("Lua script is invalid (no function found):\n%s", script))
	}

	name := submatches[0][1]

	if vm.scripts[name] != nil {
		return errors.New(fmt.Sprintf("Function %s was registered before", name))
	}

	// Load into VM
	err := vm.luaVM.DoString(script)
	if err != nil {
		return err
	}

	// Create and save handle
	vm.scripts[name] = luar.NewLuaObjectFromName(vm.luaVM, name)

	return nil
}

// CallScripts calls all saved script function handles
func (vm *Vm) CallScripts() error {
	for scriptName := range vm.scripts {
		_, err := vm.scripts[scriptName].Call(nil)
		if err != nil {
			return err
		}
	}

	return nil
}
