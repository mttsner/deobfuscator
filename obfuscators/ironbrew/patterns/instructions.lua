for Idx=1,gBits32() do 
		local Descriptor = gBits8();
		if (gBit(Descriptor, 1, 1) == 0) then
			local Type = gBit(Descriptor, 2, 3);
			local Mask = gBit(Descriptor, 4, 6);
			
			local Inst=
			{
				gBits16(),
				gBits16(),
				nil,
				nil
			};
                                    
			if (Type == 0) then 
				Inst[3] = gBits16(); 
				Inst[4] = gBits16();
			elseif(Type==1) then 
					Inst[3] = gBits32();
			elseif(Type==2) then 
				Inst[3] = gBits32() - (2 ^ 16)
			elseif(Type==3) then 
				Inst[3] = gBits32() - (2 ^ 16)
				Inst[4] = gBits16();
			end;
	
			if (gBit(Mask, 1, 1) == 1) then Inst[2] = Consts[Inst[2]] end
			if (gBit(Mask, 2, 2) == 1) then Inst[3] = Consts[Inst[3]] end
			if (gBit(Mask, 3, 3) == 1) then Inst[4] = Consts[Inst[4]] end
										
			Instrs[Idx] = Inst;
		end
    end;