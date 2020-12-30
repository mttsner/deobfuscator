package opcodemap

var opPow = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]]^Stk[Inst[OP_C]];" : "OpPow",
	"Stk[Inst[OP_A]]= Inst[OP_B] ^ Stk[Inst[OP_C]];" : "OpPowB",
	"Stk[Inst[OP_A]]= Stk[Inst[OP_B]]^ Inst[OP_C];" : "OpPowC",
	"Stk[Inst[OP_A]] = Inst[OP_B] ^ Inst[OP_C];" : "OpPowBC",
}