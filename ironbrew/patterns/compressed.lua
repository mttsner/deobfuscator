local function decompress(b)
    local c,d,e="","",{}
    local f=256;
    local g={}
    for h=0,f-1 do 
        g[h]=Char(h)
    end;
    local i=1;
    local function k()
        local l=ToNumber(Sub(b, i,i),36)
        i=i+1;
        local m=ToNumber(Sub(b, i,i+l-1),36)
        i=i+l;
        return m 
    end;
    c=Char(k())
    e[1]=c;
    while i<#b do 
        local n=k()
        if g[n]then 
            d=g[n]
        else 
            d=c..Sub(c, 1,1)
        end;
        g[f]=c..Sub(d, 1,1)
        e[#e+1],c,f=d,d,f+1 
    end;
    return table.concat(e)
end;
local ByteString=decompress(_StringExpr_);