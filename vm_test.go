package main

import (
	"testing"
)

func TestVM(t *testing.T) {

	scene := Scene{nil, nil, nil, Color{0, 0, 0, 0}, 0, false, 0}

	vm := NewVM(&scene)
	_, err := vm.LoadScripts("./testdata")
	if err != nil {
		t.Error(err)
	}

	err = vm.CallScripts()
	if err != nil {
		t.Error(err)
	}

	vm.Destroy()
}
