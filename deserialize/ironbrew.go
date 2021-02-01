package deserialize

import (
	"github.com/notnoobmaster/opcodemap"
	"github.com/yuin/gopher-lua"
)

func (data *vmdata) Ironbrew(Opcodemap map[int][]string) *lua.FunctionProto {
	function := newFunctionProto()
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
	for _, v := range data.Order {
		switch v {
		case 1:
			function.NumParameters = uint8(data.gBits8())
		case 2:
			protosCount := data.gBits32()
			for i := 0; i < protosCount; i++ {
				function.FunctionPrototypes = append(function.FunctionPrototypes, data.IronBrew(Opcodemap))
			}
		case 3:
			var pc int
			instructions := data.gBits32()
			for i := 0; i < instructions; i++ {
				descriptor := data.gBits8()
				if gBit(descriptor,0,0) == 0 {
					Type := gBit(descriptor,1,2)
					createOpcode := Opcodemap[data.gBits16()]

					instruction := &opcodemap.Instruction{PC: pc}

					switch Type {
					case 0: // ABC
						instruction.A = data.gBits16()
						instruction.B = data.gBits16()
						instruction.C = data.gBits16()
					case 1: // ABx
						instruction.A = data.gBits16()
						instruction.B = data.gBits32()
					case 2: // AsBx
						instruction.A = data.gBits16()
						instruction.B = data.gBits32() - 65536
					case 3: // AsBxC
						instruction.A = data.gBits16()
						instruction.B = data.gBits32() - 65536
						instruction.C = data.gBits16()
					}
					function = append(function, instruction.createOpcode())
					PC++ // Maybe it needs to be outside the if statement
				}
			}
		}
	}
	return function
}