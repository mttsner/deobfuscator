package opcodemap

var opGetTable = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]][Stk[Inst[OP_C]]];" : "OpGetTable",
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]][Inst[OP_C]];" : "OpGetTableConst",
}