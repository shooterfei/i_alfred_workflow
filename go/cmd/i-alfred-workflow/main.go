package main

import (
	"flag"
	"fmt"
	"i-alfred-workflow/internal/service/search"
	"os"
)

var f = flag.String("f", "", "string类型参数")
var g = flag.String("g", "", "string类型参数")
var a = flag.String("a", "", "string类型参数")

func main() {
	flag.Parse()
	param := ""
	if len(os.Args) > 3 {
		param = os.Args[3]
	}
	switch *f {
	case "npm_search":
		fmt.Println(search.NpmSearch(param))
		break
	case "mvn_search":
		fmt.Println(search.MvnSearch(param))
		break
	case "versions":
		fmt.Println(search.Versions(*g, *a))
	case "gitlab_search":
		fmt.Println(search.GitlabSearch(param))
	default:
		fmt.Println("error")
		break
	}
}
