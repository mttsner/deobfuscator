package opcodemap

var opLe = map[string]string{
	"if(Stk[Inst[OP_A]]<=Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLe",
	"if(Inst[OP_A] <= Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLeB",
	"if(Stk[Inst[OP_A]] <= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLeC",
	"if(Inst[OP_A] <= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpLeBC",
}