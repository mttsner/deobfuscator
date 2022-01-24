package ironbrew

import (
	"github.com/notnoobmaster/deobfuscator/helper"
)

func (data *vmdata) deserialize() (*helper.FunctionProto, error) {
	function := helper.NewFunctionProto()
	for _, v := range data.Order {
		switch v {
		case constants:
			constCount := data.readUint32()
			for i := 0; i < constCount; i++ {
				switch data.readUint8() {
				case data.Bool:
					function.Constants = append(function.Constants, data.readUint8() != 0)
				case data.Float:
					function.Constants = append(function.Constants, data.readFloat())
				case data.String:
					function.Constants = append(function.Constants, data.readString())
				}
			}
		case parameters:
			function.NumParameters = uint8(data.readUint8())
		case prototypes:
			protosCount := data.readUint32()
			for i := 0; i < protosCount; i++ {
				proto, err := data.deserialize()
				if err != nil {
					return nil, err
				}
				function.FunctionPrototypes = append(function.FunctionPrototypes, proto)
			}
		case instructions:
			var pc int
			
			instructions := data.readUint32()
			for i := 0; i < instructions; i++ {
				descriptor := data.readUint8()
				if GetBit(int(descriptor),0,0) == 0 {
					Type := GetBit(int(descriptor),1,2)
					opcode := data.readUint16()

					for _, instruction := range data.Opcodemap[opcode] {
						instruction.PC = pc
						switch Type {
						case 0: // ABC
							instruction.A = data.readUint16()
							instruction.B = data.readUint16()
							instruction.C = data.readUint16()
						case 1: // ABx
							instruction.A = data.readUint16()
							instruction.B = data.readUint32()
						case 2: // AsBx
							instruction.A = data.readUint16()
							instruction.B = data.readUint32() - 65536
						case 3: // AsBxC
							instruction.A = data.readUint16()
							instruction.B = data.readUint32() - 65536
							instruction.C = data.readUint16()
						}
						function.Code = append(function.Code, instruction.Create())
						pc++
					}
				}
			}
		case lineinfo:
			lineCount := data.readUint32()
			for i := 0; i < lineCount; i++ {
				function.DbgSourcePositions = append(function.DbgSourcePositions, data.readUint32())
			}
		}
	}
	return function, nil
}