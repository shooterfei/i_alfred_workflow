#!/usr/bin/env bash


echo $1
while getopts "d:bc" opt
do
  case $opt in
    d)
      echo "d's arg:$OPTARG"
    ;;
    b)
      echo "b's arg:$OPTARG"
    ;;
    c)
      echo "c's arg:$OPTARG"
    ;;
    ?)
      echo "?'s arg:$OPTARG"
    ;;
  esac
done
