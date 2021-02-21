package ironbrew

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
	Hashmap     map[string]func(*opcodemap.Instruction)uint32
}

func (data *mapData) solveSuperOp(chunk []ast.Stmt) (*opcodemap.Instruction, error) {
	pos := 0
	for pos < len(chunk) {
		if _, ok := chunk[pos].(*ast.LocalAssignStmt); !ok {
			break
		}
		pos++
	}
	pattern := beautifier.GenerateHash(chunk[pos:], data.Variables)
	hashes := strings.Split(pattern, data.Delimiter)

	instruction := opcodemap.Instruction{}
	superop := opcodemap.SuperOperator{}

	instruction.IsSuperop = true
	instruction.Superop = superop

	for _, hash := range hashes {
		if create, ok := data.Hashmap[hash]; ok {
			inst := opcodemap.Instruction{Func: create}
			superop.Instructions = append(superop.Instructions, &inst)
		}
		return nil, errors.New("shit hit the fan")
	}
	return &instruction, nil
}

func (data *mapData) chunkToOp(chunk []ast.Stmt) (*opcodemap.Instruction, error) {
	hash := beautifier.GenerateHash(chunk, data.Variables)
	if _, ok := data.Hashmap[hash]; ok {
		return data.solveSuperOp(chunk)
	}
	inst := opcodemap.Instruction{}
	inst.Func = data.Hashmap[hash]
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

// GenerateOpcodemap solves the vm loop and generates a lookup table for vm opcodes to instruction creation functions
func GenerateOpcodemap(stmt *ast.IfStmt, variables []string, hashmap map[string]func(*opcodemap.Instruction)uint32) ([]*opcodemap.Instruction, error) {
	data := mapData{
		Variables: variables,
		Hashmap: hashmap,
	}
	err := data.solveIf(stmt)
	if err != nil {
		return nil, err
	}
	return data.Opcodemap, nil
}

func initMapping() (map[string]func(*opcodemap.Instruction)uint32, error) {
	// We need to detect some variable names or else some opcodes have the same hash.
	variables := []string{"Stk", "Inst", "Env", "Upvalues", "InstrPoint",}
	replace := map[string]byte{
		"OP_A": beautifier.NumberExpr, 
		"OP_B": beautifier.NumberExpr, 
		"OP_C": beautifier.NumberExpr, 
		"OP_ENUM": beautifier.NumberExpr, 
		"OP_MOVE": beautifier.NumberExpr,
	}
	hashmap := make(map[string]func(*opcodemap.Instruction)uint32)

	for str, function := range opcodemap.OpCodes {
		chunk, err := parse.Parse(strings.NewReader(str), "")
		if err != nil {
			return nil, errors.New("Parsing somehow fucked up when generating the hashmap!\nStr:\n"+str)
		}

		hash := beautifier.GenerateHashWithReplace(chunk, variables, replace)
		// Making sure that we accidentally don't have duplicate hashes. 
		if _, ok := hashmap[hash]; ok {
			return nil, errors.New("Same Hash\n" + str)
		}

		hashmap[hash] = function
	}
	return hashmap, nil
}