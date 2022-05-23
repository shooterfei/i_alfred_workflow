#!/bin/bash


# ./lua/run.sh "$1" "$2" 

case "$1" in
  "mvn_search")
    # ./src/mvn_search/index.py $2
    ./go/bin/i-alfred-workflow -f mvn_search "$2"
  ;;
  "npm_search")
    # ./src/npm_search/index.py $2
    ./go/bin/i-alfred-workflow -f npm_search "$2"
  ;;
  "ssh")
    luajit ./src/ssh.lua "${@:2}"
  ;;
  "ssh_action")
    luajit ./src/ssh_action.lua "$2"
  ;;
  *) echo default
  ;;
esac

