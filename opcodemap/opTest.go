package opcodemap

var opTest = map[string]string{
	"if Stk[Inst[OP_A]] then InstrPoint=InstrPoint + 1; else InstrPoint = Inst[OP_B]; end;" : "OpTest",
	"if not Stk[Inst[OP_A]] then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpTestC",
}