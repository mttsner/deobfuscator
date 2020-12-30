package opcodemap

var opGt = map[string]string{
	"if (Stk[Inst[OP_A]] <= Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGt",
	"if (Inst[OP_A] <= Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGtB",
	"if (Stk[Inst[OP_A]] <= Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGtC",
	"if (Inst[OP_A] <= Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGtBC",
}