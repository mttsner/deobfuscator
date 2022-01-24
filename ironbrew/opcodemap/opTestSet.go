package opcodemap

const (
	strTestSet  = "local B=Stk[Inst[OP_C]];if B then InstrPoint=InstrPoint+1;else Stk[Inst[OP_A]]=B;InstrPoint=Inst[OP_B];end;"
	strTestSetC = "local B=Stk[Inst[OP_C]];if not B then InstrPoint=InstrPoint+1;else Stk[Inst[OP_A]]=B;InstrPoint=Inst[OP_B];end;"
)

func (instruction *Instruction) createTestSet() uint32 {
	instruction.B = instruction.C
	instruction.C = 0
	return instruction.createABC(opTESTSET)
}

func (instruction *Instruction) createTestSetC() uint32 {
	instruction.B = instruction.C
	instruction.C = 1
	return instruction.createABC(opTESTSET)
}