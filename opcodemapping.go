package deobfuscator

import (
	"fmt"
	"strconv"
	"strings"
	"./opcodemap"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

const (
	stringExpr       = "a"
	numberExpr       = "b"
	lValueExpr       = "c"
	nilExpr          = "d"
	falseExpr        = "e"
	trueExpr         = "f"
	identExpr        = "g"
	stack            = "h"
	instruction      = "i"
	comma3Expr       = "j"
	pointer          = "k"
	arithmeticOpExpr = "l"
	logicalOpExpr    = "m"
	functionExpr     = "n"
	local            = "o"
	ifStmt           = "p"
	elseStmt         = "q"
	returnStmt       = "r"
	numberForStmt    = "s"
	funcDefStmt      = "t"
	breakStmt        = "u"
	genericForStmt   = "v"
	environment      = "w"
	upvalues         = "x"
)

type constLValueExpr struct {
	ast.ExprBase

	Value lua.LValue
}

type mapData struct {
	UseNumbers  bool

	Stack       string
	Instruction string
	Environment string
	Upvalues 	string 
	Pointer     string

	Opcodemap   map[int][]string
	Hashmap     map[string]string
}

func (data *mapData) compileArithmeticOpExpr(expr *ast.ArithmeticOpExpr) string {
	ret := data.compileExpr(expr.Lhs)
	ret += expr.Operator
	ret += data.compileExpr(expr.Rhs)
	return ret
}

func (data *mapData) compileRelationalOpExpr(expr *ast.RelationalOpExpr) string {
	ret := data.compileExpr(expr.Lhs)
	ret += expr.Operator
	ret += data.compileExpr(expr.Rhs)
	return ret
}

func compileUnaryOpExpr(expr ast.Expr) string {
	switch expr.(type) {
	case *ast.UnaryMinusOpExpr:
		return "-"
	case *ast.UnaryNotOpExpr:
		return "~"
	case *ast.UnaryLenOpExpr:
		return "#"
	}
	return "PANIC"
}

func (data *mapData) compileTableExpr(expr *ast.TableExpr) string {
	ret := "{"
	for _, field := range expr.Fields {
		ret += data.compileExpr(field.Value)
	}
	return ret
}

func (data *mapData) compileExpr(expr ast.Expr) string {
	switch ex := expr.(type) {
	case *ast.StringExpr:
		return stringExpr
	case *ast.NumberExpr:
		return numberExpr
	case *constLValueExpr:
		return lValueExpr
	case *ast.NilExpr:
		return nilExpr
	case *ast.FalseExpr:
		return falseExpr
	case *ast.TrueExpr:
		return trueExpr
	case *ast.IdentExpr:
		v := ex.Value
		if data.UseNumbers && (v == "OP_A" || v == "OP_B" || v == "OP_C") {
			return numberExpr
		} else if v == "OP_ENUM" || v == "OP_MOVE" {
			return numberExpr
		} else if v == data.Stack {
			return stack
		} else if v == data.Instruction {
			return instruction
		} else if v == data.Environment {
			return environment
		} else if v == data.Upvalues {
			return upvalues
		} else if v == data.Pointer {
			return pointer
		}
		return identExpr
	case *ast.Comma3Expr:
		return comma3Expr
	case *ast.AttrGetExpr:
		return data.compileExpr(ex.Object) + "[" + data.compileExpr(ex.Key)
	case *ast.TableExpr:
		return data.compileTableExpr(ex)
	case *ast.ArithmeticOpExpr:
		return data.compileArithmeticOpExpr(ex)
	case *ast.StringConcatOpExpr:
		return arithmeticOpExpr
	case *ast.UnaryMinusOpExpr, *ast.UnaryNotOpExpr, *ast.UnaryLenOpExpr:
		return compileUnaryOpExpr(ex)
	case *ast.RelationalOpExpr:
		return data.compileRelationalOpExpr(ex)
	case *ast.LogicalOpExpr:
		return logicalOpExpr
	case *ast.FuncCallExpr:
		return data.compileFuncCallExpr(ex)
	case *ast.FunctionExpr:
		return functionExpr
	}
	return ""
}

func (data *mapData) compileAssignStmtLeft(stmt *ast.AssignStmt) string {
	ret := ""
	for _, lhs := range stmt.Lhs {
		switch ex := lhs.(type) {
		case *ast.IdentExpr:
			ret += data.compileExpr(ex) + "="
		case *ast.AttrGetExpr:
			ret += data.compileExpr(ex.Object) + "[" + data.compileExpr(ex.Key)
		}
	}
	return ret
}

func (data *mapData) compileAssignStmtRight(stmt *ast.AssignStmt) string {
	ret := ""
	for _, rhs := range stmt.Rhs {
		ret += data.compileExpr(rhs)
	}
	return ret
}

func (data *mapData) compileAssignStmt(stmt *ast.AssignStmt) string {
	return data.compileAssignStmtLeft(stmt) + data.compileAssignStmtRight(stmt)
}
func (data *mapData) compileLocalAssignStmt(stmt *ast.LocalAssignStmt) string {
	ret := identExpr
	for i := 1; i < len(stmt.Names); i++ {
		ret += identExpr
	}
	ret += "="
	for _, v := range stmt.Exprs {
		ret += data.compileExpr(v)
	}
	return ret
}


func (data *mapData) compileFuncCallExpr(expr *ast.FuncCallExpr) string {
	ret := data.compileExpr(expr.Func) + "("
	for _, ar := range expr.Args {
		ret += data.compileExpr(ar)
	}
	return ret
}

func (data *mapData) compileBranchCondition(expr ast.Expr) string {
	ret := ""
	switch ex := expr.(type) {
	case *ast.FalseExpr, *ast.NilExpr:
		ret += falseExpr
	case *ast.TrueExpr, *ast.NumberExpr, *ast.StringExpr:
		ret += trueExpr
	case *ast.UnaryNotOpExpr:
		return "~" + data.compileBranchCondition(ex.Expr)
	case *ast.LogicalOpExpr:
		return data.compileBranchCondition(ex.Lhs) + ex.Operator + data.compileBranchCondition(ex.Rhs)
	case *ast.RelationalOpExpr:
		ret := data.compileExpr(ex.Lhs)
		ret += ex.Operator
		ret += data.compileExpr(ex.Rhs)
	}
	ret += data.compileExpr(expr)
	return ret
}

func (data *mapData) compileIfStmt(stmt *ast.IfStmt) string {
	ret := ifStmt
	ret += data.compileBranchCondition(stmt.Condition)
	for _, a := range stmt.Then {
		ret += data.compileStmt(a)
	}

	if len(stmt.Else) > 0 {
		ret += elseStmt
		for _, a := range stmt.Else {
			ret += data.compileStmt(a)
		}
	}
	return ret
}

func (data *mapData) compileReturnStmt(stmt *ast.ReturnStmt) string {
	ret := returnStmt
	for _, expr := range stmt.Exprs {
		ret += data.compileExpr(expr)
	}
	return ret
}

func (data *mapData) compileNumberForStmt(stmt *ast.NumberForStmt) string {
	ret := numberForStmt
	ret += data.compileExpr(stmt.Init)
	ret += data.compileExpr(stmt.Limit)
	ret += data.compileExpr(stmt.Step)
	for _, a := range stmt.Stmts {
		ret += data.compileStmt(a)
	}
	return ret
}

func (data *mapData) compileStmt(stmt ast.Stmt) string {
	switch st := stmt.(type) {
	case *ast.AssignStmt:
		return data.compileAssignStmt(st)
	case *ast.LocalAssignStmt:
		return data.compileLocalAssignStmt(st)
	case *ast.FuncCallStmt:
		return data.compileFuncCallExpr(st.Expr.(*ast.FuncCallExpr))
	case *ast.DoBlockStmt:
		ret := ""
		for _, a := range st.Stmts {
			ret += data.compileStmt(a)
		}
		return ret
	case *ast.WhileStmt:
		return "PANIC"
	case *ast.RepeatStmt:
		return "PANIC"
	case *ast.FuncDefStmt:
		return "FuncDefStmt"
	case *ast.ReturnStmt:
		return data.compileReturnStmt(st)
	case *ast.IfStmt:
		return data.compileIfStmt(st)
	case *ast.BreakStmt:
		return "BreakStmt"
	case *ast.NumberForStmt:
		return data.compileNumberForStmt(st)
	case *ast.GenericForStmt:
		return "GenericForStmt"
	}
	return ""
}

func (data *mapData) solveSuperOp(chunk []ast.Stmt) []string {
	localassing := identExpr+"="
	ops := []string{"SuperOperator"}
	hash := ""
	for _,exp := range chunk {
		temp := data.compileStmt(exp)
		if data.Hashmap[temp] == "Delimiter" {
			ops = append(ops,data.Hashmap[hash])
			hash = ""
			continue
		} else if temp == localassing || data.Hashmap[temp] == "Delimiter2" {
			continue
		}
		hash += temp
	}
	return append(ops,data.Hashmap[hash])
}

func (data *mapData) chunkToOp(chunk []ast.Stmt) []string {
	var hash string
	for _, l := range chunk {
		hash += data.compileStmt(l)
	}
	fmt.Println(hash)
	ret := data.Hashmap[hash]
	if ret == "" {
		return data.solveSuperOp(chunk)
	} 
	return []string{ret}
}


func (data *mapData) solveIf(chunk []ast.Stmt) {
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

// GenerateHashmap allows you to generate the hashmap once and then reuse it.
func GenerateHashmap() map[string]string {
	data := &mapData{
		UseNumbers:  true,
		Stack:       "Stk",
		Instruction: "Inst",
		Environment: "Env",
		Upvalues: "Upvalues",
		Pointer:  "InstrPoint",
	}
	opMap := make(map[string]string)
	for _, i := range opcodemap.Map() {
		for str, op := range i {
			chunk, err := parse.Parse(strings.NewReader(str), "")
			if err != nil {
				panic(err)
			}
			hash := ""
			for _, stmt := range chunk {
				hash += data.compileStmt(stmt)
			}
			if opMap[hash] != "" {
				panic("Same Hash")//, op, opMap[hash])
			}
			opMap[hash] = op
		}
	}
	return opMap
}

func generateOpcodemap(vm *vmdata, hashmap map[string]string) map[int][]string {
	data := &mapData{
		UseNumbers:  false,

		Stack:       vm.Stack,
		Instruction: vm.Instruction,
		Environment: vm.Env,
		Upvalues:    vm.Upvalues,
		Pointer:     vm.PC,

		Hashmap:     hashmap,
		Opcodemap: make(map[int][]string),
	}
	data.solveIf(vm.VMLoop.Stmts)
	return data.Opcodemap
}
