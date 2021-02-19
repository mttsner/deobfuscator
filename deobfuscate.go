package deobfuscator

import (
	"github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/parse"
	"github.com/yuin/gopher-lua/ast"
	"github.com/notnoobmaster/beautifier"
	"io"
)

type placeholder struct {

}

var hashmap map[string]func()uint32
var patterns map[string]string
var initialized = false

func (p *placeholder) Deserialize() (*lua.FunctionProto, error) {
	// Placeholder function
	return &lua.FunctionProto{}, nil
}

func getvmdata(_ []ast.Stmt) (*placeholder, error) {
	for pattern, callback := obfuscator.VmData {
		success, exprs := beautifier.Match(ast, pattern)
		if success {
			callback(exprs)
		}
	} 
	// Placeholder function
	return &placeholder{}, nil
}

// Initialize the the pre-deobfuscation tasks.
func Initialize() {
	hashmap = GenerateHashmap()
}
// Deobfuscate virtualized lua code.
func Deobfuscate(file io.Reader, debug bool) (*lua.FunctionProto, error) {
	if !initialized {
		Initialize()
	}

	ast, err := parse.Parse(file, "")
	if err != nil {
		return nil, err
	}
	
	/*ast =*/ beautifier.Optimize(ast)
	data, err := getvmdata(ast)
	if err != nil {
		return nil, err
	}

	proto, err := data.Deserialize() 

	return proto, err
}