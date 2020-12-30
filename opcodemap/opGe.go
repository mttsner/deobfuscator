package opcodemap

var opGe = map[string]string{
	"if (Stk[Inst[OP_A]]<Stk[Inst[OP_C]])then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGe",
	"if (Inst[OP_A] < Stk[Inst[OP_C]]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGeB",
	"if (Stk[Inst[OP_A]] < Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGeC",
	"if (Inst[OP_A] < Inst[OP_C]) then InstrPoint=Inst[OP_B]; else InstrPoint=InstrPoint+1; end;" : "OpGeBC",
}