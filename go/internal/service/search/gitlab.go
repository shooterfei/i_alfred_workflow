package search

import (
	"encoding/json"
	"fmt"
	"i-alfred-workflow/internal/conf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GitlabSearch(q string) string {
	var gitlab_host, gitlab_port, token, protocol string = os.Getenv("GITLAB_HOST"),
		os.Getenv("GITLAB_PORT"),
		os.Getenv("GITLAB_ACCESS_TOKEN"),
		os.Getenv("PROTOCOL")

	if gitlab_host == "" {
		gitlab_host = "gitlab.shooterfei.com"
	}

	if gitlab_port == "" {
		gitlab_port = "443"
	}

	if token == "" {
		token = "oUAZ5rsNx5PM_MvChJ2L"
	}

	if protocol == "" {
		protocol = "https"
	}

	url := fmt.Sprintf("%s://%s:%s/api/v4/projects.json?search=%s", protocol, gitlab_host, gitlab_port, q)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("PRIVATE-TOKEN", token)

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("http error")
	}
	defer resp.Body.Close()


	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var gitlabs []conf.GitLab
	jsonErr := json.Unmarshal(body, &gitlabs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	var alfy conf.Alfy

	for index, v := range gitlabs {
		// result := regExp.FindAllStringSubmatch(v.Name, -1)

		// title := strings.TrimSpace(regExp2.ReplaceAllString(result[0][1], ""))
		// time := strings.TrimSpace(regExp2.ReplaceAllString(result[0][2], ""))
		// description := strings.TrimSpace(regExp2.ReplaceAllString(result[0][3], ""))

		temp := conf.Alfy_Items{
			Uid:      fmt.Sprint(index),
			Title:    v.PathWithNamespace,
			Subtitle: v.Description,
			Arg:      v.WebUrl,
			Mods: map[string]*conf.Alfy_Items_Mod{
				"cmd": {
					Arg:      v.SshUrlToRepo,
					Subtitle: "copy ssh to clipboard",
				},
				"alt": {
					Arg:      v.HttpUrlToRepo,
					Subtitle: "copy http to clipboard",
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
