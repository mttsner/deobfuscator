package opcodemap


// Map returns all the opcodes as array
func Map() []map[string]string {
	return []map[string]string{
		opDiv, opGetTable, opLoadBool, opMul, opPushStk, opSetTable, opTestSet, 
		opAdd, opEq, opGetUpval, opLoadK, opNe, opReturn, opSetTop, opTForLoop, 
		opCall, opForLoop, opGt, opLoadNil, opNewStk, opSelf, opSetUpval, opUnm, 
		opClose, opForPrep, opJmp, opLt, opNewTable, opSetFEnv, opSub, opVarArg, 
		opClosure, opGe, opLe, opMod, opNot, opSetGlobal, opTailCall, 
		opConcat, opGetGlobal, opLen, opMove, opPow, opSetList, opTest, delimiter,
	}
}

var voptoop = map[string]int{
	// MOVE
	"OpMove" : 0,
	// LOADK
	"OpLoadK" : 1,
	// LOADBOOL
	"OpLoadBool" : 2,
	"OpLoadBoolC" : 2,
	// LOADNIL
	"OpLoadNil" : 3,
	// GETUPVAL
	"OpGetUpval" : 4,
	// GETGLOBAL
	"OpGetGlobal" : 5,
	// GETTABLE
	"OpGetTable" : 6,
	"OpGetTableConst" : 6,
	// SETGLOBAL
	"OpSetGlobal" : 7,
	// SETUPVAL
	"OpSetUpval" : 8,
	// SETTABLE
	"OpSetTable" : 9,
	"OpSetTableB" : 9,
	"OpSetTableC" : 9,
	"OpSetTableBC" : 9,
	// NEWTABLE
	"OpNewTableB0" : 10,
	// SELF
	"OpSelf" : 11,
	"OpSelfC" : 11,
	// ADD
	"OpAddB" : 12,
	"OpAddC" : 12,
	"OpAddBC" : 12,
	"OpAdd" : 12,
	// SUB
	"OpSub" : 13,
	"OpSubB" : 13,
	"OpSubC" : 13,
	"OpSubBC" : 13,
	// MUL
	"OpMulC" : 14,
	"OpMulBC" : 14,
	"OpMul" : 14,
	"OpMulB" : 14,
	// DIV
	"OpDivC" : 15,
	"OpDivBC" : 15,
	"OpDiv" : 15,
	"OpDivB" : 15,
	// MOD
	"OpModB" : 16,
	"OpModC" : 16,
	"OpModBC" : 16,
	"OpMod" : 16,
	// POW
	"OpPow" : 17,
	"OpPowB" : 17,
	"OpPowC" : 17,
	"OpPowBC" : 17,
	// UNM
	"OpUnm" : 18,
	// NOT
	"OpNot" : 19,
	// LEN
	"OpLen" : 20,
	// CONCAT
	"OpConcat" : 21,
	// JMP
	"OpJmp" : 22,
	// EQ
	"OpEqB" : 23,
	"OpEqC" : 23,
	"OpEqBC" : 23,
	"OpEq" : 23,
	// LT
	"OpLt" : 24,
	"OpLtB" : 24,
	"OpLtC" : 24,
	"OpLtBC" : 24,

	"OpGe" : 24,
	"OpGeB" : 24,
	"OpGeC" : 24,
	"OpGeBC" : 24,
	// LE
	"OpLe" : 25,
	"OpLeB" : 25,
	"OpLeC" : 25,
	"OpLeBC" : 25,
	// TEST
	"OpTest" : 26,
	"OpTestC" : 26,
	// TESTSET
	"OpTestSet" : 27,
	"OpTestSetC" : 27,
	// CALL
	"OpCallC1B2" : 28,
	"OpCallB1C2" : 28,
	"OpCallC0" : 28,
	"OpCallC0B2" : 28,
	"OpCallB1C0" : 28,
	"OpCallB0" : 28,
	"OpCallB0C0" : 28,
	"OpCallB1" : 28,
	"OpCallC1" : 28,
	"OpCallB1C1" : 28,
	"OpCallC2" : 28,
	"OpCallC2B2" : 28,
	"OpCallB0C2" : 28,
	"OpCall" : 28,
	"OpCallB2" : 28,
	"OpCallB0C1" : 28,
	// TAILCALL
	"OpTailCall" : 29,
	"OpTailCallB0" : 29,
	"OpTailCallB1" : 29,
	// RETURN
	"OpReturnB1" : 30,
	"OpReturn" : 30,
	"OpReturnB2" : 30,
	"OpReturnB3" : 30,
	"OpReturnB0" : 30,
	// FORLOOP
	"OpForLoop" : 31,
	// FORPREP
	"OpForPrep" : 32,
	// TFORLOOP
	"OpTForLoop" : 33,
	// SETLIST
	"OpSetList" : 34,
	"OpSetListB0" : 34,
	"OpSetListC0" : 34,
	// CLOSE
	"OpClose" : 35,
	// CLOSURE
	"OpClosure" : 36,
	"OpClosureNU" : 36,
	// VARARG
	"OpVarArg" : 37,
	"OpVarArgB0" : 37,
}

// VOpToOp turns ib2 vopcode string to lua opcode integer
func VOpToOp(vop string) int {
	return voptoop[vop]
}
