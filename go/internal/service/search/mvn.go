package search

import (
	"encoding/json"
	"fmt"
	"i-alfred-workflow/internal/conf"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func MvnSearch(q string) string {
	url := fmt.Sprintf("https://mvn.coderead.cn/search?keyword=%s", q)
	req, _ := http.NewRequest("GET", url, nil)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("http error")
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
  var mvn conf.Mvn
	jsonErr := json.Unmarshal(body, &mvn)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var alfy conf.Alfy

  regExp := regexp.MustCompile(`'text' >(.*?)<\/span>.*description'>(.*?)<\/span>.*description'>(.*?)<\/span>`)
  regExp2 := regexp.MustCompile(`<.?em>`)
	for index, v := range mvn.Results {
    result := regExp.FindAllStringSubmatch(v.Name, -1)

    title := strings.TrimSpace(regExp2.ReplaceAllString(result[0][1], ""))
    time := strings.TrimSpace(regExp2.ReplaceAllString(result[0][2], ""))
    description := strings.TrimSpace(regExp2.ReplaceAllString(result[0][3], ""))

		temp := conf.Alfy_Items{
      Uid: fmt.Sprint(index),
      Title: description + ":"+ title,
      Subtitle: fmt.Sprintf("update by %s", time),
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
