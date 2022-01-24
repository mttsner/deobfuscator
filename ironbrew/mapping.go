package ironbrew

import (
	"errors"
	"strings"

	"github.com/notnoobmaster/deobfuscator/helper"
	"github.com/notnoobmaster/deobfuscator/ironbrew/opcodemap"

	"github.com/notnoobmaster/luautil/ast"
	"github.com/notnoobmaster/luautil/parse"
)

type mapData struct {
	Variables   []string
	Opcodemap   map[int][]*opcodemap.Instruction
	Hashmap     map[string]func(*opcodemap.Instruction)uint32
}

// Generated during the initialization phase
var delimiter string

func (data *mapData) solveSuperOp(chunk []ast.Stmt) (instructions []*opcodemap.Instruction, err error) {
	// Remove local variables generated at the start of each superoperator
	for i, l := 0, len(chunk); i < l; i++ {
		if _, ok := chunk[i].(*ast.LocalAssignStmt); !ok {
			chunk = chunk[i:]
			break
		}
	}
	// Generated signatrue for the superoperator
	pattern := helper.GenerateSignature(chunk, data.Variables)
	// Split signature by the delimiter that is between each instruction in the superoperator
	signatures := strings.Split(pattern, delimiter)
	// Create all the instructions in the superoperator
	for _, hash := range signatures {
		if create, ok := data.Hashmap[hash]; ok {
			instructions = append(instructions, &opcodemap.Instruction{Func: create})
			continue
		}
		return nil, errors.New("superoperator contained an unknown signature")
	}

	return instructions, nil
}

func (data *mapData) chunkToOp(chunk []ast.Stmt) ([]*opcodemap.Instruction, error) {
	hash := helper.GenerateSignature(chunk, data.Variables)
	if create, ok := data.Hashmap[hash]; ok {
		inst := &opcodemap.Instruction{Func: create}
		return []*opcodemap.Instruction{inst}, nil
		
	}
	return data.solveSuperOp(chunk)
}

func (data *mapData) traverseExecLoop(stmt *ast.IfStmt) (err error) {
	relational, ok := stmt.Condition.(*ast.RelationalOpExpr)
	if !ok {
		return errors.New("malformed if statement, no relational operator")
	}
	number, ok := relational.Rhs.(*ast.NumberExpr)
	if !ok {
		return errors.New("malformed if statement, right side is not a number")
	}

	op := int(number.Value)

	switch relational.Operator {
	case "==":
		data.Opcodemap[op], err = data.chunkToOp(stmt.Then)
		if err != nil {
			return err
		}
		data.Opcodemap[op+1], err = data.chunkToOp(stmt.Else)
		if err != nil {
			return err
		}
	case ">":
		data.Opcodemap[op], err = data.chunkToOp(stmt.Else)
		if err != nil {
			return err
		}
		data.Opcodemap[op+1], err = data.chunkToOp(stmt.Then)
		if err != nil {
			return err
		}
	case "<=":
		then, ok := stmt.Then[0].(*ast.IfStmt)
		if ok && len(stmt.Then) == 1{
			err = data.traverseExecLoop(then) 
			if err != nil {
				return err
			}
		}
		els, ok := stmt.Else[0].(*ast.IfStmt)
		if ok && len(stmt.Else) == 1 {
			return data.traverseExecLoop(els)
		}
		return errors.New("malformed if statement in execution loop")
	default:
		return errors.New("malformed if statement, invalid relational operator")
	}
	return nil
}

// GenerateOpcodemap solves the vm exec loop and generates a lookup table of  vm opcodes to instruction creation functions
func GenerateOpcodemap(stmt *ast.IfStmt, variables []string, hashmap map[string]func(*opcodemap.Instruction)uint32) (map[int][]*opcodemap.Instruction, error) {
	data := mapData{
		Variables: variables,
		Hashmap: hashmap,
		Opcodemap: make(map[int][]*opcodemap.Instruction),
	}
	err := data.traverseExecLoop(stmt)
	if err != nil {
		return nil, err
	}
	return data.Opcodemap, nil
}

func initMapping() (map[string]func(*opcodemap.Instruction)uint32) {
	delim, err := parse.Parse(strings.NewReader(opcodemap.Delimiter), "")
	if err != nil {
		panic(err)
	}
	// We need to detect some variable names or else some opcodes have the same signature.
	variables := []string{
		"Stk", 
		"Inst", 
		"Env", 
		"Upvalues", 
		"InstrPoint",
	}
	
	replace := map[string]byte{
		"OP_A": helper.NumberExpr, 
		"OP_B": helper.NumberExpr, 
		"OP_C": helper.NumberExpr, 
		"OP_ENUM": helper.NumberExpr, 
		"OP_MOVE": helper.NumberExpr,
	}

	delimiter = helper.GenerateSignature(delim, variables)

	hashmap := make(map[string]func(*opcodemap.Instruction)uint32)

	for str, function := range opcodemap.OpCodes {
		chunk, err := parse.Parse(strings.NewReader(str), "")
		if err != nil {
			panic("parsing failed when generating the hashmap!\nStr:\n"+str)
		}

		hash := helper.GenerateSignatureWithReplace(chunk, variables, replace)
		// Making sure there aren't duplicate hashes. 
		if _, ok := hashmap[hash]; ok {
			panic("same Hash\n" + str)
		}

		hashmap[hash] = function
	}
	return hashmap
}