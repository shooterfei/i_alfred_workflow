#!/usr/bin/env lua

local alfy = require("core.alfy")
local xmlua = require("xmlua")

local M = {}

M.search = function(groupId, artifactId)
  local url = string.format("https://mvn.coderead.cn/version?groupId=%s&artifactId=%s", groupId, artifactId)
  local res = alfy.fetch_html(url, {})
  local data = {}

  local vHtmls = xmlua.HTML.parse(res):search("//tr[@onclick='doFold($(this))']")

  for index, vHtml in ipairs(vHtmls) do
    local tds = vHtml:search("td")
    local version = tds[1]:content()
    local downloadCount = tds[3]:content()
    local time = tds[4]:content()
    local temp = {
      ["title"] = version,
      ["arg"] = index,
      ["subtitle"] = string.format("update by%s \t downloadCount: %s", time, downloadCount),
      ["mods"] = {
        ["ctrl"] = {
          ["arg"] = string.format("<dependency>\n\t<groupid>%s</groupid>\n\t<artifactid>%s</artivactid>\n\t<version>%s</version>\n</dependency>",
            groupId,
            artifactId,
            version),
          ["subtitle"] = "copy for Maven",
        },
        ["cmd"] = {
          ["arg"] = string.format("implementation '%s:%s:%s'", groupId, artifactId, version),
          ["subtitle"] = "copy for Grandle",
        },
        ["alt"] = {
          ["arg"] = string.format("implementation('%s:%s:%s')", groupId, artifactId, version),
          ["subtitle"] = "copy for Kotlin",
        },
      },
    }
    table.insert(data, temp)
  end

  return data
end

return M
