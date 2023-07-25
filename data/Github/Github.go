package Github

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/akrck02/github-data-fetch/Logger"
)

type Repository struct {
	Name string `json:"name"`
}

func GetUserRepositories(user string, token string) []string {

	var gitInfo []string
	var currentPage []Repository
	var page = 0

	for ok := true; ok; ok = len(currentPage) > 0 {

		currentPage := []Repository{}
		var requestURL = "https://api.github.com/users/" + user + "/repos?per_page=100&page=" + strconv.FormatInt(int64(page), 10)

		Logger.Info("Fetching repositories on " + requestURL)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", requestURL, nil)

		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		resp, err := client.Do(req)

		if err != nil {
			Logger.Error("while getting user repositories on " + requestURL + " : " + err.Error())
		}

		//We Read the response body on the line below.
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			Logger.Error("while reading user repositories on " + requestURL + " : " + err.Error())
		}

		//Convert the body to type string
		sb := string(body)
		json.Unmarshal([]byte(sb), &currentPage)

		for _, repo := range currentPage {
			gitInfo = append(gitInfo, repo.Name)
		}

		page++
	}

	Logger.Info("Total repositories: " + strconv.FormatInt(int64(len(gitInfo)), 10))

	return gitInfo
}
