package opcodemap

const (
	strGe   = "if (Stk[Inst[OP_A]]<Stk[Inst[OP_C]])then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGeB  = "if (Inst[OP_A] < Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGeC  = "if (Stk[Inst[OP_A]] < Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
	strGeBC = "if (Inst[OP_A] < Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;"
)

func (instruction *Instruction) createGe() uint32 {
	instruction.B = instruction.A
	instruction.A = 1
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createGeB() uint32 {
	instruction.B = instruction.A + 255
	instruction.A = 1
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createGeC() uint32 {
	instruction.B = instruction.A
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opLT)
}

func (instruction *Instruction) createGeBC() uint32 {
	instruction.B = instruction.A + 255
	instruction.C += 255
	instruction.A = 1
	return instruction.createABC(opLT)
}