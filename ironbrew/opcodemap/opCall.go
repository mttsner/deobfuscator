package opcodemap

const (
	strCall = `local A = Inst[OP_A]
	local Results = { Stk[A](Unpack(Stk, A + 1, Inst[OP_B])) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`

	strCallB2 = `local A = Inst[OP_A]
	local Results = { Stk[A](Stk[A + 1]) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`

	strCallB0 = `local A = Inst[OP_A]
	local Results = { Stk[A](Unpack(Stk, A + 1, Top)) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`

	strCallB1 = `local A = Inst[OP_A]
	local Results = { Stk[A]() };
	local Limit = Inst[OP_C];
	local Edx = 0;
	for Idx = A, Limit do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`

	strCallC0 = `local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Unpack(Stk, A + 1, Inst[OP_B])))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`

	strCallC0B2 = `local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Stk[A + 1]))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;`

	strCallC1 = `local A = Inst[OP_A]
	Stk[A](Unpack(Stk, A + 1, Inst[OP_B]))`

	strCallC1B2 = `local A = Inst[OP_A]
	Stk[A](Stk[A + 1])`

	strCallB0C0 = `local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Unpack(Stk, A + 1, Top)))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;`

	strCallB0C1 = `local A = Inst[OP_A]
	Stk[A](Unpack(Stk, A + 1, Top))`

	strCallB1C0 = `local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A]())
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;`

	strCallB1C1 = "Stk[Inst[OP_A]]();"

	strCallC2 = `local A = Inst[OP_A]
	Stk[A] = Stk[A](Unpack(Stk, A + 1, Inst[OP_B]))`

	strCallC2B2 = `local A = Inst[OP_A]
	Stk[A] = Stk[A](Stk[A + 1]) `

	strCallB0C2 = `local A = Inst[OP_A]
	Stk[A] = Stk[A](Unpack(Stk, A + 1, Top))`

	strCallB1C2 = `local A = Inst[OP_A]
	Stk[A] = Stk[A]()`
)

func (instruction *Instruction) createCall() uint32 {
	instruction.B = instruction.B - instruction.A + 1
	instruction.C = instruction.C - instruction.A + 2
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB2() uint32 {
	instruction.C = instruction.C - instruction.A + 2
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB0() uint32 {
	instruction.C = instruction.C - instruction.A + 2
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB1() uint32 {
	instruction.C = instruction.C - instruction.A + 2
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC0() uint32 {
	instruction.B = instruction.B - instruction.A + 1
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC0B2() uint32 {
	instruction.B = instruction.B - instruction.A + 1
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC1() uint32 {
	instruction.B = instruction.B - instruction.A + 1
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC1B2() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB0C0() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB0C1() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB1C0() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB1C1() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC2() uint32 {
	instruction.B = instruction.B - instruction.A + 2
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallC2B2() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB0C2() uint32 {
	return instruction.createABC(opCALL)
}

func (instruction *Instruction) createCallB1C2() uint32 {
	return instruction.createABC(opCALL)
}