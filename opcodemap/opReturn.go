package opcodemap

var opReturn = map[string]string{
	`local A = Inst[OP_A];
	do return Unpack(Stk, A, A + Inst[OP_B]) end;` : "OpReturn",
	`do return Stk[Inst[OP_A]] end` : "OpReturnB2",
	`local A = Inst[OP_A]; 
	do return Stk[A], Stk[A + 1] end` : "OpReturnB3",
	`local A = Inst[OP_A]; 
	do return Unpack(Stk, A, Top) end;` : "OpReturnB0",
	"do return end;" : "OpReturnB1",
}