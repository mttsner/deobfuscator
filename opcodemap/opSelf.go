package opcodemap

var opSelf = map[string]string{
	"local A=Inst[OP_A];local B=Stk[Inst[OP_B]];Stk[A+1]=B;Stk[A]=B[Stk[Inst[OP_C]]];" : "OpSelf",
	"local A=Inst[OP_A];local B=Stk[Inst[OP_B]];Stk[A+1]=B;Stk[A]=B[Inst[OP_C]];" : "OpSelfC",
}