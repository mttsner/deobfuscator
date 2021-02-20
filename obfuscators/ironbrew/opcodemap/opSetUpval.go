package opcodemap

const strSetUpval = "Upvalues[Inst[OP_B]]=Stk[Inst[OP_A]];"

func (instruction *Instruction) createSetUpval() uint32 {
	return instruction.createABC(opSETUPVAL)
}