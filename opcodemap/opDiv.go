package opcodemap

const (
	strDiv   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]] / Stk[Inst[OP_C]];"
	strDivB  = "Stk[Inst[OP_A]] = Inst[OP_B] / Stk[Inst[OP_C]];"
	strDivC  = "Stk[Inst[OP_A]] = Stk[Inst[OP_B]] / Inst[OP_C];"
	strDivBC = "Stk[Inst[OP_A]] =  Inst[OP_B] / Inst[OP_C];"
)

func (instruction *Instruction) createDiv() uint32 {
	return instruction.createABC(opDIV)
}

func (instruction *Instruction) createDivB() uint32 {
	instruction.B += 255
	return instruction.createABC(opDIV)
}

func (instruction *Instruction) createDivC() uint32 {
	instruction.C += 255
	return instruction.createABC(opDIV)
}

func (instruction *Instruction) createDivBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opDIV)
}