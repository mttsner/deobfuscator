package opcodemap

var opMul = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]]*Stk[Inst[OP_C]];" : "OpMul",
	"Stk[Inst[OP_A]]=Inst[OP_B]*Stk[Inst[OP_C]];" : "OpMulB",
	"Stk[Inst[OP_A]] = Stk[Inst[OP_B]] * Inst[OP_C];" : "OpMulC",
	"Stk[Inst[OP_A]]=Inst[OP_B] * Inst[OP_C]" : "OpMulBC",
}