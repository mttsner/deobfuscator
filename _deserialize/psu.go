package deserialize

func DeobfuscatePSU(bytestring string ) {
	data := &vmdata{}
	data.Bytecode = Decompress(bytestring)
	//fmt.Println(data.Bytecode)
	data.Key = 185
	data.Bool = 3
	data.Float = 41
	data.String = 37
	data.Obfuscator = obfuscator{Name: "PSU"}
	constCount := data.gBits32()
	fmt.Println(constCount)
	for i := 0; i < constCount; i++ {
		temp := data.gBits8()
		switch temp {
		case data.Bool:
			_ = lua.LBool(data.gBits8() != 0)
		case data.Float:
			_ = lua.LNumber(data.gFloat())
		case data.String:
			_ = lua.LString(data.gString())
		}
	}
	instr := data.gBits32()
	fmt.Println(instr)
	for i := 0; i < instr; i++ {
		descriptor := data.gBits8()
		fmt.Println("Descriptor:", descriptor)
		if descriptor != 0 {
			descriptor--
			t := gBit(descriptor, 0, 2)
			//fmt.Println("Type:", t)
			switch t {
			case 6: 
				// Does nothing 
			case 3:
				_ = data.gBits16()
				_ =	data.gBits8()
				_ =	data.gBits32()
				_ = data.gBits16()
			case 2:
				_ = data.gBits16()
				_ = data.gBits8()
				_ = data.gBits32()
			case 0:
				_ = data.gBits16()
				_ = data.gBits8()
				_ = data.gBits16()
				_ = data.gBits16()
			case 5:
				_ = data.gBits16()
				_ =	data.gBits8()
				_ = data.gBits32()
				C := data.gBits16()
				for i := 0; i < C; i++ {
					_ = data.gBits8()
					_ = data.gBits16()
				}
			case 1:	
				_ = data.gBits16()
				_ =	data.gBits8()
				_ = data.gBits32()
			}
			if gBit(descriptor, 7, 7) == 1 {
				_ = data.gBits32()
			}

			if gBit(descriptor, 6, 6) == 1 {
				_ = data.gBits32()
			}
		}
	}
}