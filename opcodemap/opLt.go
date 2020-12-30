package opcodemap

var opLt = map[string]string{
	"if(Stk[Inst[OP_A]] < Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLt",
	"if(Inst[OP_A] < Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLtB",
	"if(Stk[Inst[OP_A]] < Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLtC",
	"if(Inst[OP_A] < Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLtBC",
}