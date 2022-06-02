local cjson = require("cjson")
local http = require("socket.http")
local ltn12 = require("ltn12")
local io = require("io")
local inspect = require("core.inspect")

local M = {}

M.fetch = function(url, body)
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

M.inspect = inspect.inspect

M.fetch_html = function(url, body)
	local response_body = {}
	body["url"] = url
	body["sink"] = ltn12.sink.table(response_body)
	local res = { http.request(body) }
	if res[2] == 200 then
		if type(response_body) == "table" then
			return table.concat(response_body)
		end
	else
		return {
			["code"] = false,
		}
	end
end

M.output = function(data)
	local items = {
		["items"] = data,
	}
	print(cjson.encode(items))
end

M.json_file_load = function(filePath)
	local file = io.open(filePath, "r")
  if file ~= nil then
    local json = file:read("*a")
    file:close()
    return cjson.decode(json)
  end
  return {}
end


M.write_json_file = function(filePath, json_data)
  local file = io.open(filePath, "w+")
  if file ~= nil then
    file:write(json_data)
    file:close()
    return true
  end
  print("err")
  return false
end


return M
