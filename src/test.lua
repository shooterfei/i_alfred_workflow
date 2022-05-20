local cjson = require("cjson")
local http = require("socket.http")
local ltn12 = require("ltn12")
local io = require("io")
--引入mine包
local mime = require("mime")


local fetch = function(url, body)
	local response_body = {}
	body["url"] = url
	body["sink"] = ltn12.sink.table(response_body)
	local res = { http.request(body) }
	if res[2] == 200 then
		if type(response_body) == "table" then
			return cjson.decode(table.concat(response_body))
		end
	else
		return {
			["code"] = false,
		}
	end
end



local fh, errStr = io.open( arg[1], "rb" )

local contents = fh:read( "*a" )

--base64编码
local enc = mime.b64(contents)

--Post传输数据时，[+]号会被转换成空格，导致服务端解码出错
--这边做个转换，替换成[%2B]。这边要用转义字符[%]，否则报错
enc = string.gsub(enc,"+","%%2B")
print(enc)

io.close(fh)


-- --返回数据
-- local function networkIndex(event)
-- 	print(event.response)
-- end

-- local params = {}
-- params.body = "data=" .. enc

local body = {
  method = "POST",
  body = {
    format = "json",
    data = enc
  }
}

local res = fetch(arg[2], body)
print(cjson.encode(res))
-- --POST 数据
