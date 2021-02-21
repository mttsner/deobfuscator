package deobfuscator

import (
	"errors"
	"io"

	"github.com/notnoobmaster/beautifier"
	"github.com/notnoobmaster/deobfuscator/obfuscators/ironbrew"
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/ast"
	"github.com/yuin/gopher-lua/parse"
)

// Obfuscators that are supported
var Obfuscators = map[string]func([]ast.Stmt)(*lua.FunctionProto, error) {
	"Ironbrew": ironbrew.Deobfuscate,
}

var initialized = false
// Initialize the the pre-deobfuscation tasks.
func Initialize() {
	ironbrew.Initialize()
}

// Deobfuscate virtualized lua code.
func Deobfuscate(file io.Reader, debug bool) (*lua.FunctionProto, error) {
	if !initialized {
		Initialize()
	}

	chunk, err := parse.Parse(file, "")
	if err != nil {
		return nil, err
	}
	
	/*ast =*/ beautifier.Optimize(chunk)

	for _, deobfuscate := range Obfuscators {
		if proto, err := deobfuscate(chunk); err == nil {
			return proto, err
		}
	}

	return nil, errors.New("Couldn't deobfuscate")
}