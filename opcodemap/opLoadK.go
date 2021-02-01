package opcodemap

const strLoadK = "Stk[Inst[OP_A]] = Inst[OP_B];"

func (instruction *Instruction) createLoadK() uint32 {
	instruction.B--
	return instruction.createABx(opLOADK)
}