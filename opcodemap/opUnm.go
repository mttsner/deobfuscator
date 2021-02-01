package opcodemap

const strUnm = "Stk[Inst[OP_A]]=-Stk[Inst[OP_B]];"

func (instruction *Instruction) createUnm() uint32 {
	return instruction.createABC(opUNM)
}