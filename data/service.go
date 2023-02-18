package main

import (
	"encoding/json"
	"github.com/akrck02/github-data-fetch/Configuration"
	"github.com/akrck02/github-data-fetch/Github"
	"github.com/akrck02/github-data-fetch/Logger"
	"os"
	"strings"
)

type GitFile struct {
	Username     string   `json:"username"`
	Repositories []string `json:"repositories"`
}

func main() {

	Logger.ShowLogAppTitle()
	var configuration = Configuration.LoadConfiguration()
	Logger.Log("Username: " + configuration.Username)
	Logger.Log("Token: " + configuration.Token)
	Logger.Log("Backup path: " + configuration.BackupPath)

	var usernames = strings.Split(configuration.Username, ",")
	var gitFiles = []GitFile{}

	for _, username := range usernames {
		Logger.Title("Fetching data for " + username)
		var gitFile = GitFile{}
		gitFile.Username = username

		var repositories = Github.GetUserRepositories(username, configuration.Token)
		gitFile.Repositories = repositories

		for _, repository := range repositories {
			Logger.Log("Repository: " + repository)
		}

		gitFiles = append(gitFiles, gitFile)

	}

	json, _ := json.Marshal(gitFiles)
	_ = os.WriteFile(configuration.BackupPath+"/git.json", json, 0644)

	Logger.Title("Finished")
}
