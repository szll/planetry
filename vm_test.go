package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewVM(t *testing.T) {
	vm := NewVM(nil)

	assert.Nil(t, vm, "vm should be nil")

	vm2 := NewVM(createTestingScene())

	assert.NotNil(t, vm2, "vm2 should not be nil")

	vm2.Destroy()
}

func TestLoadScripts(t *testing.T) {
	vm := NewVM(createTestingScene())
	ls, err := vm.LoadScripts("./testdata")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, ls == 1, "two scripts should have been loaded")

	vm.Destroy()
}

func TestCallScripts(t *testing.T) {
	vm := NewVM(createTestingScene())
	_, err := vm.LoadScripts("./testdata")
	if err != nil {
		t.Error(err)
	}

	err = vm.CallScripts()
	if err != nil {
		t.Error(err)
	}
}
