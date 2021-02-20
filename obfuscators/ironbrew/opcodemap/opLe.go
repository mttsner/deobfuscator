package opcodemap

const (
	strLe   = "if(Stk[Inst[OP_A]]<=Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLeB  = "if(Inst[OP_A] <= Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLeC  = "if(Stk[Inst[OP_A]] <= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strLeBC = "if(Inst[OP_A] <= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createLe() uint32 {
	instruction.B = instruction.A
	instruction.A = 0 
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createLeB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 0
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createLeC() uint32 {
	instruction.B = instruction.A
	instruction.C += 255
	instruction.A = 0
	return instruction.createABC(opLE)
}

func (instruction *Instruction) createLeBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 0
	return instruction.createABC(opLE)
}
