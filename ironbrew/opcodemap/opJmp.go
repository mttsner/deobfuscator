package opcodemap

const strJmp = "InstrPoint=Inst[OP_B];"

func (instruction *Instruction) createJmp() uint32 {
	instruction.sBx = instruction.B - instruction.PC - 1 
	return instruction.createAsBx(opJMP)
}