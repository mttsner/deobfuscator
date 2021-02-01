package opcodemap

const (
	strLoadBool  = "Stk[Inst[OP_A]]=(Inst[OP_B]~=0);"
	strLoadBoolC = "Stk[Inst[OP_A]]=(Inst[OP_B]~=0);InstrPoint=InstrPoint+1;"
)

func (instruction *Instruction) createLoadBool() uint32 {
	return instruction.createABC(opLOADBOOL)
}

func (instruction *Instruction) createLoadBoolC() uint32 {
	instruction.C = 1
	return instruction.createABC(opLOADBOOL)
}