package search

import (
	"encoding/json"
	"fmt"
	"i-alfred-workflow/internal/conf"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func MvnSearchOld(q string) string {
	url := fmt.Sprintf("https://search.maven.org/solrsearch/select?q=%v&start=%d&rows=%d", q, 0, 20)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

  var mvn conf.MvnOld
	json.Unmarshal(body, &mvn)

	var alfy conf.Alfy
	for _, v := range mvn.Response.Docs {
		web_url := fmt.Sprintf("https://search.maven.org/search?q=g:%v", v.G)
		temp := conf.Alfy_Items{
			Title: fmt.Sprintf("%v:%v:%v", v.G, v.A, v.LatestVersion),
			Subtitle: fmt.Sprintf("update by %v \t count: %v",
				time.Unix(v.Timestamp/1000, 0).Format("2006-01-02 15:04:05"),
				v.VersionCount),
			Arg: web_url,
      
			Mods: map[string]*conf.Alfy_Items_Mod{
				"cmd": {
					Arg:      "test",
					Subtitle: "123",
          Title: "test",
				},
			},
		}
		alfy.Items = append(alfy.Items, &temp)
	}

	b, err := json.Marshal(alfy)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
