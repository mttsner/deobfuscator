package opcodemap

var opCall = map[string]string{
	`local A = Inst[OP_A]
	local Results = { Stk[A](Unpack(Stk, A + 1, Inst[OP_B])) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end`: "OpCall",

	`local A = Inst[OP_A]
	local Results = { Stk[A](Stk[A + 1]) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end` : "OpCallB2",

	`local A = Inst[OP_A]
	local Results = { Stk[A](Unpack(Stk, A + 1, Top)) };
	local Edx = 0;
	for Idx = A, Inst[OP_C] do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end` : "OpCallB0",

	`local A = Inst[OP_A]
	local Results = { Stk[A]() };
	local Limit = Inst[OP_C];
	local Edx = 0;
	for Idx = A, Limit do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end` : "OpCallB1",

	`local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Unpack(Stk, A + 1, Inst[OP_B])))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end` : "OpCallC0",

	`local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Stk[A + 1]))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;` : "OpCallC0B2",

	`local A = Inst[OP_A]
	Stk[A](Unpack(Stk, A + 1, Inst[OP_B]))` : "OpCallC1",

	`local A = Inst[OP_A]
	Stk[A](Stk[A + 1])` : "OpCallC1B2",

	`local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A](Unpack(Stk, A + 1, Top)))
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;` : "OpCallB0C0",

	`local A = Inst[OP_A]
	Stk[A](Unpack(Stk, A + 1, Top))` : "OpCallB0C1",

	`local A = Inst[OP_A]
	local Results, Limit = _R(Stk[A]())
	Top = Limit + A - 1
	local Edx = 0;
	for Idx = A, Top do 
		Edx = Edx + 1;
		Stk[Idx] = Results[Edx];
	end;` : "OpCallB1C0",

	`Stk[Inst[OP_A]]();` : "OpCallB1C1",

	`local A = Inst[OP_A]
	Stk[A] = Stk[A](Unpack(Stk, A + 1, Inst[OP_B])) ` : "OpCallC2",

	`local A = Inst[OP_A]
	Stk[A] = Stk[A](Stk[A + 1]) ` : "OpCallC2B2",

	`local A = Inst[OP_A]
	Stk[A] = Stk[A](Unpack(Stk, A + 1, Top))` : "OpCallB0C2",
	
	`local A = Inst[OP_A]
	Stk[A] = Stk[A]()` : "OpCallB1C2",
}