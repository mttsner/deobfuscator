package opcodemap

const strLen = "Stk[Inst[OP_A]]=#Stk[Inst[OP_B]];"

func (instruction *Instruction) createLen() uint32 {
	return instruction.createABC(opLEN)
}