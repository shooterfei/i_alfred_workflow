package main

import (
	"flag"
	"fmt"
	"os"
)

var f = flag.String("f", "", "string类型参数")

func main() {
	flag.Parse()
	param := ""
	if len(os.Args) > 3 {
		param = os.Args[3]
	}
	switch *f {
	case "npm_search":
		fmt.Println(npmSearch(param))
		break
	default:
		fmt.Println("error")
		break
	}
}
