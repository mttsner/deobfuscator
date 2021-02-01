package deserialize

func newFunctionProto() *lua.FunctionProto {
	return &lua.FunctionProto{
		SourceName:         "noobmaster(254658795878219778)",
		LineDefined:        0,
		LastLineDefined:    0,
		NumUpvalues:        0,
		NumParameters:      0,
		IsVarArg:           0,
		NumUsedRegisters:   2,
		Code:               make([]uint32, 0, 128),
		Constants:          make([]lua.LValue, 0, 32),
		FunctionPrototypes: make([]*lua.FunctionProto, 0, 16),

		DbgSourcePositions: make([]int, 0, 128),
		DbgLocals:          make([]*lua.DbgLocalInfo, 0, 16),
		DbgCalls:           make([]lua.DbgCall, 0, 128),
		DbgUpvalues:        make([]string, 0, 16),
	}
}