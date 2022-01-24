package helper

type Value interface{}

type DbgLocalInfo struct {
	Name    string
	StartPc int
	EndPc   int
}

type DbgCall struct {
	Name string
	Pc   int
}

type FunctionProto struct {
	SourceName         string
	LineDefined        int
	LastLineDefined    int
	NumUpvalues        uint8
	NumParameters      uint8
	IsVarArg           uint8
	NumUsedRegisters   uint8
	Code               []uint32
	Constants          []Value
	FunctionPrototypes []*FunctionProto

	DbgSourcePositions []int
	DbgLocals          []*DbgLocalInfo
	DbgCalls           []DbgCall
	DbgUpvalues        []string

	stringConstants []string
}

func NewFunctionProto() *FunctionProto {
	return &FunctionProto{
		SourceName:         "deobfuscated",
		LineDefined:        0,
		LastLineDefined:    0,
		NumUpvalues:        0,
		NumParameters:      0,
		IsVarArg:           0,
		NumUsedRegisters:   2,
		Code:               make([]uint32, 0, 128),
		Constants:          make([]Value, 0, 32),
		FunctionPrototypes: make([]*FunctionProto, 0, 16),

		DbgSourcePositions: make([]int, 0, 128),
		DbgLocals:          make([]*DbgLocalInfo, 0, 16),
		DbgCalls:           make([]DbgCall, 0, 128),
		DbgUpvalues:        make([]string, 0, 16),
	}
}