#!/usr/bin/env lua

local alfy = require("core.alfy")

local M = {}

M.search = function(q)
  local url = string.format("https://mvn.coderead.cn/search/?keyword=%s", q)
  local res = alfy.fetch(url, {})
  local data = {}
  local reg1 = "'text' >(.*)</span>.*description'>(.*)</span>.*description'>(.*)</span>"
  local reg2 = "<.?em>"
  local reg3 = "%s+"
  for _, value in ipairs(res["results"]) do
    local title, time, description = string.match(value["name"], reg1)
    title = string.gsub(string.gsub(title, reg2, ""), reg3, "")
    time = string.gsub(time, reg3, "")
    description = string.gsub(string.gsub(description, reg2, ""), reg3, "")

    local temp = {
      ["title"] = description .. ":" .. title,
      ["arg"] = description .. ":" .. title,
      ["subtitle"] = "update by " .. time,
    }
    table.insert(data, temp)
  end
  return data
end

return M
