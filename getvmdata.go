package deobfuscator

import (
	"github.com/yuin/gopher-lua/ast"
	"../beautifier"
)


[`local function gBits32()
local W, X, Y, Z = Byte(ByteString, Pos, Pos + 3);

W = BitXOR(W, XOR_KEY)
X = BitXOR(X, XOR_KEY)
Y = BitXOR(Y, XOR_KEY)
Z = BitXOR(Z, XOR_KEY)

Pos	= Pos + 4;
return (Z*16777216) + (Y*65536) + (X*256) + W;
end;`] = gBits32

func (vm *vmdata) gBits32() {

}


func generatePatterns()  {
	map[string]func()
}

func getvmdata(chunk []ast.Stmt) (data?, error){
	//beautifier.Match(patterns)
	return data, nil
}