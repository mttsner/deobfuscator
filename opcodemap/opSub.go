package opcodemap

var opSub = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]]-Stk[Inst[OP_C]];" : "OpSub",
	"Stk[Inst[OP_A]] = Inst[OP_B] - Stk[Inst[OP_C]];" : "OpSubB",
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]] - Inst[OP_C];" : "OpSubC",
	"Stk[Inst[OP_A]] = Inst[OP_B]- Inst[OP_C];" : "OpSubBC",
}