#!/usr/bin/env lua

local alfy = require("core.alfy")
local xmlua = require("xmlua")

local M = {}

M.search = function(g, a)
  local url = string.format("https://mvn.coderead.cn/version?groupId=%s&artifactId=%s", g, a)
  local res = alfy.fetch_html(url, {})
  local vHtmls = xmlua.HTML.parse(res):search("///tr[onclick=\"dofold($(this))\"]")
  print(xmlua.to_xml(vHtmls))
  return ""
end

return M
