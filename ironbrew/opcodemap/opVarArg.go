package opcodemap

const (
	strVarArg   = "local A=Inst[OP_A];local B=Inst[OP_B];for Idx=A,B do Stk[Idx]=Vararg[Idx-A];end;"
	strVarArgB0 = "local A=Inst[OP_A];Top=A+Varargsz-1;for Idx=A,Top do local VA=Vararg[Idx-A];Stk[Idx]=VA;end;"
)

func (instruction *Instruction) createVarArg() uint32 {
	instruction.B = instruction.B - instruction.A + 1
	return instruction.createABC(opVARARG)
}

func (instruction *Instruction) createVarArgB0() uint32 {
	return instruction.createABC(opVARARG)
}