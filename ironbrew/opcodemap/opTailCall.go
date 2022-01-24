package opcodemap

const (
	strTailCall   = "local A = Inst[OP_A]; do return Stk[A](Unpack(Stk, A + 1, Inst[OP_B])) end;"
	strTailCallB0 = "local A = Inst[OP_A]; do return Stk[A](Unpack(Stk, A + 1, Top)) end;"
	strTailCallB1 = "do return Stk[Inst[OP_A]](); end;"
)

func (instruction *Instruction) createTailCall() uint32 {
	instruction.B = instruction.B - instruction.A + 1 
	return instruction.createABC(opTAILCALL)
}

func (instruction *Instruction) createTailCallB0() uint32 {
	return instruction.createABC(opTAILCALL)
}

func (instruction *Instruction) createTailCallB1() uint32 {
	return instruction.createABC(opTAILCALL)
}