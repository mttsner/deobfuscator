package ironbrew

import (
	// assigned to _ because go:embed doesn't work without requiring embed
	_ "embed"
	"errors"
	"strconv"
	"strings"

	"github.com/notnoobmaster/beautifier"
	//"github.com/notnoobmaster/deobfuscator"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

type settings struct {
	BytecodeCompress bool
	PreserveLineInfo bool
}

type vmdata struct {
	Loop     *ast.IfStmt
	Settings settings
	Deserialize *ast.FunctionExpr
	Key      byte
	Bool     int
	Float    int
	String   int
	Env      string
	Upvalues string
	Bytecode []byte
}

//go:embed "patterns/constloop.lua"
var strConstLoop string
var astConstLoop []ast.Stmt

func (data *vmdata) constants(chunk []ast.Stmt) bool {
	success, exprs, _ := beautifier.Match(chunk, astConstLoop)
	if !success {
		return false
	}
	data.Bool, _ = strconv.Atoi(exprs[0].(*ast.NumberExpr).Value)
	data.Float, _ = strconv.Atoi(exprs[1].(*ast.NumberExpr).Value)
	data.String, _ = strconv.Atoi(exprs[2].(*ast.NumberExpr).Value)
	return true
}

//go:embed patterns/compressed.lua
var strCompressed string
var astCompressed []ast.Stmt

func (data *vmdata) compressed(chunk []ast.Stmt) bool {
	success, exprs, _ := beautifier.Match(chunk, astCompressed)
	if !success {
		return success
	}
	byteString := exprs[0].(*ast.StringExpr).Value
	if bytecode, err := decompress(byteString); err == nil {
		data.Settings.BytecodeCompress = true
		data.Bytecode = bytecode
		return success
	}
	return false
}

//go:embed patterns/uncompressed.lua
var strUncompressed string
var astUncompressed []ast.Stmt

func (data *vmdata) uncompressed(chunk []ast.Stmt) bool {
	success, exprs, _ := beautifier.Match(chunk, astUncompressed)
	if !success {
		return success
	}
	data.Bytecode = []byte(exprs[0].(*ast.StringExpr).Value)
	return success
}

//go:embed patterns/normal.lua
var strNormal string
var astNormal []ast.Stmt

func (data *vmdata) normal(chunk []ast.Stmt) bool {
	success, exprs, stmts := beautifier.Match(chunk, astNormal)
	if !success {
		return success
	}
	key, _ := strconv.Atoi(exprs[0].(*ast.NumberExpr).Value)
	data.Key = byte(key)
	data.Deserialize = exprs[1].(*ast.FunctionExpr)
	data.Upvalues = exprs[2].(*ast.IdentExpr).Value
	data.Env = exprs[3].(*ast.IdentExpr).Value
	data.Loop = stmts[0].(*ast.IfStmt)
	return success
}

//go:embed patterns/lineinfo.lua
var strLineinfo string
var astLineinfo []ast.Stmt

func (data *vmdata) lineinfo(chunk []ast.Stmt) bool {

	return true
}

func (data *vmdata) GetVmdata(chunk []ast.Stmt) (err error) {
	if !(data.compressed(chunk) || data.uncompressed(chunk)) {
		return errors.New("Couldn't get VM bytecode")
	}

	if !(data.normal(chunk) || data.lineinfo(chunk)) {
		return errors.New("Couldn't get VM data")
	}

	if !data.constants(data.Deserialize.Stmts) {
		return errors.New("Couldn't get constant keys")
	}

	return nil
}

func compile(str string) []ast.Stmt {
	chunk, err := parse.Parse(strings.NewReader(str), "")
	if err != nil {
		panic(err) //panic("Ironbrew: pattern parsing failed")
	}
	return chunk
}

func initVmdata() {
	astConstLoop = compile(strConstLoop)

	astCompressed = compile(strCompressed)
	astUncompressed = compile(strUncompressed)

	astNormal = compile(strNormal)
	astLineinfo = compile(strLineinfo)
}
