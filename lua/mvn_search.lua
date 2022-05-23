#!/usr/bin/env lua

local alfy = require("core.alfy")

local url = string.format("https://search.maven.org/solrsearch/select?q=%s&start=%s&rows=%s", arg[1], 0, 20)

local res = alfy.fetch(url, {})

local data = {}

for _, value in ipairs(res["response"]["docs"]) do
	local g, a, v = value["g"], value["a"], value["v"] == nil and value["latestVersion"] or value["v"]
	local mvn = string.format(
		"<dependency>\n  <groupid>%s</groupid>\n  <artifactid>%s</artifactid>\n  <version>%s</version>\n</dependency>",
		g,
		a,
		v
	)
	local web_url = "https://search.maven.org/search?q=g:" .. g
	local gradle = string.format("implementation('%s:%s:%s')", g, a, v)
	local temp = {
		["title"] = g .. ":" .. a .. ":" .. v,
		["arg"] = web_url,
		["subtitle"] = string.format(
			"update at %s \t count: %s",
			os.date("%Y-%m-%d %H:%M:%S", math.floor(value["timestamp"] / 1000)),
			value["versionCount"]
		),
		["mods"] = {
			["alt"] = {
				["arg"] = gradle,
				["subtitle"] = "copy gradle dependency to clipboard",
			},
			["cmd"] = {
				["arg"] = mvn,
				["subtitle"] = "copy maven dependency to clipboard",
			},
		},
	}
	table.insert(data, temp)
end

alfy.output(data)
