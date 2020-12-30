package opcodemap

var opNe = map[string]string{
	"if(Stk[Inst[OP_A]]~=Stk[Inst[OP_C]])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpNe",
	"if(Inst[OP_A] ~= Stk[Inst[OP_C]]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpNeB",
	"if(Stk[Inst[OP_A]] ~= Inst[OP_C]) then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpNeC",
	"if(Inst[OP_A] ~= Inst[OP_C])then InstrPoint=InstrPoint+1;else InstrPoint=Inst[OP_B];end;" : "OpNeBC",
}