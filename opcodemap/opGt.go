package opcodemap

const (
	strGt   = "if (Stk[Inst[OP_A]] <= Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGtB  = "if (Inst[OP_A] <= Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGtC  = "if (Stk[Inst[OP_A]] <= Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGtBC = "if (Inst[OP_A] <= Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
)

func (instruction *Instruction) createGt() uint32 {
	instruction.B = instruction.A
	instruction.A = 1
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createGtB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 1
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createGtC() uint32 {
	instruction.B = instruction.A
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createGtBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opLE)
}