package deobfuscator

import (
	"strings"
	"github.com/notnoobmaster/beautifier"
	"github.com/notnoobmaster/deobfuscator/opcodemap"
	"github.com/yuin/gopher-lua/parse"
	//"github.com/yuin/gopher-lua/ast"
)

type mapData struct {
	Variables   map[string]byte
	Opcodemap   map[int][]string
	Hashmap     map[string]opcodemap.CreateSig
}
/*
func (data *mapData) solveSuperOp(chunk []ast.Stmt) []string {
	localassing := identExpr+"="
	hash := ""
	for _,exp := range chunk {
		temp := data.compileStmt(exp)
		if data.Hashmap[temp] == "Delimiter" {
			ops = append(ops,data.Hashmap[hash])
			hash = ""
			continue
		}
		hash += temp
	}
	return 
}

func (data *mapData) chunkToOp(chunk []ast.Stmt) opcodemap.CreateSig {
	hash := beautifier.GeneratePattern(chunk, data.Variables)
	//if _, ok := data.Hashmap[hash]; ok {
	//	return data.solveSuperOp(chunk)
	//}
	return data.Hashmap[hash]
}

func solveIf(chunk []ast.Stmt) {
	for _, stmt := range chunk {
		switch st := stmt.(type) {
		case *ast.IfStmt:
			ex := st.Condition.(*ast.RelationalOpExpr)
			switch ex.Operator {
			case "==":
				op, _ := strconv.Atoi(ex.Rhs.(*ast.NumberExpr).Value)
				data.Opcodemap[op] = data.chunkToOp(st.Then)
				data.Opcodemap[op+1] = data.chunkToOp(st.Else)
			case ">":
				op, _ := strconv.Atoi(ex.Rhs.(*ast.NumberExpr).Value)
				data.Opcodemap[op+1] = data.chunkToOp(st.Then)
				data.Opcodemap[op] = data.chunkToOp(st.Else)
			case "<=":
				if inner, ok := st.Then[0].(*ast.IfStmt); ok {
					if op, ok := inner.Condition.(*ast.RelationalOpExpr); ok {
						if _, ok := op.Rhs.(*ast.NumberExpr); ok {
							//op, _ := strconv.Atoi(rhs.Value)
							data.solveIf(st.Then)
							data.solveIf(st.Else)
							return 
						}
					}
				}
				op, _ := strconv.Atoi(ex.Rhs.(*ast.NumberExpr).Value)
				data.Opcodemap[op] = data.chunkToOp(st.Then)
				data.solveIf(st.Else)
			default:
				data.solveIf(st.Then)
			}
		}
	}
}

func generateOpcodemap(vm *vmdata, hashmap map[string]opcodemap.CreateSig) map[int]opcodemap.CreateSig {
	variables := map[string]byte{
		vm.Stack :        beautifier.Stack,
		vm.Instruction :       beautifier.Instruction,
		vm.Env:        beautifier.Environment,
		vm.Upvalues:   beautifier.Upvalues,
		vm.PC: beautifier.Pointer,
	}
	solveIf(vm.VMLoop.Stmts)
}
*/
// GenerateHashmap generates the lookup table for opcode functions.
func GenerateHashmap() map[string]func(*opcodemap.Instruction) {
	// We need to detect some variable names or else some opcodes have the same hash.
	variables := map[string]byte{
		"Stk":        beautifier.Stack,
		"Inst":       beautifier.Instruction,
		"Env":        beautifier.Environment,
		"Upvalues":   beautifier.Upvalues,
		"InstrPoint": beautifier.Pointer,
		"OP_A":       beautifier.A,
		"OP_B":       beautifier.B,
		"OP_C":       beautifier.C,
		"OP_ENUM":    beautifier.MOVE,
		"OP_MOVE":    beautifier.ENUM,
	}

	opMap := make(map[string]opcodemap.CreateSig)

	for str, function := range opcodemap.OpCodes {
		chunk, err := parse.Parse(strings.NewReader(str), "")
		if err != nil {
			panic("Parsing somehow fucked up when generating the hashmap!")
		}

		hash := beautifier.GeneratePattern(chunk, variables)
		// Making sure that we accidentally don't have duplicate hashes. 
		if _, ok := opMap[hash]; ok {
			panic("Same Hash\n" + str)
		}

		opMap[hash] = function
	}
	return opMap
}