package deobfuscator

import (
	"./opcodemap"
	"github.com/yuin/gopher-lua"
)

func newFunctionProto() *lua.FunctionProto {
	return &lua.FunctionProto{
		SourceName:         "noobmaster(254658795878219778)",
		LineDefined:        0,
		LastLineDefined:    0,
		NumUpvalues:        0,
		NumParameters:      0,
		IsVarArg:           7,
		NumUsedRegisters:   250,
		Code:               make([]uint32, 0, 128),
		Constants:          make([]lua.LValue, 0, 32),
		FunctionPrototypes: make([]*lua.FunctionProto, 0, 16),

		DbgSourcePositions: []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1},//make([]int, 0, 128),
		DbgLocals:          make([]*lua.DbgLocalInfo, 0, 16),
		DbgCalls:           make([]lua.DbgCall, 0, 128),
		DbgUpvalues:        make([]string, 0, 16),
	}
}

func addOp(op string, a, b, c int, function *lua.FunctionProto, f func(int,int,int,int) uint32, PC *int) {
	switch op {
	case "OpLeB":
		b = 255 + a
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpLeC":
		c = 255 + c
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpLeBC":
		b = 255 + a
		c = 255 + c
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpLtB":
		b = 255 + a
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpLtC":
		c = 255 + c
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpLtBC":
		b = 255 + a
		c = 255 + c
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpGeB":
		b = 255 + a
		a = 1
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpGeC":
		c = 255 + c
		a = 1
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpGeBC":
		b = 255 + a
		c = 255 + c
		a = 1
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpEqB":
		b = 255 + a
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpEqC":
		c = 255 + c
		b = a
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpEqBC":
		b = 255 + a
		c = 255 + c
		a = 0
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpSelfC":
		c = 255 + c
	case "OpSetList":
		b -= a
	case "OpSetListC0":
		b -= a
	case "OpSetGlobal":
		b--
	case "OpSetTableB":
		b = 255 + b
	case "OpSetTableC":
		c = 255 + c
	case "OpSetTableBC":
		b = 255 + b
		c = 255 + c
	case "OpLoadK":
		b--
	case "OpGetGlobal":
		b--
	case "OpCall": //TODO: figure out all the other calls lmao
		c = c-a+2
	case "OpCallC1":
		b = b-a+1
	case "OpTestSetC":
		b = c
		c = 1
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, b, c))
		return
	case "OpTForLoop":
		function.Code = append(function.Code, opCreateABC(opcodemap.VOpToOp(op), a, 0, c))
		return
	case "OpJmp":
		b = b-*PC
	case "OpLoadBoolC":
		*PC++
	case "OpForPrep":
		b = b-*PC-1//THIS IS VERY IFY. ME NO LIKeY :(
	case "OpForLoop":
		b = b-*PC
	case "OpVarArg":
		b = b-2
	case "OpTailCall":
		b = b-a+1
	case "OpAddB":
		b = 255 + b
	case "OpAddC":
		c = 255 + c 
	case "OpAddBC":
		b = 255 + b
		c = 255 + c
	}
	function.Code = append(function.Code, f(opcodemap.VOpToOp(op), a, b, c))
}

func (data *vmdata) deserialize(Opcodemap map[int][]string) *lua.FunctionProto {
	function := newFunctionProto()
	PC := 1
	var upvalarray []int

	constCount := data.gBits32()
	for i := 0; i < constCount; i++ {
		switch data.gBits8() {
		case data.Bool:
			function.Constants = append(function.Constants, lua.LBool(data.gBits8() != 0))
		case data.Float:
			function.Constants = append(function.Constants, lua.LNumber(data.gFloat()))
		case data.String:
			function.Constants = append(function.Constants, lua.LString(data.gString()))
		}
	}

	var superop int = 1
	var currentSuperOp int
	var isSuperOp bool

	for _, v := range data.Order {
		switch v {
		case 1:
			function.NumParameters = uint8(data.gBits8())
		case 2:
			protosCount := data.gBits32()
			for i := 0; i < protosCount; i++ {
				function.FunctionPrototypes = append(function.FunctionPrototypes, data.deserialize(Opcodemap))
			}
		case 3:
			instructions := data.gBits32()

			for i := 0; i < instructions; i++ {
				descriptor := data.gBits8()
				if gBit(descriptor,0,0) == 0 {
					Type := gBit(descriptor,1,2)
					//mask := gBit(descriptor,3,5)
					o := data.gBits16()
					op := Opcodemap[o][0]

					if op == "SuperOperator" && !isSuperOp {
						isSuperOp = true
						currentSuperOp = o
					}

					if isSuperOp {
						op = Opcodemap[currentSuperOp][superop]
						if superop == len(Opcodemap[currentSuperOp]) {
							superop = 0
							currentSuperOp = 0
							isSuperOp = false
						}
						superop++
					}
				
					a := data.gBits16()
					var b int
					var c int

					switch Type {
					case 0: //ABC
						b = data.gBits16()
						c = data.gBits16()
						addOp(op, a, b, c, function, opCreateABC, &PC)
					case 1: // ABx
						b = data.gBits32()
						
						if op == "OpClosureNU" {upvalarray = append(upvalarray, 0)}
						
						addOp(op, a, b, c, function, opCreateABx, &PC)
					case 2: // AsBx
						b = data.gBits32() - 65536
						addOp(op, a, b, c, function, opCreateASbx, &PC)
					case 3: // AsBxC
						b = data.gBits32() - 65536
						c = data.gBits16()
						
						if op == "OpClosure" {upvalarray = append(upvalarray, c)}

						addOp(op, a, b, c, function, opCreateABx, &PC)
					}
					PC++
				}

			}
		}
	}

	if len(upvalarray) != 0 {
		for i, proto := range function.FunctionPrototypes {
			proto.NumUpvalues = uint8(upvalarray[i])
		}
	}

	return function
}