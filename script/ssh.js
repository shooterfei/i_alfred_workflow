#!/usr/bin/env osascript -l JavaScript

// # display dialog $arg

// on run argv 
//   set query to argv
//   log 1234444444
//   log query
//   log 99999999999999
//   tell application id "com.runningwithcrayons.Alfred" to run trigger "ssh" in workflow "shooterfei" with argument query
// end run


function run(argv) {
  console.log(argv)
  // console.log(JSON.stringify(argv))
  // let query = "{query}";
  // let [, , url] = query.split(",");

  // return url
}
