package ironbrew

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/notnoobmaster/luautil/parse"
	"github.com/notnoobmaster/deobfuscator/helper"
)

//go:embed test.lua
var test string

func TestGetVmdata(t *testing.T) {
	chunk, err := parse.Parse(strings.NewReader(test), "")
	if err != nil {
		t.Fatal(err)
	}
	initVmdata()
	data := vmdata{}
	err = data.getVmdata(chunk)
	t.Logf("%#v", data)
	if err != nil {
		t.Error(err)
	}
}

func TestDeobfuscate(t *testing.T) {
	chunk, err := parse.Parse(strings.NewReader(test), "")
	if err != nil {
		t.Fatal(err)
	}
	//beautifier.Optimize(chunk)
	Initialize()
	out, err := Deobfuscate(chunk)
	out.DbgSourcePositions = make([]int, len(out.Code))
	t.Log(out.String())
	if err != nil {
		t.Error(out)
	}
}

const vm = `
A = Inst[2]
Stk[A](Unpack(Stk, A + 1, Inst[3]))
`

const str = `
local A = Inst[OP_A]
Stk[A](Unpack(Stk, A + 1, Inst[OP_B]))
`

func TestHash(t *testing.T) {
	chunk1, err := parse.Parse(strings.NewReader(str), "")
	if err != nil {
		t.Fatal(err)
	}
	chunk2, err := parse.Parse(strings.NewReader(vm), "")
	if err != nil {
		t.Fatal(err)
	}
	variables := []string{"Stk", "Inst", "Env", "Upvalues", "InstrPoint",}
	replace := map[string]byte{
		"OP_A": helper.NumberExpr, 
		"OP_B": helper.NumberExpr, 
		"OP_C": helper.NumberExpr, 
		"OP_ENUM": helper.NumberExpr, 
		"OP_MOVE": helper.NumberExpr,
	}
	hash1 := helper.GenerateSignatureWithReplace(chunk1, variables, replace)
	hash2 := helper.GenerateSignature(chunk2, variables)
	if hash1 != hash2 {
		t.Error(hash1, hash2)
	}
}

/*
func TestMatch(t *testing.T) {
	target, _ := parse.Parse(strings.NewReader(strTarget), "")
	pattern, err := parse.Parse(strings.NewReader(strTarget), "")

	success, _, := beautifier.Match(target, pattern)

	if !success
}*/