package opcodemap

const (
	strMul   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]]*Stk[Inst[OP_C]];"
	strMulB  = "Stk[Inst[OP_A]]=Inst[OP_B]*Stk[Inst[OP_C]];"
	strMulC  = "Stk[Inst[OP_A]] = Stk[Inst[OP_B]] * Inst[OP_C];"
	strMulBC = "Stk[Inst[OP_A]]=Inst[OP_B] * Inst[OP_C]"
)

func (instruction *Instruction) createMul() uint32 {
	return instruction.createABC(opMUL)
}

func (instruction *Instruction) createMulB() uint32 {
	instruction.B += 255
	return instruction.createABC(opMUL)
}

func (instruction *Instruction) createMulC() uint32 {
	instruction.C += 255
	return instruction.createABC(opMUL)
}

func (instruction *Instruction) createMulBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opMUL)
}