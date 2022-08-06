package search

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"i-alfred-workflow/internal/conf"
	"log"
)

func Versions(groupId string, artifactId string) string {
	doc, err := goquery.NewDocument(fmt.Sprintf("http://mvn.coderead.cn/version?groupId=%s&artifactId=%s", groupId, artifactId))
	if err != nil {
		log.Fatal(err)
	}

	vHtmls := doc.Find("tr[onclick=\"doFold($(this))\"]")
	var alfy conf.Alfy

	vHtmls.Each(func(i int, vHtml *goquery.Selection) {
		version := vHtml.Find("td:nth-child(1)").Text()
		downloadCount := vHtml.Find("td:nth-child(3)").Text()
		time := vHtml.Find("td:nth-child(4)").Text()

		temp := conf.Alfy_Items{
			Title:    version,
			Subtitle: fmt.Sprintf("update by %s \t downloadCount: %s", time, downloadCount),
			Arg:      fmt.Sprint(i),
			Mods: map[string]*conf.Alfy_Items_Mod{
				"ctrl": {
					Arg: fmt.Sprintf("<dependency>\n\t<groupid>%s</groupid>\n\t<artifactId>%s</artifactId>\n\t<version>%s</version>\n</dependency>",
						groupId,
						artifactId,
						version),
					Subtitle: "copy for Maven",
				},
				"cmd": {
					Arg:      fmt.Sprintf("implementation '%s:%s:%s'", groupId, artifactId, version),
					Subtitle: "copy for Grandle",
				},
				"alt": {
					Arg:      fmt.Sprintf("implementation('%s:%s:%s')", groupId, artifactId, version),
					Subtitle: "copy for Kotlin",
				},
			},
		}
		alfy.Items = append(alfy.Items, &temp)
	})

	b, err := json.Marshal(alfy)
	return string(b)
}
