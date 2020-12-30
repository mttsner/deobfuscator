package opcodemap

var opMod = map[string]string{
	"Stk[Inst[OP_A]]=Stk[Inst[OP_B]]%Stk[Inst[OP_C]];" : "OpMod",
	"Stk[Inst[OP_A]] = Inst[OP_B] % Stk[Inst[OP_C]];" : "OpModB",
	"Stk[Inst[OP_A]] = Stk[Inst[OP_B]] % Inst[OP_C];" : "OpModC",
	"Stk[Inst[OP_A]]= Inst[OP_B] % Inst[OP_C];" : "OpModBC",
}