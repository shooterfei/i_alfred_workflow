package search

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func Versions(groupId string, artifactId string) string {
	doc, err := goquery.NewDocument(fmt.Sprintf("http://mvn.coderead.cn/version?groupId=%s&artifactId=%s", groupId, artifactId))
	if err != nil {
		log.Fatal(err)
	}

	vHtmls := doc.Find("tr[onclick=\"doFold($(this))\"]")

	vHtmls.Each(func(i int, vHtml *goquery.Selection) {
		version := vHtml.Find("td:nth-child(1)").Text()
		downloadCount := vHtml.Find("td:nth-child(3)").Text()
		time := vHtml.Find("td:nth-child(4)").Text()
		fmt.Printf("version=%v\n", version)
		fmt.Printf("downloadCount=%v\n", downloadCount)
		fmt.Printf("time=%v\n", time)
	})

	// Each(func(i int, selection * goquery.Selection) {
	//   fmt.Println(selection.Text())
	// })
	return ""
}
