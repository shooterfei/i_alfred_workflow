#!/bin/bash


export PATH=/usr/local/bin:$PATH


CWD=$(cd "$(dirname "$0")" || exit;pwd)
cd "$CWD" || exit

cmd=$(luarocks --lua-version 5.1 path)
eval "$cmd"

export LUA_PATH="./lualib/?.lua;./lualib/?/?.lua;$LUA_PATH"

case "$1" in
  "mvn_search")
    luajit src/mvn_search.lua "$2"
  ;;
  "mvn_package_version")
    read -r -a arr <<< "${2//:/ }"
    ./i-alfred-workflow -f versions  -g "${arr[0]}" -a "${arr[1]}"
  ;;
  "npm_search")
    luajit src/npm_search.lua "$2"
  ;;
  "gitlab_search")
    ./i-alfred-workflow -f gitlab_search "$2"
  ;;
  "ssh")
    luajit src/ssh.lua "${@:2}"
  ;;
  "ssh_action")
    luajit src/ssh_action.lua "$2"
  ;;
  *) echo default
  ;;
esac
