package opcodemap

var opSetList = map[string]string{
	`local A = Inst[OP_A];
	local T = Stk[A];
	for Idx = A + 1, Inst[OP_B] do 
		Insert(T, Stk[Idx])
	end;` : "OpSetList",

	`local A = Inst[OP_A];
	local T = Stk[A];
	for Idx = A + 1, Top do 
		Insert(T, Stk[Idx])
	end;` : "OpSetListB0",

	`InstrPoint = InstrPoint + 1
	local A = Inst[OP_A];
	local T = Stk[A];
	for Idx = A + 1, Inst[OP_B] do 
		Insert(T, Stk[Idx])
	end;` : "OpSetListC0",
}