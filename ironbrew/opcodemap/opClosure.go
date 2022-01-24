package opcodemap

const (
	strClosure   = "local NewProto=Proto[Inst[OP_B]];local NewUvals;local Indexes={};NewUvals=Setmetatable({},{__index=function(_,Key)local Val=Indexes[Key];return Val[1][Val[2]];end,__newindex=function(_,Key,Value)local Val=Indexes[Key] Val[1][Val[2]]=Value;end;});for Idx=1,Inst[OP_C] do InstrPoint=InstrPoint+1;local Mvm=Instr[InstrPoint];if Mvm[OP_ENUM]==OP_MOVE then Indexes[Idx-1]={Stk,Mvm[OP_B]};else Indexes[Idx-1]={Upvalues,Mvm[OP_B]};end;Lupvals[#Lupvals+1]=Indexes;end;Stk[Inst[OP_A]]=Wrap(NewProto,NewUvals,Env);"
	strClosureNU = "Stk[Inst[OP_A]]=Wrap(Proto[Inst[OP_B]],nil,Env);"
)
// This needs some special stuff to account for upvalue shit
func (instruction *Instruction) createClosure() uint32 {
	return instruction.createABx(opCLOSURE)
}
func (instruction *Instruction) createClosureNU() uint32 {
	return instruction.createABx(opCLOSURE)
}