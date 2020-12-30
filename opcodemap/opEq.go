package opcodemap

var opEq = map[string]string{
	"if(Stk[Inst[OP_A]]==Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpEq",
	"if(Inst[OP_A] == Stk[Inst[OP_C]]) then InstrPoint = InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpEqB",
	"if(Stk[Inst[OP_A]] == Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpEqC",
	"if(Inst[OP_A] == Inst[OP_C]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpEqBC",
}