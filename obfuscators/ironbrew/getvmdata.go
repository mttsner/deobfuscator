package ironbrew

import (
	// assigned to _ because go:embed doesn't work without requiring embed
	_ "embed"
	"errors"
	"strconv"
	"strings"

	"github.com/notnoobmaster/beautifier"
	//"github.com/notnoobmaster/deobfuscator"
	"github.com/yuin/gopher-lua/parse"
	"github.com/yuin/gopher-lua/ast"
)

type settings struct {
	BytecodeCompress bool
	PreserveLineInfo bool
}

type vmdata struct {
	Key byte
	Bool int 
	Float int
	String int
	Env string
	Upvalues string
	Bytecode []byte
	Settings settings
}

//go:embed "patterns/gbits32.lua"
var strBtis32 string
var astBits32 []ast.Stmt

func (data *vmdata) getBits32Data(chunk []ast.Stmt) bool {
	success, exprs := beautifier.Match(chunk, astBits32)
	if !success {
		return false
	}
	key, _ := strconv.Atoi(exprs[0].(*ast.NumberExpr).Value)
	data.Key = byte(key)
	return true
}

//go:embed "patterns/constloop.lua"
var strConstLoop string
var astConstLoop []ast.Stmt

func (data *vmdata) getConstLoopData(chunk []ast.Stmt) bool {
	success, exprs := beautifier.Match(chunk, astConstLoop)
	if !success {
		return false
	}
	data.Bool,   _ = strconv.Atoi(exprs[0].(*ast.NumberExpr).Value)
	data.Float,  _ = strconv.Atoi(exprs[1].(*ast.NumberExpr).Value)
	data.String, _ = strconv.Atoi(exprs[2].(*ast.NumberExpr).Value)
	return true
}

//go:embed "patterns/wrap.lua"
var strWrap string
var astWrap []ast.Stmt

//go:embed "patterns/wraplineinfo.lua"
var strWrapLineInfo string
var astWrapLineInfo []ast.Stmt

func (data *vmdata) getWrapData(chunk []ast.Stmt) bool {
	success, exprs := beautifier.Match(chunk, astWrap)
	if success {
		data.Upvalues = exprs[0].(*ast.IdentExpr).Value
		data.Env      = exprs[1].(*ast.IdentExpr).Value
		return true
	}
	success, exprs = beautifier.Match(chunk, astWrapLineInfo)
	if success {
		data.Upvalues = exprs[0].(*ast.IdentExpr).Value
		data.Env      = exprs[1].(*ast.IdentExpr).Value
		data.Settings.PreserveLineInfo = true
		return true
	}
	return false
}

//go:embed "patterns/compressed.lua"
var strCompressed string
var astCompressed []ast.Stmt

//go:embed "patterns/bytestring.lua"
var strNormal string
var astNormal []ast.Stmt

func (data *vmdata) getBytecode(chunk []ast.Stmt) bool {
	success, exprs := beautifier.Match(chunk, astNormal)
	if success {
		data.Bytecode = []byte(exprs[0].(*ast.StringExpr).Value)
		return true
	}
	success, exprs = beautifier.Match(chunk, astCompressed)
	if success {
		byteString := exprs[0].(*ast.StringExpr).Value
		if bytecode, err := decompress(byteString); err == nil {
			data.Settings.BytecodeCompress = true
			data.Bytecode = bytecode
			return true
		}
	}
	return false
}

func (data *vmdata) GetVmdata(chunk []ast.Stmt) (err error){
	if !data.getBits32Data(chunk) {
		return errors.New("Couldn't get the decryption key")
	}
	if !data.getConstLoopData(chunk) {
		return errors.New("Couldn't get constant keys")
	}
	if !data.getWrapData(chunk) {
		return errors.New("Couldn't get wrap function variables")
	}
	if !data.getBytecode(chunk) {
		return errors.New("Couldn't get the bytecode")
	}
	return
}

func compile(str string) []ast.Stmt {
	chunk, err := parse.Parse(strings.NewReader(str), "")
	if err != nil {
		panic(err)//panic("Ironbrew: pattern parsing failed")
	}
	return chunk
}

func initVmdata() {
	astBits32 = compile(strBtis32)
	astConstLoop = compile(strConstLoop)
	astWrap = compile(strWrap)
	astWrapLineInfo = compile(strWrapLineInfo)
	astCompressed = compile(strCompressed)
	astNormal = compile(strNormal)
}