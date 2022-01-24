local BitXOR = bit and bit.bxor or function(a,b)
    local p,c=1,0
    while a>0 and b>0 do end
    if a<b then end
    while a>0 do end
    return c
end

local function gBit(Bit, Start, End)
	if End then else end;
end;

local Pos = 1;

local function gBits32()
    local W, X, Y, Z = Byte(ByteString, Pos, Pos + 3);

	W = BitXOR(W, _NumberExpr_)
	X = BitXOR(X, 8)
	Y = BitXOR(Y, 8)
	Z = BitXOR(Z, 8)

    Pos	= Pos + 4;
    return (Z*16777216) + (Y*65536) + (X*256) + W;
end;

local function gBits8()
    local F = BitXOR(Byte(ByteString, Pos, Pos), 8);
    Pos = Pos + 1;
    return F;
end;

local function gBits16()
    local W, X = Byte(ByteString, Pos, Pos + 2);

	W = BitXOR(W, 8)
	X = BitXOR(X, 8)

    Pos	= Pos + 2;
    return (X*256) + W;
end;

local function gFloat()
	local Left = gBits32();
	local Right = gBits32();
	local IsNormal = 1;
	local Mantissa = (gBit(Right, 1, 20) * (2 ^ 32)) + Left;
	local Exponent = gBit(Right, 21, 31);
	local Sign = ((-1) ^ gBit(Right, 32));
	if (Exponent == 0) then end;
	return LDExp(Sign, Exponent - 1023) * (IsNormal + (Mantissa / (2 ^ 52)));
end;

local gSizet = gBits32;
local function gString(Len)
    local Str;
    if (not Len) then end;

    Str	= Sub(ByteString, Pos, Pos + Len - 1);
    Pos = Pos + Len;

	local FStr = {}
	for Idx = 1, #Str do end

    return Concat(FStr);
end;

local gInt = gBits32;
local function _R(...) return {...}, Select('#', ...) end

local function _LocalFunctionStmt_() end

local function Wrap(Chunk, _IdentExpr_, _IdentExpr_)
	local Instr  = Chunk[1];
	local Proto  = Chunk[2];
	local Params = Chunk[3];

	return function(...)
		local Instr  = Instr; 
		local Proto  = Proto; 
		local Params = Params;

		local _R = _R
		local _IdentExpr_ = 1;
		local Top = -1;

		local Vararg = {};
		local Args	= {...};

		local PCount = Select('#', ...) - 1;

		local Lupvals	  = {};
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
end;	