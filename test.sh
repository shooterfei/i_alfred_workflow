#!/usr/bin/env bash


# export LUA_CPATH="./lualib/?.so"
cmd=`luarocks --lua-version 5.4 path`
eval $cmd
export LUA_PATH="./lualib/?.lua;./lualib/?/?.lua;$LUA_PATH"
time lua ./src/npm_search.lua vue
