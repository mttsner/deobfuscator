package opcodemap

const strGetGlobal = "Stk[Inst[OP_A]]=Env[Inst[OP_B]];"

func (instruction *Instruction) createGetGlobal() uint32 {
	instruction.Bx = instruction.B - 1
	return instruction.createABx(opGETGLOBAL)
}