#!/bin/bash

export PATH=/usr/local/bin:$PATH


CWD=$(cd "$(dirname "$0")" || exit;pwd)
cd "$CWD" || exit

cmd=$(luarocks --lua-version 5.1 path)
eval "$cmd"
export LUA_PATH="./lualib/?.lua;./lualib/?/?.lua;$LUA_PATH"


case "$1" in
  "mvn_search")
    luajit mvn_search.lua "$2"
  ;;
  "npm_search")
    luajit npm_search.lua "$2"
  ;;
  "ssh")
    luajit ssh.lua "${@:2}"
  ;;
  "ssh_action")
    luajit ssh_action.lua "$2"
  ;;
  *) echo default
  ;;
esac



