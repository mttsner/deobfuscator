package opcodemap

const (
	strEq   = "if(Stk[Inst[OP_A]]==Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strEqB  = "if(Inst[OP_A] == Stk[Inst[OP_C]]) then InstrPoint = InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strEqC  = "if(Stk[Inst[OP_A]] == Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
	strEqBC = "if(Inst[OP_A] == Inst[OP_C]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createEq() uint32 {
	instruction.B = instruction.A
	instruction.A = 0 
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createEqB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 0
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createEqC() uint32 {
	instruction.B = instruction.A
	instruction.C += 255
	instruction.A = 0
	return instruction.createABC(opEQ)
}

func (instruction *Instruction) createEqBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 0 
	return instruction.createABC(opEQ)
}