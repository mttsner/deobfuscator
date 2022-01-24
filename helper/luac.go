package helper

import (
	"bytes"
	"encoding/binary"
)

type header struct {
	Signature     uint32
	Version       byte
	Format        byte
	Endianness    byte
	Int           byte
	SizeT         byte
	Instruction   byte
	LuaNumber 	  byte
	IntegralFlag  byte
}

func dumpString(buf *bytes.Buffer, str string) {
	binary.Write(buf, binary.LittleEndian, int32(len(str) + 1))
	buf.WriteString(str)
	buf.WriteByte(0)
}

func dumpInt(buf *bytes.Buffer, i int) {
	binary.Write(buf, binary.LittleEndian, int32(i))
}

func dumpNumber(buf *bytes.Buffer, num float64) {
	binary.Write(buf, binary.LittleEndian, num)
}

func dumpHeader(buf *bytes.Buffer) {
	binary.Write(buf, binary.BigEndian, header{
		Signature:    0x1B4C7561,
		Version:      0x51,
		Format:       0,
		Endianness:   1,
		Int:          4,
		SizeT:        4,
		Instruction:  4,
		LuaNumber: 	  8,
		IntegralFlag: 0,
	})
}

func dumpFunction(buf *bytes.Buffer, f *FunctionProto) {
	dumpString(buf, f.SourceName)
	dumpInt(buf, f.LineDefined) 	 
	dumpInt(buf, f.LastLineDefined)  
	buf.WriteByte(f.NumUpvalues) 	 
	buf.WriteByte(f.NumParameters) 
	buf.WriteByte(f.IsVarArg)
	buf.WriteByte(f.NumUsedRegisters)

	dumpCode(buf, f.Code)
	dumpConstants(buf, f.Constants)
	dumpPrototypes(buf, f.FunctionPrototypes)
	// Filling debug fields with zeros NOTE: This should use the functionproto instead.
	dumpInt(buf, 0)
	dumpInt(buf, 0)
	dumpInt(buf, 0)
}

func dumpCode(buf *bytes.Buffer, code []uint32) {
	dumpInt(buf, len(code))

	for _, inst := range code {
		binary.Write(buf, binary.LittleEndian, inst)
	}
}

func dumpPrototypes(buf *bytes.Buffer, protos []*FunctionProto) {
	dumpInt(buf, len(protos))

	for _, proto := range protos {
		dumpFunction(buf, proto)
	}
}

func dumpConstants(buf *bytes.Buffer, consts []Value) {
	dumpInt(buf, len(consts))

	for _, cons := range consts {
		switch v := cons.(type) {
		case nil:
			buf.WriteByte(0)
		case bool:
			buf.WriteByte(1)
			binary.Write(buf, binary.LittleEndian, v)
		case float64:
			buf.WriteByte(3)
			dumpNumber(buf, v)
		case string:
			buf.WriteByte(4)
			dumpString(buf, v)
		}
	}
}

// DumpLua turns FunctionProto into a luac formated byte slice
func DumpLua(proto *FunctionProto) []byte {
	buf := new(bytes.Buffer)
	dumpHeader(buf)
	dumpFunction(buf, proto)
	return buf.Bytes()
}