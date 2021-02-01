package opcodemap

const (
	strGetTable      = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]][Stk[Inst[OP_C]]];"
	strGetTableConst = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]][Inst[OP_C]];"
)

func (instruction *Instruction) createGetTable() uint32 {
	return instruction.createABC(opGETTABLE)
}

func (instruction *Instruction) createGetTableConst() uint32 {
	instruction.C += 255
	return instruction.createABC(opGETTABLE)
}