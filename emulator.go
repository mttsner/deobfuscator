package deobfuscator

import (
	"math"
	"encoding/binary"
)

func (data *vmdata) read(bs []byte) []byte{
	for i, b := range bs {
		bs[i] = b ^ data.Key
		if data.Obfuscator.Name == "PSU" {
			data.Key = byte(math.Mod(float64(bs[i]), 256))
		} 
	}
	data.Pos += len(bs)
	return bs
}

func (data *vmdata) gBits8() int {
	F := data.Bytecode[data.Pos] ^ data.Key
	if data.Obfuscator.Name == "PSU" {
		data.Key = byte(math.Mod(float64(F), 256))
	} 
	data.Pos++
	return int(F)
}

func (data *vmdata) gBits16() int {
	d := data.read(data.Bytecode[data.Pos:data.Pos+2])
	return int(binary.LittleEndian.Uint16(d))
}

func (data *vmdata) gBits32() int {
	d := data.read(data.Bytecode[data.Pos:data.Pos+4])
	return int(binary.LittleEndian.Uint32(d))
}

func (data *vmdata) gFloat() float64 {
	d := data.read(data.Bytecode[data.Pos:data.Pos+8])
	x := binary.LittleEndian.Uint64(d)
	return math.Float64frombits(x)
}


func (data *vmdata) gString() string {
	length := int(data.gBits32())
	return string(data.read(data.Bytecode[data.Pos:data.Pos+length]))
}

func gBit(num, start, end int) int {
	mask := ((1<<(end-start+1))-1) << start
	return (num & mask) >> start
}