package opcodemap

var opLoadNil = map[string]string{
	"for Idx=Inst[OP_A],Inst[OP_B] do Stk[Idx]=nil;end;" : "OpLoadNil",
}