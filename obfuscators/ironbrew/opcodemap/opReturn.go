package opcodemap

const (
	strReturn   = "local A = Inst[OP_A]; do return Unpack(Stk, A, A + Inst[OP_B]) end;"
	strReturnB2 = "do return Stk[Inst[OP_A]] end"
	strReturnB3 = "local A = Inst[OP_A]; do return Stk[A], Stk[A + 1] end"
	strReturnB0 = "local A = Inst[OP_A]; do return Unpack(Stk, A, Top) end;"
	strReturnB1 = "do return end;"
)

func (instruction *Instruction) createReturn() uint32 {
	instruction.B += 2
	return instruction.createABC(opRETURN)
}

func (instruction *Instruction) createReturnB2() uint32 {
	return instruction.createABC(opRETURN)
}

func (instruction *Instruction) createReturnB3() uint32 {
	return instruction.createABC(opRETURN)
}

func (instruction *Instruction) createReturnB0() uint32 {
	return instruction.createABC(opRETURN)
}

func (instruction *Instruction) createReturnB1() uint32 {
	return instruction.createABC(opRETURN)
}