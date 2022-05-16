#!/usr/bin/env lua

-- local cjson = require("cjson")
local alfy = require "core.alfy"

alfy.write_json_file("./data/temp.json", table.concat(arg))
os.execute("node ./dist/index.js temp")

-- -- print(os.getenv("data"))
-- -- print("212324244")
-- -- print("args----> " .. arg[1])

-- local json_data = cjson.decode(os.getenv("data"))

-- local items = {}
-- if json_data["label"] == "group" then
-- 	for k, v in ipairs(json_data["items"]) do
-- 		local temp = {
-- 			title = json_data["group"] .. "/" .. v["host"],
-- 			subtitle = v["description"],
-- 			arg = cjson.encode(v),
-- 		}
-- 		table.insert(items, temp)
-- 	end
--   alfy.output(items)
-- end
