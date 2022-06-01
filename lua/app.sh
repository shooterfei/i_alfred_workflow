#!/bin/bash


export PATH=/usr/local/bin:$PATH


CWD=$(cd "$(dirname "$0")" || exit;pwd)
cd "$CWD" || exit

cmd=$(luarocks --lua-version 5.1 path)
eval "$cmd"

export LUA_PATH="./lualib/?.lua;./lualib/?/?.lua;$LUA_PATH"

case "$1" in
  "mvn_search")
    luajit main.lua -f "$1" "$2"
  ;;
  "mvn_package_version")
    read -r -a arr <<< "${2//:/ }"
    luajit main.lua -f versions  -g "${arr[0]}" -a "${arr[1]}"
  ;;
  "npm_search")
    luajit main.lua -f "$1" "$2"
  ;;
  "gitlab_search")
    ./i-alfred-workflow -f gitlab_search "$2"
  ;;
  "ssh")
    luajit main.lua -f "$1" "${@:2}"
  ;;
  "ssh_action")
    luajit main.lua -f "$1" "$2"
  ;;
  *) echo default
  ;;
esac
