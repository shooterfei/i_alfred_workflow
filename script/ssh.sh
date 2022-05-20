#!/usr/bin/env bash

# export host=192.168.3.41
# export port=22
# export username=dl
# export password=zeusight
# export root="~"
# export timeout=10
func() {
    echo "u: username"
    echo "a: ip addr"
    echo "p: port"
    echo "d: rootdir"
    echo "w: password"
    echo "t: username"
    exit 0
}


while getopts 'u:host:p:r:password:t' OPT; do
    case $OPT in
        u) export host=$OPTARG;;
        a) export port=$OPTARG;;
        p) export port=$OPTARG;;
        d) export port=$OPTARG;;
        w) export port=$OPTARG;;
        t) export port=$OPTARG;;
        ?) func;;
    esac
done


# ./template_default.exp 
