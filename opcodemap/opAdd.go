package opcodemap

var opAdd = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]]+Stk[Inst[OP_C]]" : "OpAdd",
	"Stk[Inst[OP_A]] = Inst[OP_B] + Stk[Inst[OP_C]]"  : "OpAddB",
	"Stk[Inst[OP_A]] = Stk[Inst[OP_B]] + Inst[OP_C]"  : "OpAddC",
	"Stk[Inst[OP_A]] = Inst[OP_B] + Inst[OP_C]"       : "OpAddBC",
}