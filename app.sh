#!/bin/bash

export PATH=/usr/local/bin:$PATH

cmd=`luarocks --lua-version 5.1 path`
eval $cmd
export LUA_PATH="./lualib/?.lua;./lualib/?/?.lua;$LUA_PATH"

case "$1" in
  "mvn_search")
    # ./src/mvn_search/index.py $2
    luajit ./src/mvn_search.lua $2
  ;;
  "npm_search")
    luajit ./src/npm_search.lua $2
    # ./src/npm_search/index.py $2
  ;;
  "ssh")
    luajit ./src/ssh.lua "${@:2}"
  ;;
  "ssh_action")
    luajit ./src/ssh_action.lua $2
  ;;
  *) echo default
  ;;
esac

