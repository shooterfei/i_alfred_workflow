#!/bin/bash

export PATH=/usr/local/bin:$PATH

case "$1" in
  "mvn_search")
    # ./src/mvn_search/index.py $2
    node ./dist/bundle.mjs  -f mvn_search "$2"
  ;;
  "mvn_package_version")
    read -r -a arr <<< "${2//:/ }"
    node ./dist/bundle.mjs -f versions  -g "${arr[0]}" -a "${arr[1]}"
  ;;
  "npm_search")
    node ./dist/bundle.mjs  -f npm_search "$2"
  ;;
  "gitlab_search")
    ./i-alfred-workflow -f gitlab_search "$2"
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
