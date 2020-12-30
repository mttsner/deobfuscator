package opcodemap

var opLoadBool = map[string]string{
	"Stk[Inst[OP_A]]=(Inst[OP_B]~=0);" : "OpLoadBool",
	"Stk[Inst[OP_A]]=(Inst[OP_B]~=0);InstrPoint=InstrPoint+1;" : "OpLoadBoolC",
}