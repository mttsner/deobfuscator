package deobfuscator

import (
	"errors"
	"strings"
	"strconv"

	"github.com/notnoobmaster/beautifier"
	"github.com/notnoobmaster/deobfuscator/opcodemap"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

type mapData struct {
	Variables   map[string]byte
	Opcodemap   map[int][]string
	Hashmap     map[string]opcodemap.CreateSig
}

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

func (s *state) solveIf(stmt *ast.IfStmt) error {
	relational, ok := stmt.Condition.(*ast.RelationalOpExpr)
	if !ok {
		return errors.New("Malformed if statement, no relational operator")
	}
	number, ok := relational.Rhs.(*ast.NumberExpr)
	if !ok {
		return errors.New("Malformed if statement, right side is not a number")
	}
	op, err := strconv.Atoi(number.Value)
	if err != nil {
		return errors.New("Malformed if statement, string conversion failed")
	}
	switch relational.Operator {
	case "==":
		s.Opcodemap[op] = s.chunkToOp(stmt.Then)
		s.Opcodemap[op+1] = s.chunkToOp(stmt.Else)
	case ">":
		s.Opcodemap[op+1] = s.chunkToOp(stmt.Then)
		s.Oopcodemap[op] = s.chunkToOp(stmt.Else)
	case "<=":
		if inner, ok := stmt.Then[0].(*ast.IfStmt); ok && len(stmt.Then) == 1 {
			s.solveIf(inner)
		} else {
			return errors.New("Malformed if statement, expected elseif")
		}
	default:
		return errors.New("Malformed if statement, invalid relational operator")
	}
	return nil
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