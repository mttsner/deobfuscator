local function gBits32()
	local W, X, Y, Z = Byte(ByteString, Pos, Pos + 3);
	
	W = BitXOR(W, _NumberExpr_)
	X = BitXOR(X, XOR_KEY)
	Y = BitXOR(Y, XOR_KEY)
	Z = BitXOR(Z, XOR_KEY)
	
	Pos	= Pos + 4;
	return (Z*16777216) + (Y*65536) + (X*256) + W;
end;