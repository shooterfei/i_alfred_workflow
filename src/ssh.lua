#!/usr/bin/env lua

local alfy = require("core.alfy")
local cjson = require("cjson")

-- print(os.getenv("data"))

-- local data = nil
-- if arg[1] == nil then
--   arg[1] = ""
-- end
-- for key, value in ipairs(arg) do
--   -- print(key, value)
-- 	if value == "-d" then
-- 		data = arg[key + 1]
-- 	end
-- end

local mode = os.getenv("mode")

local json_file = "./data/ssh.json"
if mode == "temp" then
  json_file = "./data/temp.json"
end

local json_data = alfy.json_file_load(json_file)

local items = {}
for k, v in pairs(json_data) do
	local temp = {}
	if k ~= "default" then
    v["group"] = k
		if arg[1] == "" then
			temp = {
				title = k,
				subtitle = v["label"],
				arg = cjson.encode(v),
			}
			table.insert(items, temp)
		else
			for key, value in ipairs(json_data[k]["items"]) do
        if string.find(value["host"], arg[1]) or string.find(value["description"], arg[1]) or string.find(k, arg[1]) then
					temp = {
						title = k .. "/" .. value["host"],
						subtitle = value["description"],
						arg = cjson.encode(value),
					}
					table.insert(items, temp)
				end
			end
		end
	else
		for key, value in ipairs(json_data[k]["items"]) do
			if arg[1] == nil or string.find(value["host"], arg[1]) ~= nil then
				temp = {
					title = value["host"],
					subtitle = value["description"],
          arg = cjson.encode(value),
				}
				table.insert(items, temp)
			end
		end
	end
end

alfy.output(items)
