package main

import (
	"encoding/json"
	"fmt"
	"i-alfred-workflow/internal/conf"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func npmSearch(q string) string {
	url := fmt.Sprintf("https://www.npmjs.com/search?q=%s", q)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("x-spiferack", "1")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("http error")
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var npm conf.Npm
	jsonErr := json.Unmarshal(body, &npm)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	time.Now()
	// items := make([]*conf.Alfy_Items, 0)
	var alfy conf.Alfy
	for _, v := range npm.Objects {
		temp := conf.Alfy_Items{
			Title: v.Package.Name,
			Subtitle: fmt.Sprintf("update by %v \t p: %.2f q: %.2f m: %.2f",
				time.Unix(v.Package.Date.Ts / 1000, 0).Format("2006-01-02 15:04:05"),
				v.Score.Detail.Popularity,
				v.Score.Detail.Quality,
				v.Score.Detail.Maintenance),
			Arg: "test",
		}
		alfy.Items = append(alfy.Items, &temp)
	}
	b, err := json.Marshal(alfy)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
