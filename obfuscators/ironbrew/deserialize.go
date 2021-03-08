package ironbrew

import (
	"github.com/notnoobmaster/deobfuscator/obfuscators/ironbrew/opcodemap"
	"github.com/notnoobmaster/deobfuscator/helper"
	"github.com/yuin/gopher-lua"
)

func (data *vmdata) deserialize() (*lua.FunctionProto, error) {
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
				proto, err := data.deserialize()
				if err != nil {
					return nil, err
				}
				function.FunctionPrototypes = append(function.FunctionPrototypes, proto)
			}
		case instructions:
			var pc int
			var isSuperop bool
			var superop *opcodemap.SuperOperator
			
			instructions := data.gBits32()
			for i := 0; i < instructions; i++ {
				descriptor := data.gBits8()
				if helper.GetBit(descriptor,0,0) == 0 {
					Type := helper.GetBit(descriptor,1,2)
					opcode := data.gBits16()
					instruction := data.Opcodemap[opcode]
					
					if isSuperop {
						if opcode == 0 {
							instruction = superop.Instructions[superop.Pos]
							superop.Pos++
						} else {
							isSuperop = false
						}
					} else if (*instruction).IsSuperop {
						superop = instruction.Superop
						instruction = superop.Instructions[0]
						superop.Pos++
						isSuperop = true
					}

					instruction.PC = pc // := &opcodemap.Instruction{PC: pc}
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
					function.Code = append(function.Code, instruction.Create())
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
	return function, nil
}