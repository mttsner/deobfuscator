package deobfuscator

import (

)

const opSizeCode = 6
const opSizeA = 8
const opSizeB = 9
const opSizeC = 9
const opSizeBx = 18
const opSizesBx = 18

const opMaxArgsA = (1 << opSizeA) - 1
const opMaxArgsB = (1 << opSizeB) - 1
const opMaxArgsC = (1 << opSizeC) - 1
const opMaxArgBx = (1 << opSizeBx) - 1
const opMaxArgSbx = opMaxArgBx >> 1

func opGetOpCode(inst uint32) int {
	return int(inst >> 26)
}

func opSetOpCode(inst *uint32, opcode int) {
	*inst = (*inst & 0x3ffffff) | uint32(opcode<<26)
}

func opGetArgA(inst uint32) int {
	return int(inst>>18) & 0xff
}

func opSetArgA(inst *uint32, arg int) {
	*inst = (*inst & 0xfc03ffff) | uint32((arg&0xff)<<18)
}

func opGetArgB(inst uint32) int {
	return int(inst & 0x1ff)
}

func opSetArgB(inst *uint32, arg int) {
	*inst = (*inst & 0xfffffe00) | uint32(arg&0x1ff)
}

func opGetArgC(inst uint32) int {
	return int(inst>>9) & 0x1ff
}

func opSetArgC(inst *uint32, arg int) {
	*inst = (*inst & 0xfffc01ff) | uint32((arg&0x1ff)<<9)
}

func opGetArgBx(inst uint32) int {
	return int(inst & 0x3ffff)
}

func opSetArgBx(inst *uint32, arg int) {
	*inst = (*inst & 0xfffc0000) | uint32(arg&0x3ffff)
}

func opGetArgSbx(inst uint32) int {
	return opGetArgBx(inst) - opMaxArgSbx
}

func opSetArgSbx(inst *uint32, arg int) {
	opSetArgBx(inst, arg+opMaxArgSbx)
}

func opCreateABC(op, a, b, c int) uint32 {
	var inst uint32 = 0
	opSetOpCode(&inst, op)
	opSetArgA(&inst, a)
	opSetArgB(&inst, b)
	opSetArgC(&inst, c)
	return inst
}

func opCreateABx(op, a, bx, _ int) uint32 {
	var inst uint32 = 0
	opSetOpCode(&inst, op)
	opSetArgA(&inst, a)
	opSetArgBx(&inst, bx)
	return inst
}

func opCreateASbx(op, a, sbx, _ int) uint32 {
	var inst uint32 = 0
	opSetOpCode(&inst, op)
	opSetArgA(&inst, a)
	opSetArgSbx(&inst, sbx)
	return inst
}