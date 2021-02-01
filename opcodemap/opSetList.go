package opcodemap

const (
	strSetList   = "local A = Inst[OP_A]; local T = Stk[A]; for Idx = A + 1, Inst[OP_B] do Insert(T, Stk[Idx]) end;"
	strSetListB0 = "local A = Inst[OP_A]; local T = Stk[A]; for Idx = A + 1, Top do Insert(T, Stk[Idx]) end;"
	strSetListC0 = "InstrPoint = InstrPoint + 1 local A = Inst[OP_A]; local T = Stk[A]; for Idx = A + 1, Inst[OP_B] do Insert(T, Stk[Idx]) end;"
)

func (instruction *Instruction) createSetList() uint32 {
	instruction.B -= instruction.A
	return instruction.createABC(opSETLIST)
}

func (instruction *Instruction) createSetListB0() uint32 {
	return instruction.createABC(opSETLIST)
}

func (instruction *Instruction) createSetListC0() uint32 {
	panic("Figure out a way to get the instrution data of some x thing/ of the next instruction. opSetList.go")
	return instruction.createABC(opSETLIST)
}