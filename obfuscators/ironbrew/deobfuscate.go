package ironbrew

import (
	"fmt"

	"github.com/notnoobmaster/deobfuscator/obfuscators/ironbrew/opcodemap"
	lua "github.com/yuin/gopher-lua"
	"github.com/yuin/gopher-lua/ast"
)

//const debug = false
var hashmap map[string]func(*opcodemap.Instruction)uint32

// Deobfuscate ironbrew
func Deobfuscate(chunk []ast.Stmt)(*lua.FunctionProto, error)  {
	data := vmdata{}
	err := data.getVmdata(chunk)
	if err != nil {
		return nil, err
	}
	variables := []string{data.Stack, data.Inst, data.Env, data.Upvalues, data.InstPtr}
	data.Opcodemap, err = GenerateOpcodemap(data.Loop, variables, hashmap)
	if err != nil {
		return nil, err
	}
	return data.deserialize()
}

// Initialize everything needed for deobfuscation
func Initialize() error {
	err := initVmdata()
	if err != nil {
		return err
	}
	hashmap, err = initMapping()
	fmt.Print(len(hashmap))
	if err != nil {
		return err
	}
	return nil
}

/*
	if debug {
		fmt.Println("Obfuscator: Ironbrew")
		fmt.Printf("Obfuscation settings: %+v\n", data.Settings)
		fmt.Println("Key:",data.Key)
		fmt.Println("Bool:",data.Bool)
		fmt.Println("Float:",data.Float)
		fmt.Println("String:",data.String)
		fmt.Println("Order:", data.Order)
		fmt.Println("Upvalues:", data.Upvalues)
		fmt.Println("Environment:", data.Env)
		//fmt.Println("Stack:", data.Stack)
		//fmt.Println("Instruction:", data.Instruction)
		//fmt.Println("PC:", data.PC)
		//fmt.Println("data loop found:", data.dataLoop != nil)
		//fmt.Println("Bytecode:", data.Bytecode)
	}
*/