package opcodemap

var opNewStk = map[string]string{
	"Stk = {};for Idx = 0, PCount do if Idx < Params then Stk[Idx] = Args[Idx + 1]; else break end; end;" : "OpNewStk",
}