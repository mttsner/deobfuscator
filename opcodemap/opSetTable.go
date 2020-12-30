package opcodemap

var opSetTable = map[string]string{
	"Stk[Inst[OP_A]][Stk[Inst[OP_B]]]=Stk[Inst[OP_C]];" : "OpSetTable",
	"Stk[Inst[OP_A]][Inst[OP_B]] = Stk[Inst[OP_C]];" : "OpSetTableB",
	"Stk[Inst[OP_A]][Stk[Inst[OP_B]]] = Inst[OP_C];" : "OpSetTableC",
	"Stk[Inst[OP_A]][Inst[OP_B]] = Inst[OP_C];" : "OpSetTableBC",
}