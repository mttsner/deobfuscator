package opcodemap

const strNewTableB0 = "Stk[Inst[OP_A]]={};"

func (instruction *Instruction) createNewTableB0() uint32 {
	return instruction.createABC(opNEWTABLE)
}