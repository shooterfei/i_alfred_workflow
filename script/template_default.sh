#!/usr/bin/expect

set host $env(host)
set port $env(port)
set username $env(username)
set password $env(password)
set root $env(root)
set timeout $env(timeout)


spawn ssh $username@$host -p $port -t "cd $root; bash --login"


expect {
 "yes/no" {send "yes\r"; exp_continue }
 "password:" {send "$password"}
}

interact
