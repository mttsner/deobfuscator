package opcodemap

const (
	strPow   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]]^Stk[Inst[OP_C]];"
	strPowB  = "Stk[Inst[OP_A]]= Inst[OP_B] ^ Stk[Inst[OP_C]];"
	strPowC  = "Stk[Inst[OP_A]]= Stk[Inst[OP_B]]^ Inst[OP_C];"
	strPowBC = "Stk[Inst[OP_A]] = Inst[OP_B] ^ Inst[OP_C];"
)

func (instruction *Instruction) createPow() uint32 {
	return instruction.createABC(opPOW)
}

func (instruction *Instruction) createPowB() uint32 {
	instruction.B += 255
	return instruction.createABC(opPOW)
}

func (instruction *Instruction) createPowC() uint32 {
	instruction.C += 255
	return instruction.createABC(opPOW)
}

func (instruction *Instruction) createPowBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opPOW)
}