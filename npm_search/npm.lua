#!/usr/bin/env lua

-- local pac = {
-- 	title = "test",
-- 	subtitle = "test",
-- }

local http = require("socket.http")

local item = [[
    {
      "items": [
        {
          "title": "test",
          "subtitle": "test",
          "arg": "test"
        }
      ]
    }
]]
print(item)

print(package.path)
print(http.request("http://www.cs.princeton.edu/~diego/professional/luasocket"))
