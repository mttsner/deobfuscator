package opcodemap

const strConcat = "local B=Inst[OP_B];local K=Stk[B] for Idx=B+1,Inst[OP_C] do K=K..Stk[Idx];end;Stk[Inst[OP_A]]=K;"

func (instruction *Instruction) createConcat() uint32 {
	return instruction.createABC(opCONCAT)
}