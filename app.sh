#!/bin/bash

export PATH=/usr/local/bin:$PATH


case "$1" in
  "mvn_search")
    ./mvn_search/index.py $2
  ;;
  "npm_search")
    # ./npm_search/npm.lua $2
    ./npm_search/index.py $2
  ;;
  *) echo default
  ;;
esac

