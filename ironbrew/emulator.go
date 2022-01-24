package ironbrew

import (
	"math"
	"encoding/binary"
)

func (data *vmdata) read(count int) []byte {
	bytes := make([]byte, count)

	for i := 0; i < count; i++ {
		b, _ := data.Bytecode.ReadByte() // Who needs errors anyways?
		bytes[i] = b ^ data.Key
	}

	return bytes
}

func (data *vmdata) readUint8() uint8 {
	return data.read(1)[0]
}

func (data *vmdata) readUint16() int {
	return int(binary.LittleEndian.Uint16(data.read(2)))
}

func (data *vmdata) readUint32() int {
	return int(binary.LittleEndian.Uint32(data.read(4)))
}

func (data *vmdata) readFloat() float64 {
	d := data.read(8)
	x := binary.LittleEndian.Uint64(d)
	return math.Float64frombits(x)
}


func (data *vmdata) readString() string {
	length := int(data.readUint32())
	return string(data.read(length))
}