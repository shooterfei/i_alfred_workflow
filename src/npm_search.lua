#!/usr/bin/env lua

local alfy = require("core.alfy")

local url = "https://www.npmjs.com/search?q=" .. arg[1]

local res = alfy.fetch(url, {
	headers = {
		["x-spiferack"] = "1",
	},
})

local data = {}

for k, v in ipairs(res["objects"]) do
	local score = v["score"]["detail"]
	local temp = {
		["title"] = v["package"]["name"],
		["subtitle"] = string.format(
			"update by %s \t p: %.2f q: %.2f m: %.2f",
			os.date("%Y-%m-%d %H:%M:%S", v["package"]["date"]["ts"] // 1000),
			score["popularity"],
			score["quality"],
			score["maintenance"]
		),
		["arg"] = v["package"]["links"]["npm"],
	}
	table.insert(data, temp)
end

alfy.output(data)
