package opcodemap

const strLoadNil = "for Idx=Inst[OP_A],Inst[OP_B] do Stk[Idx]=nil;end;"

func (instruction *Instruction) createLoadNil() uint32 {
	return instruction.createABC(opLOADNIL)
}