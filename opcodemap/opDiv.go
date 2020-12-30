package opcodemap

var opDiv = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]] / Stk[Inst[OP_C]];" : "OpDiv",
	"Stk[Inst[OP_A]] = Inst[OP_B] / Stk[Inst[OP_C]];" : "OpDivB",
	"Stk[Inst[OP_A]] = Stk[Inst[OP_B]] / Inst[OP_C];" : "OpDivC",
	"Stk[Inst[OP_A]] =  Inst[OP_B] / Inst[OP_C];" : "OpDivBC",
}