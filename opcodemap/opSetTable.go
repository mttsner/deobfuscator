package opcodemap

const (
	strSetTable   = "Stk[Inst[OP_A]][Stk[Inst[OP_B]]]=Stk[Inst[OP_C]];"
	strSetTableB  = "Stk[Inst[OP_A]][Inst[OP_B]] = Stk[Inst[OP_C]];"
	strSetTableC  = "Stk[Inst[OP_A]][Stk[Inst[OP_B]]] = Inst[OP_C];"
	strSetTableBC = "Stk[Inst[OP_A]][Inst[OP_B]] = Inst[OP_C];"
)

func (instruction *Instruction) createSetTable() uint32 {
	return instruction.createABC(opSETTABLE)
}

func (instruction *Instruction) createSetTableB() uint32 {
	instruction.B += 255
	return instruction.createABC(opSETTABLE)
}

func (instruction *Instruction) createSetTableC() uint32 {
	instruction.C += 255
	return instruction.createABC(opSETTABLE)
}

func (instruction *Instruction) createSetTableBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opSETTABLE)
}