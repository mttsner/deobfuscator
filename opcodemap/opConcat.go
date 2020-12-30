package opcodemap

var opConcat = map[string]string{
	"local B=Inst[OP_B];local K=Stk[B] for Idx=B+1,Inst[OP_C] do K=K..Stk[Idx];end;Stk[Inst[OP_A]]=K;" : "OpConcat",
}