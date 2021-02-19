package deobfuscator

import (
	"../beautifier"
	"strconv"
	"io"
	"fmt"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"github.com/yuin/gopher-lua/ast"
)

type settings struct {
	BytecodeCompress bool
	PreserveLineInfo bool
}

type obfuscator struct {
	Name string
	FuncLen int
	ChunkLen int
	WrapLen int
	WrapPos int
	BytePos int
}

type vmdata struct {
	Pos int
	Chunk []ast.Stmt
	Obfuscator obfuscator//interface{}
	Functions []*ast.FunctionExpr
	Settings settings
	Bytecode []byte
	Key byte
	Bool int 
	Float int
	String int
	Order []int // 1: Arg 2: Protos 3: Instructions
	Env string
	Upvalues string
	Stack string
	Instruction string
	PC string
	VMLoop *ast.WhileStmt
}

func (vm *vmdata) getFunctions() {
	for _, stmt := range vm.Chunk {
		if local, ok := stmt.(*ast.LocalAssignStmt); ok {
			if fExpr, ok := local.Exprs[0].(*ast.FunctionExpr); ok {
				vm.Functions = append(vm.Functions, fExpr)
			}
		}
	}
}

func (vm *vmdata) getSettings() {
	switch len(vm.Functions) {
	case vm.Obfuscator.FuncLen:
		vm.Settings.PreserveLineInfo = true
	case vm.Obfuscator.FuncLen+1:
		vm.Settings.BytecodeCompress = true
	default:
		fmt.Print(len(vm.Functions))
		panic("Couldn't detect obfuscation settings.")
	}

	switch len(vm.Chunk) {
	case vm.Obfuscator.ChunkLen:
		vm.Settings.PreserveLineInfo = false
	case vm.Obfuscator.ChunkLen+1:
	case vm.Obfuscator.ChunkLen+2:
		vm.Settings.PreserveLineInfo = true
		vm.Settings.BytecodeCompress = true
	default:
		fmt.Print(len(vm.Chunk))
		panic("Couldn't detect obfuscation settings.")
	}
}

func (vm *vmdata) getBytecode() {
	msg := "Couldn't get the bytecode."
	if vm.Settings.BytecodeCompress {
		compressed := vm.Chunk[vm.Obfuscator.BytePos+1]
		local := getLocal(compressed, msg)
		call := getCall(local.Exprs[0], msg)
		if str, ok := call.Args[0].(*ast.StringExpr); ok {
			vm.Bytecode = Decompress(str.Value)
		}
	} else {
		bytestring := vm.Chunk[vm.Obfuscator.BytePos]
		if local, ok := bytestring.(*ast.AssignStmt); ok {
			if str, ok := local.Rhs[0].(*ast.StringExpr); ok {
				vm.Bytecode = []byte(str.Value)
			}
		}
	}
	if len(vm.Bytecode) == 0 {
		panic("Couldn't get the bytecode.")
	}
}

func (vm *vmdata) getKey() {
	var gBits32 *ast.FunctionExpr
	var success bool

	if vm.Settings.BytecodeCompress {
		gBits32 = vm.Functions[2]
	} else {
		gBits32 = vm.Functions[1]
	}

	if len(gBits32.Stmts) == 7 {
		if xor, ok := gBits32.Stmts[1].(*ast.AssignStmt); ok {
			if call, ok := xor.Rhs[0].(*ast.FuncCallExpr); ok {
				if key, ok := call.Args[1].(*ast.NumberExpr); ok {
					key, err := strconv.Atoi(key.Value)
					if err == nil {
						success = true
						vm.Key = byte(key)
					}
				}
			}
		}
	} 
	if !success {
		panic("Invalid gBits32 function.")
	}
}

func getKeys(stmt *ast.IfStmt, keys *[]int) (success bool) {

	if operator, ok := stmt.Condition.(*ast.RelationalOpExpr); ok {
		if operator.Operator == "==" {
			if key, ok := operator.Rhs.(*ast.NumberExpr);  ok {
				key, err := strconv.Atoi(key.Value)
				if err == nil {
					success = true
					*keys = append(*keys, key)
				}
			}
		}
	}

	if !success {
		panic("Couldn't retrieve constant decryption keys.")
	}

	if len(stmt.Else) == 1 {
		if stmt, ok := stmt.Else[0].(*ast.IfStmt); ok {
			success = getKeys(stmt, keys)
		}
	}
	return success
}

func (vm *vmdata) getConstKeys() {
	var Deserialize *ast.FunctionExpr
	var success bool
	var keys []int

	if vm.Settings.BytecodeCompress {
		Deserialize = vm.Functions[8]
	} else {
		Deserialize = vm.Functions[7]
	}

	if len(Deserialize.Stmts) > 8 {
		if loop, ok := Deserialize.Stmts[6].(*ast.NumberForStmt); ok {
			if ifstmt, ok := loop.Stmts[2].(*ast.IfStmt); ok {
				success = getKeys(ifstmt, &keys)
			}
		}
	}

	if !success {
		panic("Invalid Deserialize function.")
	}

	vm.Bool = keys[0]
	vm.Float = keys[1]
	vm.String = keys[2]
}

func (vm *vmdata) getOrder() {
	var Deserialize *ast.FunctionExpr
	success := true

	if vm.Settings.BytecodeCompress {
		Deserialize = vm.Functions[8]
	} else {
		Deserialize = vm.Functions[7]
	}

	for _, ex := range Deserialize.Stmts[7:len(Deserialize.Stmts)-1] {
		switch st := ex.(type) {
		case *ast.NumberForStmt:
			switch len(st.Stmts) {
			case 1:
				vm.Order = append(vm.Order, 2)
			case 2:
				vm.Order = append(vm.Order, 3)
			default:
				success = false
				break
			}
		case *ast.AssignStmt:
			vm.Order = append(vm.Order, 1)
		default:
			success = false
			break
		}
	}
	if !success {
		panic("Failed to get the deserializing order.")
	}
}


func getPC(ex *ast.FunctionExpr, idx int) string {
	if pc, ok := ex.Stmts[idx].(*ast.LocalAssignStmt); ok {
		return pc.Names[0]
	}
	panic("Couldn't get the PC variable.")
}

func getLocal(stmt ast.Stmt, msg string) *ast.LocalAssignStmt {
	if local, ok := stmt.(*ast.LocalAssignStmt); ok {
		return local
	}
	panic(msg)
}

func getReturnStmt(stmt ast.Stmt, msg string) *ast.ReturnStmt {
	if returnStmt, ok := stmt.(*ast.ReturnStmt); ok {
		return returnStmt
	}
	panic(msg)
}

func getFunction(expr ast.Expr, msg string) *ast.FunctionExpr {
	if function, ok := expr.(*ast.FunctionExpr); ok {
		return function
	}
	panic(msg)
}

func getWhile(stmt ast.Stmt, msg string) *ast.WhileStmt {
	if while, ok := stmt.(*ast.WhileStmt); ok {
		return while
	}
	panic(msg)
}

func getCall(expr ast.Expr, msg string) *ast.FuncCallExpr {
	if call, ok := expr.(*ast.FuncCallExpr); ok {
		return call
	}
	panic(msg)
}

func (vm *vmdata) getVariables() {
	var Wrap *ast.FunctionExpr
	var Inner *ast.FunctionExpr
	stackIdx := 10 // 8
	instIdx := 13 // 11
	loopIdx := 15 // 13

	if vm.Settings.BytecodeCompress {
		Wrap = vm.Functions[9]
	} else {
		Wrap = vm.Functions[8]
	}

	if len(Wrap.Stmts) != vm.Obfuscator.WrapLen {panic("Invalid Wrap function length.")}

	returnStmt := getReturnStmt(Wrap.Stmts[vm.Obfuscator.WrapPos], "Missing return statement in Wrap function.")
	Inner = getFunction(returnStmt.Exprs[0], "Missing inner function in Wrap.")

	if vm.Settings.PreserveLineInfo {
		pcall := getLocal(Inner.Stmts[4], "Invalid inner function format. (PreserveLineinfo)")
		Inner = getFunction(pcall.Exprs[0], "Loop function not defined. (PreserveLineinfo)")
		stackIdx = 7
		instIdx = 10
		loopIdx = 12
	}

	if len(Inner.Stmts) != 16 {
		panic("Invalid inner function inside Wrap.")
	}

	names := Wrap.ParList.Names
	vm.Upvalues = names[1]
	vm.Env = names[2]
	vm.PC = getPC(Inner, 4)
	vm.VMLoop = getWhile(Inner.Stmts[loopIdx], "Missing vmloop.")
	vm.Stack = getLocal(Inner.Stmts[stackIdx], "Couldn't get the stack variable.").Names[0]
	vm.Instruction = getLocal(Inner.Stmts[instIdx], "Couldn't get the instruction variable.").Names[0]
}

func (vm *vmdata) getObfuscator() {
	if len(vm.Chunk) == 1 {
		vm.Obfuscator = obfuscator{
			Name: "Aztup Brew",
			FuncLen: 9,
			ChunkLen: 24,
			BytePos: 8,
			WrapPos: 1,
			WrapLen: 2,
		}
		returnStmt := getReturnStmt(vm.Chunk[0], "Bad vm")
		call := getCall(returnStmt.Exprs[0], "No call")
		function := getFunction(call.Func, "Worse vm")
		vm.Chunk = function.Stmts
	} else {
		vm.Obfuscator = obfuscator{
			Name: "Vanilla IronBrew",
			FuncLen: 9,
			ChunkLen: 26,
			BytePos: 11,
			WrapPos: 3,
			WrapLen: 4,
		}
	}
}

// DeobfuscateIronbrew returns the function proto of the deobfuscated ironbrew script.
func DeobfuscateIronbrew(file io.Reader, hashmap map[string]string, debug bool) *lua.FunctionProto {
	ast, err := parse.Parse(file, "")
	if err != nil {
		panic(err)
	}

	beautifier.Optimize(ast)
	
	vm := &vmdata{Chunk: ast}
	vm.getObfuscator()
	vm.getFunctions()
	vm.getSettings()
	vm.getBytecode()
	vm.getKey()
	vm.getConstKeys()
	vm.getOrder()
	vm.getVariables()

	if debug {
		fmt.Println("Obfuscator:", vm.Obfuscator.Name)
		fmt.Printf("Obfuscation settings: %+v\n", vm.Settings)
		fmt.Println("Key:",vm.Key)
		fmt.Println("Bool:",vm.Bool)
		fmt.Println("Float:",vm.Float)
		fmt.Println("String:",vm.String)
		fmt.Println("Order:", vm.Order)
		fmt.Println("Upvalues:", vm.Upvalues)
		fmt.Println("Environment:", vm.Env)
		fmt.Println("Stack:", vm.Stack)
		fmt.Println("Instruction:", vm.Instruction)
		fmt.Println("PC:", vm.PC)
		fmt.Println("VM loop found:", vm.VMLoop != nil)
		fmt.Println("Bytecode:", vm.Bytecode)
	}

	if vm.Settings.PreserveLineInfo {
		panic("PreserverLineInfo is not supported yet.")
	}

	opmap := generateOpcodemap(vm, hashmap)
	//fmt.Println(opmap)
	return vm.deserialize(opmap)
}