package opcodemap

const (
	strTest  = "if Stk[Inst[OP_A]] then InstrPoint=InstrPoint + 1; else InstrPoint = Inst[OP_B]; end;"
	strTestC = "if not Stk[Inst[OP_A]] then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createTest() uint32 {
	instruction.B = 0 // B is ignored in TEST
	return instruction.createABC(opTEST)
}

func (instruction *Instruction) createTestC() uint32 {
	instruction.B = 0 // B is ignored in TEST
	return instruction.createABC(opTEST)
}