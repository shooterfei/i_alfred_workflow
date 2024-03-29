#!/bin/bash


case "$1" in
  "mvn_search")
    # ./src/mvn_search/index.py $2
    ./i-alfred-workflow -f mvn_search "$2"
  ;;
  "mvn_package_version")
    read -r -a arr <<< "${2//:/ }"
    ./i-alfred-workflow -f versions  -g "${arr[0]}" -a "${arr[1]}"
  ;;
  "npm_search")
    # ./src/npm_search/index.py $2
    ./i-alfred-workflow -f npm_search "$2"
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

