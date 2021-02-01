package opcodemap

const (
	strSub   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]]-Stk[Inst[OP_C]];"
	strSubB  = "Stk[Inst[OP_A]] = Inst[OP_B] - Stk[Inst[OP_C]];"
	strSubC  = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]] - Inst[OP_C];"
	strSubBC = "Stk[Inst[OP_A]] = Inst[OP_B]- Inst[OP_C];"
)

func (instruction *Instruction) createSub() uint32 {
	return instruction.createABC(opSUB)
}

func (instruction *Instruction) createSubB() uint32 {
	instruction.B += 255
	return instruction.createABC(opSUB)
}

func (instruction *Instruction) createSubC() uint32 {
	instruction.C += 255
	return instruction.createABC(opSUB)
}

func (instruction *Instruction) createSubBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opSUB)
}