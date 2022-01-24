package opcodemap

const (
	strNe   = "if(Stk[Inst[OP_A]]~=Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strNeB  = "if(Inst[OP_A] ~= Stk[Inst[OP_C]]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strNeC  = "if(Stk[Inst[OP_A]] ~= Inst[OP_C]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strNeBC = "if(Inst[OP_A] ~= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createNe() uint32 {
	instruction.B = instruction.A
	instruction.A = 1
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createNeB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 1
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createNeC() uint32 {
	instruction.B = instruction.A
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createNeBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opEQ)
}