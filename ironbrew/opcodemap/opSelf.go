package opcodemap

const (
	strSelf  = "local A=Inst[OP_A];local B=Stk[Inst[OP_B]];Stk[A+1]=B;Stk[A]=B[Stk[Inst[OP_C]]];"
	strSelfC = "local A=Inst[OP_A];local B=Stk[Inst[OP_B]];Stk[A+1]=B;Stk[A]=B[Inst[OP_C]];"
)

func (instruction *Instruction) createSelf() uint32 {
	return instruction.createABC(opSELF)
}

func (instruction *Instruction) createSelfC() uint32 {
	instruction.C += 255
	return instruction.createABC(opSELF)
}