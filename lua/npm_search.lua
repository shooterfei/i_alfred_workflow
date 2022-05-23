local alfy = require("core.alfy")

local url = "https://www.npmjs.com/search?q=" .. arg[1]

local res = alfy.fetch(url, {
	headers = {
		["x-spiferack"] = "1",
	},
})

local data = {}

for _, v in ipairs(res["objects"]) do
	local score = v["score"]["detail"]

	local temp = {
		["title"] = v["package"]["name"],
		["subtitle"] = string.format(
			"update by %s \t p: %.2f q: %.2f m: %.2f",
			os.date("%Y-%m-%d %H:%M:%S", math.floor(v["package"]["date"]["ts"] / 1000)),
			score["popularity"],
			score["quality"],
			score["maintenance"]
		),
		["arg"] = v["package"]["links"]["npm"],
		["mods"] = {
      ["ctrl"] = {
        ["arg"] = "npm install " .. v["package"]["name"],
        ["subtitle"] = "copy for npm"
      },
			["cmd"] = {
				["arg"] = "yarn add " .. v["package"]["name"],
				["subtitle"] = "copy for yarn",
			},
			["alt"] = {
				["arg"] = "pnpm install " .. v["package"]["name"],
				["subtitle"] = "copy for pnpm",
			},
		},
	}
	table.insert(data, temp)
end

alfy.output(data)
