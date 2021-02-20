package opcodemap

// Instruction holds all the data relevant for creating a instruction. 
type Instruction struct {
	A   int
	B   int
	C   int
	Bx  int
	sBx int
	PC  int
}

func (instruction *Instruction) createABC(op int) uint32 {
	return uint32(op)<<0          |
		uint32(instruction.A)<<6  |
		uint32(instruction.B)<<23 |
		uint32(instruction.C)<<14
}

func (instruction *Instruction) createABx(op int) uint32 {
	return uint32(op)<<0          |
		uint32(instruction.A)<<6  |
		uint32(instruction.Bx)<<14
}

func (instruction *Instruction) createAsBx(op int) uint32 {
	instruction.Bx += instruction.sBx + 131071
	return instruction.createABx(op)
}