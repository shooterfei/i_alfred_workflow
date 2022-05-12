#!/usr/bin/env node


const https = require("https")

// const fetch = require("node-fetch")
// let pac = {
//   items: {
//     title: "test",
//     subtitle: "test",
//     arg: "test"
//   }
// }

// console.log(JSON.stringify(pac))

let url = "https://www.npmjs.com/search?q=vue"



const options = {
  hostname: "www.npmjs.com",
  port: 443,
  path: "/search",
  method: 'GET'
}

https.requests(options, res => {
  console.log(`状态码: ${res.statusCode}`)
})

// let search = "vue"

// let params = {
//      'q': search,
//      'page': 0,
//      'perPage': 20
//  }


// fetch(url, {

// }).then(res => res.json())
//   .then(res => {
//     console.log(res)
//   })


 // res = fetch.get(url, params=params)
 // url = "https://www.npmjs.com/search"

 // params = {
 //     'q': search,
 //     'page': 0,
 //     'perPage': 20
 // }

 // res = requests.get(url, params=params)

