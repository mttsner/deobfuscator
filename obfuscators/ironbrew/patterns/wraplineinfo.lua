local function Wrap(Chunk, _IdentExpr_, _IdentExpr_) -- Upvalues and Env
	local Instr = Chunk[1];
	local Proto = Chunk[2];
	local Params = Chunk[3];

	return function(...)
		local InstrPoint = 1;
		local Top = -1;

		local Args = {...};
		local PCount = Select('#', ...) - 1;

		local function Loop()
			local Instr  = Instr; 
			local Const  = Const; 
			local Proto  = Proto; 
			local Params = Params;

			local _R = _R
			local Vararg = {};

			local Lupvals	= {};
			local Stk		= {};
	
			for Idx = 0, PCount do
				if (Idx >= Params) then
					Vararg[Idx - Params] = Args[Idx + 1];
				else
					Stk[Idx] = Args[Idx + 1];
				end;
			end;
	
			local Varargsz = PCount - Params + 1

			local Inst;
			local Enum;	

			while true do
				-- This is ignored.
			end;
		end;

		A, B = _R(PCall(Loop))
		if not A[1] then
			local line = Chunk[7][InstrPoint] or '?'
			error('ERROR IN IRONBREW SCRIPT [LINE ' .. line .. ']:' .. A[2])
		else
			return Unpack(A, 2, B)
		end;
	end;
end;	