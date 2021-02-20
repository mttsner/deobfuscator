package ironbrew

import (
	"github.com/notnoobmaster/deobfuscator/obfuscators/ironbrew/opcodemap"
	"github.com/notnoobmaster/deobfuscator/helper"
	"github.com/yuin/gopher-lua"
)

func (data *vmdata) deserialize(Opcodemap map[int]func(*opcodemap.Instruction)uint32) *lua.FunctionProto {
	function := helper.NewFunctionProto()
	for _, v := range data.Order {
		switch v {
		case constants:
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
		case parameters:
			function.NumParameters = uint8(data.gBits8())
		case prototypes:
			protosCount := data.gBits32()
			for i := 0; i < protosCount; i++ {
				function.FunctionPrototypes = append(function.FunctionPrototypes, data.deserialize(Opcodemap))
			}
		case instructions:
			var pc int
			instructions := data.gBits32()
			for i := 0; i < instructions; i++ {
				descriptor := data.gBits8()
				if helper.GetBit(descriptor,0,0) == 0 {
					Type := helper.GetBit(descriptor,1,2)
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
					function.Code = append(function.Code, createOpcode(instruction))
					pc++
				}
			}
		case lineinfo:
			lineCount := data.gBits32()
			for i := 0; i < lineCount; i++ {
				function.DbgSourcePositions = append(function.DbgSourcePositions, data.gBits32())
			}
		}
	}
	return function
}