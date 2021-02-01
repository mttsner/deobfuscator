package opcodemap

const strMove = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]];"

func (instruction *Instruction) createMove() uint32 {
	return instruction.createABC(opMOVE)
}