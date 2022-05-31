local alfy = require("core.alfy")
local cjson = require("cjson")

local mode = os.getenv("mode")
local pwd = os.getenv("PWD")

local json_file = "../data/ssh.json"
if mode == "temp" then
  json_file = "../data/temp.json"
end

local json_data = alfy.json_file_load(json_file)

local items = {}
-- luacheck: ignore 311
local temp = {}
if mode ~= "temp" then
  for k, v in pairs(json_data) do
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
        for _, value in ipairs(json_data[k]["items"]) do
          if string.find(value["host"], arg[1])
              or string.find(value["description"], arg[1])
              or string.find(k, arg[1])
          then
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
      for _, value in ipairs(json_data[k]["items"]) do
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
else
  if json_data["label"] == "group" then
    for _, v in ipairs(json_data["items"]) do
      temp = {
        title = json_data["group"] .. "/" .. v["host"],
        subtitle = v["description"],
        arg = cjson.encode(v),
      }
      table.insert(items, temp)
    end
  end
end
alfy.output(items)
