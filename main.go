package main

import (
	"os"
	"path/filepath"

	"github.com/notnoobmaster/deobfuscator/ironbrew"
	"github.com/notnoobmaster/deobfuscator/helper"
	"github.com/notnoobmaster/luautil/parse"
)

func init() {
	ironbrew.Initialize()
}

func main() {
	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		panic(err)
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	chunk, err := parse.Parse(file, "")
	if err != nil {
		panic(err)
	}

	proto, err := ironbrew.Deobfuscate(chunk)
	if err != nil {
		panic(err)
	}

	luac := helper.DumpLua(proto)
	if os.WriteFile("out.luac", luac, 0644) != nil {
		panic(err)
	}
}