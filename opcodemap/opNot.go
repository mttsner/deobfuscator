package opcodemap

const strNot = "Stk[Inst[OP_A]]=(not Stk[Inst[OP_B]]);"

func (instruction *Instruction) createNot() uint32 {
	return instruction.createABC(opNOT)
}