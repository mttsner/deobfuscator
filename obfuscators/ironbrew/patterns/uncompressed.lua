--[[local Byte         = string.byte;
local Char         = string.char;
local Sub          = string.sub;
local Concat       = table.concat;
local Insert       = table.insert;
local LDExp        = math.ldexp;
local GetFEnv      = getfenv or function() return _ENV end;
local Setmetatable = setmetatable;
]]
local Select       = select;
local Unpack = unpack or table.unpack;
local ToNumber = tonumber;
local ByteString = _StringExpr_