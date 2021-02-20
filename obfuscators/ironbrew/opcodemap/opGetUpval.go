package opcodemap

const strGetUpval = "Stk[Inst[OP_A]]=Upvalues[Inst[OP_B]];"

func (instruction *Instruction) createGetUpval() uint32 {
	return instruction.createABC(opGETUPVAL)
}