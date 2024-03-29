package opcodemap

const strTForLoop = `
	local A = Inst[OP_A];
	local C = Inst[OP_C];
	local CB = A + 2
	local Result = {Stk[A](Stk[A + 1],Stk[CB])};
	for Idx = 1, C do 
		Stk[CB + Idx] = Result[Idx];
	end;
	local R = Result[1]
	if R then 
		Stk[CB] = R 
		InstrPoint = Inst[OP_B];
	else
		InstrPoint = InstrPoint + 1;
	end;
`

func (instruction *Instruction) createTForLoop() uint32 {
	instruction.B = 0
	return instruction.createABC(opTFORLOOP)
}