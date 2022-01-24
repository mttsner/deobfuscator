local function gBit(Bit, Start, End)
end;

local Pos = 1;

local function gBits32()
    local W, X, Y, Z = string.byte(ByteString, Pos, Pos + 3);

	W = BitXOR(W, _NumberExpr_)
	X = BitXOR(X, 8)
	Y = BitXOR(Y, 8)
	Z = BitXOR(Z, 8)

    Pos	= Pos + 4;
    return (Z*16777216) + (Y*65536) + (X*256) + W;
end;

local function gBits8() end;

local function gBits16() end;

local function gFloat() end;

local gSizet = gBits32;
local function gString(Len)

end;

local gInt = gBits32;
local function _R(...) end

local function _LocalFunctionStmt_() end

local PCall = pcall
local function Wrap(Chunk, _IdentExpr_, _IdentExpr_)
	local Instr = Chunk[1];
	local Proto = Chunk[2];
	local Params = Chunk[3];
	return function(...)
		local _IdentExpr_ = 1;
		local Top = -1;

		local Args = {...};
		local PCount = Select('#', ...) - 1;
		local Instr  = Chunk[1];
		local Proto  = Chunk[2];
		local Params = Chunk[3];

		local function Loop()
			local _R = _R
			local Vararg = {};
			local Lupvals	= {};
			local _IdentExpr_ = {};
	
			for Idx = 0, PCount do end;
	
			local Varargsz = PCount - Params + 1

			local _IdentExpr_;
			local Enum;	

			while true do
				Inst		= Instr[InstrPoint];
				Enum		= Inst[1];
				_IfStmt_()
				InstrPoint	= InstrPoint + 1;
			end;
		end;

		A, B = _R(PCall(Loop))
		if not A[1] then else end;
	end;
end;	