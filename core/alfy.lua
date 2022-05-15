local cjson = require("cjson")
local http = require("socket.http")
local ltn12 = require("ltn12")
local io = require("io")

local M = {}

local fetch = function(url, body)
	local response_body = {}
	body["url"] = url
	body["sink"] = ltn12.sink.table(response_body)
	local res = { http.request(body) }
	if res[2] == 200 then
		if type(response_body) == "table" then
			return cjson.decode(table.concat(response_body))
		end
	else
		return {
			["code"] = false,
		}
	end
end

local output = function(data)
	local items = {
		["items"] = data,
	}
	print(cjson.encode(items))
end

local json_file_load = function(filePath)
	local file = io.open(filePath, "r")
  if file ~= nil then
    local json = file:read("a")
    file:close()
    return cjson.decode(json)
  end
  return {}
end

M.fetch = fetch
M.output = output
M.json_file_load = json_file_load
return M
