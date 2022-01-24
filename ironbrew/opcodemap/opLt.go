package opcodemap

const (
	strLt   = "if(Stk[Inst[OP_A]] < Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLtB  = "if(Inst[OP_A] < Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLtC  = "if(Stk[Inst[OP_A]] < Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLtBC = "if(Inst[OP_A] < Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createLt() uint32 {
	instruction.B = instruction.A
	instruction.A = 0
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createLtB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 0
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createLtC() uint32 {
	instruction.B = instruction.A 
	instruction.C += 255
	instruction.A = 0
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createLtBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 0
	return instruction.createABC(opLT)
}