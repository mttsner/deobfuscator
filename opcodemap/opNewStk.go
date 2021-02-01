package opcodemap

const strNewStk = "Stk = {};for Idx = 0, PCount do if Idx < Params then Stk[Idx] = Args[Idx + 1]; else break end; end;"

func (instruction *Instruction) createNewStk() uint32 {
	panic("Controlflow detected. opNewStk.go")
}