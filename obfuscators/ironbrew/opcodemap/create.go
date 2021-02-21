package opcodemap

// SuperOperator holds all the data needed to work with superoperators
type SuperOperator struct {
	Instructions []*Instruction
	Pos int
}

// Instruction holds all the data relevant for creating a instruction. 
type Instruction struct {
	A   int
	B   int
	C   int
	Bx  int
	sBx int
	PC  int
	IsSuperop bool
	Superop SuperOperator
	Func func(*Instruction)uint32
}

func (instruction *Instruction) Create() uint32 {
	return instruction.Func(instruction)
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