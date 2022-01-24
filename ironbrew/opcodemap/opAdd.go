package opcodemap

const (
	strAdd   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]]+Stk[Inst[OP_C]]"
	strAddB  = "Stk[Inst[OP_A]] = Inst[OP_B] + Stk[Inst[OP_C]]"
	strAddC  = "Stk[Inst[OP_A]] = Stk[Inst[OP_B]] + Inst[OP_C]"
	strAddBC = "Stk[Inst[OP_A]] = Inst[OP_B] + Inst[OP_C]"
)

func (instruction *Instruction) createAdd() uint32 {
	return instruction.createABC(opADD)
}

func (instruction *Instruction) createAddB() uint32 {
	instruction.B += 255
	return instruction.createABC(opADD)
}

func (instruction *Instruction) createAddC() uint32 {
	instruction.C += 255
	return instruction.createABC(opADD)
}

func (instruction *Instruction) createAddBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opADD)
}