package opcodemap

var opVarArg = map[string]string{
	"local A=Inst[OP_A];local B=Inst[OP_B];for Idx=A,B do Stk[Idx]=Vararg[Idx-A];end;" : "OpVarArg",
	"local A=Inst[OP_A];Top=A+Varargsz-1;for Idx=A,Top do local VA=Vararg[Idx-A];Stk[Idx]=VA;end;" : "OpVarArgB0",
}