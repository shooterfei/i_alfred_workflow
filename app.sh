#!/bin/bash

export PATH=/usr/local/bin:$PATH


case "$1" in
  "mvn_search")
    # ./src/mvn_search/index.py $2
    ./src/mvn_search.lua $2
  ;;
  "npm_search")
    ./src/npm_search.lua $2
    # ./npm_search/index.py $2
  ;;
  "ssh")
    ./src/ssh.lua "${@:2}"
    # ./npm_search/index.py $2
  ;;
  "ssh_action")
    ./src/ssh_action.lua $2
    # ./npm_search/index.py $2
  ;;
  *) echo default
  ;;
esac

