package vm

type cpu struct {
	// The 16 virtual registers
	v [16]uint8
	// The address register
	i uint16
	// The program counter
	pc uint16
	// The stack pointer
	sp uint8
	// The VM memory
	memory [4096]uint8
}

type instruction struct {
	nibble uint8
	opcode uint16
	nnn    uint16
	nn     uint8
	n      uint8
	x      uint8
	y      uint8
}

func NewCPU() *cpu {
	return &cpu{
		v:      [16]uint8{},
		i:      0x200,
		pc:     0x200,
		sp:     0,
		memory: [4096]uint8{},
	}
}

func (cpu *cpu) loadROM(program []byte) {
	// Load the program into memory after 0x200
	copy(cpu.memory[0x200:], program[:])
}

func (cpu *cpu) fetchNext() instruction {
	// Get the current program counter and increase it by two
	var i = cpu.pc
	cpu.pc += 2

	// Fetch the 16-bit instruction
	var op = uint16(cpu.memory[i])<<8 | uint16(cpu.memory[i+1])

	return instruction{
		nibble: uint8(op >> 12),
		opcode: op,
		nnn:    op & 0x0FFF,
		nn:     uint8(op & 0x00FF),
		n:      uint8(op & 0x000F),
		x:      uint8(op & 0x0F00),
		y:      uint8(op & 0x00F0),
	}
}

func (cpu *cpu) execute(instr instruction) {
	if instr.opcode == 0x00E0 {
		vm.cls()
	}
}
