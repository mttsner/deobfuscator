package opcodemap

const strGetGlobal = "Stk[Inst[OP_A]]=Env[Inst[OP_B]];"

func (instruction *Instruction) createGetGlobal() uint32 {
	instruction.B--
	return instruction.createABx(opGETGLOBAL)
}