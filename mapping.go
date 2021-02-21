package deobfuscator

import (
	"errors"
	"strconv"
	"strings"

	"github.com/notnoobmaster/beautifier"
	"github.com/notnoobmaster/deobfuscator/obfuscators/ironbrew/opcodemap"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

type mapData struct {
	Delimiter   string
	Variables   []string
	Opcodemap   []*opcodemap.Instruction
	Hashmap     map[string]func(opcodemap.Instruction)
}

func (data *mapData) solveSuperOp(chunk []ast.Stmt) (*opcodemap.Instruction, error) {
	pos := 0
	for pos < len(chunk) {
		if _, ok := chunk[pos].(*ast.LocalAssignStmt); !ok {
			break
		}
		pos++
	}
	pattern := beautifier.GeneratePattern(chunk[pos:], data.Variables)
	hashes := strings.Split(pattern, data.Delimiter)

	for _, hash := range hashes {
		if create, ok := data.Hashmap[hash]; ok {
			// add to something
		}
		return errors.New("shit hit the fan")
	}
	return nil
}

func (data *mapData) chunkToOp(chunk []ast.Stmt) (*opcodemap.Instruction, error) {
	hash := beautifier.GeneratePattern(chunk, data.Variables)
	if _, ok := data.Hashmap[hash]; ok {
		return data.solveSuperOp(chunk)
	}
	inst := opcodemap.Instruction{}
	inst.Create = data.Hashmap[hash]
	return &inst, nil
}

func (data *mapData) solveIf(stmt *ast.IfStmt) error {
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
		op1, err := data.chunkToOp(stmt.Then)
		if !ok {
			return err
		}
		op2, err := data.chunkToOp(stmt.Else)
		if !ok {
			return err
		}
		data.Opcodemap[op] = op1
		data.Opcodemap[op+1] = op2
	case ">":
		op1, err := data.chunkToOp(stmt.Else)
		if !ok {
			return err
		}
		op2, err := data.chunkToOp(stmt.Then)
		if !ok {
			return err
		}
		data.Opcodemap[op] = op1
		data.Opcodemap[op+1] = op2
	case "<=":
		if inner, ok := stmt.Then[0].(*ast.IfStmt); ok && len(stmt.Then) == 1 {
			data.solveIf(inner)
		} else {
			return errors.New("Malformed if statement, expected elseif")
		}
	default:
		return errors.New("Malformed if statement, invalid relational operator")
	}
	return nil
}

func generateOpcodemap(vm *vmdata, hashmap map[string]opcodemap.CreateSig) map[int]opcodemap.CreateSig {
	data := mapData{
		Variables: []string{vm.Stack, vm.Instruction, vm.Env, vm.Upvalues, vm.PC},
		Hashmap: hashmap,
	}
	return data.solveIf(vm.VMLoop.Stmts)
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