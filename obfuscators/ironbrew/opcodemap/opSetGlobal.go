package opcodemap

const strSetGlobal = "Env[Inst[OP_B]] = Stk[Inst[OP_A]];"

func (instruction *Instruction) createSetGlobal() uint32 {
	instruction.B--
	return instruction.createABx(opSETGLOBAL)
}