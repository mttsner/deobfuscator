package patterns

import "github.com/yuin/gopher-lua/ast"

const strBits32 = 
`
local function gBits32()
	local W, X, Y, Z = Byte(ByteString, Pos, Pos + 3);
	
	W = BitXOR(W, XOR_KEY)
	X = BitXOR(X, XOR_KEY)
	Y = BitXOR(Y, XOR_KEY)
	Z = BitXOR(Z, XOR_KEY)
	
	Pos	= Pos + 4;
	return (Z*16777216) + (Y*65536) + (X*256) + W;
end;
`

func getBits32Data(bits32 []ast.Stmt) {
	key := bits32[1].(ast.AssignStmt).Rhs.(ast.FuncCallExpr).Args[1].(ast.NumberExpr).Value
}

const strConstLoop = 
`
for Idx=1, ConstCount do 
	local Type =gBits8();
	local Cons;

	if(Type==CONST_BOOL) then Cons = (gBits8() ~= 0);
	elseif(Type==CONST_FLOAT) then Cons = gFloat();
	elseif(Type==CONST_STRING) then Cons = gString();
	end;
	
	Consts[Idx] = Cons;
end;
`

func getConstLoopData(loop []ast.Stmt) {
	ifstmt  := loop[2].(ast.IfStmt)
	elseif  := ifstmt.Else[0].(ast.IfStmt) 
	elseif2 := elseif.Else[0].(ast.IfStmt) 

	BOOL   := ifstmt.Condition.(ast.RelationalOpExpr).Rhs.(ast.NumberExpr).Value
	FLOAT  := elseif.Condition.(ast.RelationalOpExpr).Rhs.(ast.NumberExpr).Value
	STRING := elseif2.Condition.(ast.RelationalOpExpr).Rhs.(ast.NumberExpr).Value
}

const strWrap =
`
function(Chunk, Upvalues, Env)
	local Instr = Chunk[1];
	local Proto = Chunk[2];
	local Params = Chunk[3];
	return function(...) end;
end;
`

func getWrapData(wrap []ast.Stmt) {
	function := 
	upvalues := wrap.() 
	env      := 
}