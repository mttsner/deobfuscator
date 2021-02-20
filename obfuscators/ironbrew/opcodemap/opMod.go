package opcodemap

const (
	strMod   = "Stk[Inst[OP_A]]=Stk[Inst[OP_B]]%Stk[Inst[OP_C]];"
	strModB  = "Stk[Inst[OP_A]] = Inst[OP_B] % Stk[Inst[OP_C]];"
	strModC  = "Stk[Inst[OP_A]] = Stk[Inst[OP_B]] % Inst[OP_C];"
	strModBC = "Stk[Inst[OP_A]]= Inst[OP_B] % Inst[OP_C];"
)

func (instruction *Instruction) createMod() uint32 {
	return instruction.createABC(opMOD)
}

func (instruction *Instruction) createModB() uint32 {
	instruction.B += 255
	return instruction.createABC(opMOD)
}

func (instruction *Instruction) createModC() uint32 {
	instruction.C += 255
	return instruction.createABC(opMOD)
}

func (instruction *Instruction) createModBC() uint32 {
	instruction.B += 255
	instruction.C += 255
	return instruction.createABC(opMOD)
}