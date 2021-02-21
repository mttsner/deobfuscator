package opcodemap

const (
	opMOVE int = iota
	opLOADK
	opLOADBOOL
	opLOADNIL
	opGETUPVAL
	opGETGLOBAL
	opGETTABLE
	opSETGLOBAL
	opSETUPVAL
	opSETTABLE
	opNEWTABLE
	opSELF
	opADD
	opSUB
	opMUL
	opDIV
	opMOD
	opPOW
	opUNM
	opNOT
	opLEN
	opCONCAT
	opJMP
	opEQ
	opLT
	opLE
	opTEST
	opTESTSET
	opCALL
	opTAILCALL
	opRETURN
	opFORLOOP
	opFORPREP
	opTFORLOOP
	opSETLIST
	opCLOSE
	opCLOSURE
	opVARARG
)

var OpCodes = map[string]func(*Instruction)uint32 {
	// MOVE
	strMove: (*Instruction).createMove,
	// LOADK
	strLoadK: (*Instruction).createLoadK,
	// LOADBOOL
	strLoadBool:  (*Instruction).createLoadBool,
	strLoadBoolC: (*Instruction).createLoadBoolC,
	// LOADNIL
	strLoadNil: (*Instruction).createLoadNil,
	// GETUPVAL
	strGetUpval: (*Instruction).createGetUpval,
	// GETGLOBAL
	strGetGlobal: (*Instruction).createGetGlobal,
	// GETTABLE
	strGetTable:      (*Instruction).createGetTable,
	strGetTableConst: (*Instruction).createGetTableConst,
	// SETUPVAL
	strSetUpval: (*Instruction).createSetUpval,
	// SETTABLE
	strSetTable:   (*Instruction).createSetTable,
	strSetTableB:  (*Instruction).createSetTableB,
	strSetTableC:  (*Instruction).createSetTableC,
	strSetTableBC: (*Instruction).createSetTableBC,
	// NEWTABLE
	strNewTableB0: (*Instruction).createNewTableB0,
	// SELF
	strSelf:  (*Instruction).createSelf,
	strSelfC: (*Instruction).createSelfC,
	// ADD
	strAdd:   (*Instruction).createAdd,
	strAddB:  (*Instruction).createAddB,
	strAddC:  (*Instruction).createAddC,
	strAddBC: (*Instruction).createAddBC,
	// SUB
	strSub:   (*Instruction).createSub,
	strSubB:  (*Instruction).createSubB,
	strSubC:  (*Instruction).createSubC,
	strSubBC: (*Instruction).createSubBC,
	// MUL
	strMul:   (*Instruction).createMul,
	strMulB:  (*Instruction).createMulB,
	strMulC:  (*Instruction).createMulC,
	strMulBC: (*Instruction).createMulBC,
	// DIV
	strDiv:   (*Instruction).createDiv,
	strDivB:  (*Instruction).createDivB,
	strDivC:  (*Instruction).createDivC,
	strDivBC: (*Instruction).createDivBC,
	// MOD
	strMod:   (*Instruction).createMod,
	strModB:  (*Instruction).createModB,
	strModC:  (*Instruction).createModC,
	strModBC: (*Instruction).createModBC,
	// POW
	strPow:   (*Instruction).createPow,
	strPowB:  (*Instruction).createPowB,
	strPowC:  (*Instruction).createPowC,
	strPowBC: (*Instruction).createPowBC,
	// UNM
	strUnm: (*Instruction).createUnm,
	// NOT
	strNot: (*Instruction).createNot,
	// LEN
	strLen: (*Instruction).createLen,
	// CONCAT
	strConcat: (*Instruction).createConcat,
	// JMP
	strJmp: (*Instruction).createJmp,
	// EQ
	strEq:   (*Instruction).createEq,
	strEqB:  (*Instruction).createEqB,
	strEqC:  (*Instruction).createEqC,
	strEqBC: (*Instruction).createEqBC,
	// LT
	strLt:   (*Instruction).createLt,
	strLtB:  (*Instruction).createLtB,
	strLtC:  (*Instruction).createLtC,
	strLtBC: (*Instruction).createLtBC,
	// LE
	strLe:   (*Instruction).createLe,
	strLeB:  (*Instruction).createLeB,
	strLeC:  (*Instruction).createLeC,
	strLeBC: (*Instruction).createLeBC,
	// TEST
	strTest:  (*Instruction).createTest,
	strTestC: (*Instruction).createTestC,
	// TESTSET
	strTestSet:  (*Instruction).createTestSet,
	strTestSetC: (*Instruction).createTestSetC,
	// Call
	strCall:     (*Instruction).createCall,
	strCallB2:   (*Instruction).createCallB2,
	strCallB0:   (*Instruction).createCallB0,
	strCallB1:   (*Instruction).createCallB1,
	strCallC0:   (*Instruction).createCallC0,
	strCallC0B2: (*Instruction).createCallC0B2,
	strCallC1:   (*Instruction).createCallC1,
	strCallC1B2: (*Instruction).createCallC1B2,
	strCallB0C0: (*Instruction).createCallB0C0,
	strCallB0C1: (*Instruction).createCallB0C1,
	strCallB1C0: (*Instruction).createCallB1C0,
	strCallB1C1: (*Instruction).createCallB1C1,
	strCallC2:   (*Instruction).createCallC2,
	strCallC2B2: (*Instruction).createCallC2B2,
	strCallB0C2: (*Instruction).createCallB0C2,
	strCallB1C2: (*Instruction).createCallB1C2,
	// TAILCALL
	strTailCall:   (*Instruction).createTailCall,
	strTailCallB0: (*Instruction).createTailCallB0,
	strTailCallB1: (*Instruction).createTailCallB1,
	// RETURN
	strReturn:   (*Instruction).createReturn,
	strReturnB0: (*Instruction).createReturnB0,
	strReturnB1: (*Instruction).createReturnB1,
	strReturnB2: (*Instruction).createReturnB2,
	strReturnB3: (*Instruction).createReturnB3,
	// FORLOOP
	strForLoop: (*Instruction).createForLoop,
	// FORPREP
	strForPrep: (*Instruction).createForPrep,
	// TFORLOOP
	strTForLoop: (*Instruction).createTForLoop,
	// SETLIST
	strSetList:   (*Instruction).createSetList,
	strSetListB0: (*Instruction).createSetListB0,
	strSetListC0: (*Instruction).createSetListC0,
	// CLOSE
	strClose: (*Instruction).createClose,
	// CLOSURE
	strClosure:   (*Instruction).createClosure,
	strClosureNU: (*Instruction).createClosureNU,
	// VARARG
	strVarArg:   (*Instruction).createVarArg,
	strVarArgB0: (*Instruction).createVarArgB0,
}

/*/ OpCodes maps vm function strings to functions that create the correct opcode representation of that function.
var OpCodes = map[string]string {
	// MOVE
	strMove: "(*Instruction).createMove",
	// LOADK
	strLoadK: "(*Instruction).createLoadK",
	// LOADBOOL
	strLoadBool:  "(*Instruction).createLoadBool",
	strLoadBoolC: "(*Instruction).createLoadBoolC",
	// LOADNIL
	strLoadNil: "(*Instruction).createLoadNil",
	// GETUPVAL
	strGetUpval: "(*Instruction).createGetUpval",
	// GETGLOBAL
	strGetGlobal: "(*Instruction).createGetGlobal",
	// GETTABLE
	strGetTable:      "(*Instruction).createGetTable",
	strGetTableConst: "(*Instruction).createGetTableConst",
	// SETUPVAL
	strSetUpval: "(*Instruction).createSetUpval",
	// SETTABLE
	strSetTable:   "(*Instruction).createSetTable",
	strSetTableB:  "(*Instruction).createSetTableB",
	strSetTableC:  "(*Instruction).createSetTableC",
	strSetTableBC: "(*Instruction).createSetTableBC",
	// NEWTABLE
	strNewTableB0: "(*Instruction).createNewTableB0",
	// SELF
	strSelf:  "(*Instruction).createSelf",
	strSelfC: "(*Instruction).createSelfC",
	// ADD
	strAdd:   "(*Instruction).createAdd",
	strAddB:  "(*Instruction).createAddB",
	strAddC:  "(*Instruction).createAddC",
	strAddBC: "(*Instruction).createAddBC",
	// SUB
	strSub:   "(*Instruction).createSub",
	strSubB:  "(*Instruction).createSubB",
	strSubC:  "(*Instruction).createSubC",
	strSubBC: "(*Instruction).createSubBC",
	// MUL
	strMul:   "(*Instruction).createMul",
	strMulB:  "(*Instruction).createMulB",
	strMulC:  "(*Instruction).createMulC",
	strMulBC: "(*Instruction).createMulBC",
	// DIV
	strDiv:   "(*Instruction).createDiv",
	strDivB:  "(*Instruction).createDivB",
	strDivC:  "(*Instruction).createDivC",
	strDivBC: "(*Instruction).createDivBC",
	// MOD
	strMod:   "(*Instruction).createMod",
	strModB:  "(*Instruction).createModB",
	strModC:  "(*Instruction).createModC",
	strModBC: "(*Instruction).createModBC",
	// POW
	strPow:   "(*Instruction).createPow",
	strPowB:  "(*Instruction).createPowB",
	strPowC:  "(*Instruction).createPowC",
	strPowBC: "(*Instruction).createPowBC",
	// UNM
	strUnm: "(*Instruction).createUnm",
	// NOT
	strNot: "(*Instruction).createNot",
	// LEN
	strLen: "(*Instruction).createLen",
	// CONCAT
	strConcat: "(*Instruction).createConcat",
	// JMP
	strJmp: "(*Instruction).createJmp",
	// EQ
	strEq:   "(*Instruction).createEq",
	strEqB:  "(*Instruction).createEqB",
	strEqC:  "(*Instruction).createEqC",
	strEqBC: "(*Instruction).createEqBC",
	// LT
	strLt:   "(*Instruction).createLt",
	strLtB:  "(*Instruction).createLtB",
	strLtC:  "(*Instruction).createLtC",
	strLtBC: "(*Instruction).createLtBC",
	// LE
	strLe:   "(*Instruction).createLe",
	strLeB:  "(*Instruction).createLeB",
	strLeC:  "(*Instruction).createLeC",
	strLeBC: "(*Instruction).createLeBC",
	// TEST
	strTest:  "(*Instruction).createTest",
	strTestC: "(*Instruction).createTestC",
	// TESTSET
	strTestSet:  "(*Instruction).createTestSet",
	strTestSetC: "(*Instruction).createTestSetC",
	// Call
	strCall:     "(*Instruction).createCall",
	strCallB2:   "(*Instruction).createCallB2",
	strCallB0:   "(*Instruction).createCallB0",
	strCallB1:   "(*Instruction).createCallB1",
	strCallC0:   "(*Instruction).createCallC0",
	strCallC0B2: "(*Instruction).createCallC0B2",
	strCallC1:   "(*Instruction).createCallC1",
	strCallC1B2: "(*Instruction).createCallC1B2",
	strCallB0C0: "(*Instruction).createCallB0C0",
	strCallB0C1: "(*Instruction).createCallB0C1",
	strCallB1C0: "(*Instruction).createCallB1C0",
	strCallB1C1: "(*Instruction).createCallB1C1",
	strCallC2:   "(*Instruction).createCallC2",
	strCallC2B2: "(*Instruction).createCallC2B2",
	strCallB0C2: "(*Instruction).createCallB0C2",
	strCallB1C2: "(*Instruction).createCallB1C2",
	// TAILCALL
	strTailCall:   "(*Instruction).createTailCall",
	strTailCallB0: "(*Instruction).createTailCallB0",
	strTailCallB1: "(*Instruction).createTailCallB1",
	// RETURN
	strReturn:   "(*Instruction).createReturn",
	strReturnB0: "(*Instruction).createReturnB0",
	strReturnB1: "(*Instruction).createReturnB1",
	strReturnB2: "(*Instruction).createReturnB2",
	strReturnB3: "(*Instruction).createReturnB3",
	// FORLOOP
	strForLoop: "(*Instruction).createForLoop",
	// FORPREP
	strForPrep: "(*Instruction).createForPrep",
	// TFORLOOP
	strTForLoop: "(*Instruction).createTForLoop",
	// SETLIST
	strSetList:   "(*Instruction).createSetList",
	strSetListB0: "(*Instruction).createSetListB0",
	strSetListC0: "(*Instruction).createSetListC0",
	// CLOSE
	strClose: "(*Instruction).createClose",
	// CLOSURE
	strClosure:   "(*Instruction).createClosure",
	strClosureNU: "(*Instruction).createClosureNU",
	// VARARG
	strVarArg:   "(*Instruction).createVarArg",
	strVarArgB0: "(*Instruction).createVarArgB0",
}
*/