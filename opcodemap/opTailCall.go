package opcodemap

var opTailCall = map[string]string{
	`local A = Inst[OP_A];
	do return Stk[A](Unpack(Stk, A + 1, Inst[OP_B])) end;` : "OpTailCall",
	`local A = Inst[OP_A];
	do return Stk[A](Unpack(Stk, A + 1, Top)) end;` : "OpTailCallB0",
	`do return Stk[Inst[OP_A]](); end;` : "OpTailCallB1",
}