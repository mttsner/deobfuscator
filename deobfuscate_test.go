package deobfuscator

import (
	"os"
	"testing"

	"github.com/notnoobmaster/beautifier"
	"github.com/yuin/gopher-lua/parse"
)

func TestDeobfuscate(t *testing.T) {
	file, err := os.Open("obfuscated.lua")
	if err != nil {
		t.Fatal(err)
	}
	proto, err := Deobfuscate(file, false)
	t.Log(proto)
	if err != nil {
		t.Error(err)
	}
}

func TestOptimize(t *testing.T) {
	file, err := os.Open("obfuscated.lua")
	if err != nil {
		t.Fatal(err)
	}
	chunk, err := parse.Parse(file, "")
	if err != nil {
		t.Fatal(err)
	}
	
	beautifier.Optimize(chunk)
	
	newfile, err := os.Create("beautified.lua")

	newfile.WriteString(beautifier.Beautify(&chunk))
}