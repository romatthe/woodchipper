package vm

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

type VM struct {
	cpu cpu
}

func (vm *VM) fetch() uint {
	return 0
}

// Fetch the next instruction and execute it
func (vm *VM) Step() error {
	// Fetch the next instruction
	var instruction = vm.fetch()

	switch instruction {
	default:
		return fmt.Errorf("Unknown opcode: %04X", instruction)
	}

	return nil
}

// Reset the CHIP-8 VM
func (vm *VM) Reset() {

}

// Load a program ROM from a file
func (vm *VM) LoadFile(file string) error {
	if program, err := ioutil.ReadFile(file); err != nil {
		return err
	} else {
		for _, c := range string(program) {
			if unicode.IsSpace(c) || unicode.IsGraphic(c) {
				continue
			}

			// Load the program rom into the CPU memory
			vm.cpu.loadROM(program)
		}

		return nil
	}
}
