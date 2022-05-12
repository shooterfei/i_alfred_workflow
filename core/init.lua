JSON = (loadfile "core/json.lua")



-- JSON = require("JSON")
JSON.decodeNumbersAsObjects = true
local lua_value = JSON:decode('{"num":10201348204923842304}')
local raw_json_text = JSON:encode(lua_value)
print(raw_json_text)
