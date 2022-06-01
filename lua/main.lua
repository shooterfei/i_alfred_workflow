#!/usr/bin/env luajit

local cjson = require("cjson")

local opt = {}

_G.string.starts = function(String, Start)
	return string.sub(String, 1, string.len(Start)) == Start
end

local lens = 0;
for index, value in ipairs(arg) do
	if _G.string.starts(value, "-") == true then
		opt[value] = arg[index + 1]
	end
  lens = lens + 1
end

local q = arg[lens]

local _switch = {
	["mvn_search"] = function()
    return require("src.mvn_search").search(q)
	end,
	["versions"] = function()
    return require("src.mvn_package_version").search(opt["-g"], opt["-a"])
	end,
	["npm_search"] = function()
    return require("src.npm_search").search(q)
	end,
}

print(cjson.encode(_switch[opt["-f"]]()))

