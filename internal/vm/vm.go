package vm

import (
	"fmt"
)

type VM struct {
}

func (vm *VM) fetch() uint {
	return 0
}

func (vm *VM) Step() error {
	// Fetch the next instruction
	var instruction = vm.fetch()

	switch instruction {
	default:
		return fmt.Errorf("Unknown opcode: %04X", instruction)
	}

	return nil
}
