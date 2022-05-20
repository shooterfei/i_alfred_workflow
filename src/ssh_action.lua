#!/usr/bin/env lua

local alfy = require("core.alfy")
local cjson = require("cjson")



local args = table.concat(arg)
alfy.write_json_file("./data/temp.json", args)

local ssh_config = cjson.decode(args)

print(ssh_config["label"])

if ssh_config["label"] == "group" then
  os.execute("node ./dist/index.js temp")
else
  local shell_script = [[
  export host=%s
  export port=%s
  export username=%s
  export password=%s
  export rootdir=%s
  export timeout=%s
  node ./dist/index.js run
]]

  os.execute(string.format(shell_script,
    ssh_config["host"],
    ssh_config["port"],
    ssh_config["username"],
    ssh_config["password"],
    ssh_config["rootdir"],
    ssh_config["timeout"]
  ))
end

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
