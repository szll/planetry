package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewVM(t *testing.T) {
	vm := NewVM(nil)

	assert.Nil(t, vm, "vm should be nil")

	vm2 := NewVM(createTestingScene([]string{}))

	assert.NotNil(t, vm2, "vm2 should not be nil")

	vm2.Destroy()
}

func TestLoadScripts(t *testing.T) {
	vm := NewVM(createTestingScene([]string{"script-test.lua", "script-sun-disappears-after-365-steps.lua"}))
	ls, err := vm.LoadScripts("./testdata")
	if err != nil {
		t.Error(err)
	}

	assert.True(t, ls == 2, "two scripts should have been loaded")

	vm.Destroy()
}

func TestCallScripts(t *testing.T) {
	vm := NewVM(createTestingScene([]string{}))
	_, err := vm.LoadScripts("./testdata")
	if err != nil {
		t.Error(err)
	}

	err = vm.CallScripts()
	if err != nil {
		t.Error(err)
	}
}
