package Service

import (
	"encoding/json"
	"github.com/akrck02/github-backup-script/Command"
	"github.com/akrck02/github-backup-script/Configuration"
	"github.com/akrck02/github-backup-script/Core"
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Util"
)

type GitBackup struct {
	Username     string   `json:"username"`
	Repositories []string `json:"repositories"`
}

func Backup(configuration Configuration.Configuration, stats Core.ServiceStats) Core.ServiceStats {
	var err = Command.NewDirectory(configuration.BackupPath)

	if err.Code != Core.NO_ERROR {
		Core.Stop(err)
		return stats
	}

	Logger.Jump()
	Logger.Info("She doesn't love me, let's code T.T")
	Logger.Jump()

	var gitInfoFile = Util.ReadFile("./Resources/git.json")
	var gitInfo []GitBackup

	json.Unmarshal([]byte(gitInfoFile), &gitInfo)
	Logger.Info("GIT ACCOUNTS: " + Util.Int2String(len(gitInfo)))

	for _, account := range gitInfo {

		var accountPath = configuration.BackupPath + "/" + account.Username
		Command.NewDirectory(accountPath)

		for _, repo := range account.Repositories {

			var filePath = accountPath + "/" + repo

			if Util.FileExists(filePath) {

				var updated = updateRepository(filePath, account.Username, repo)

				if updated {
					stats.Updated++
					Logger.Info("Updated " + account.Username + "/" + repo)
				} else {
					stats.Failed++
					Logger.Error("Failed to update " + account.Username + "/" + repo)
				}
				continue
			}

			var cloned = cloneRepository(accountPath, account.Username, repo)
			if cloned {
				stats.Cloned++
				Logger.Info("Cloned " + account.Username + "/" + repo)
			} else {
				stats.Failed++
				Logger.Error("Failed to clone " + account.Username + "/" + repo)
			}

		}
	}

	return stats
}

func cloneRepository(url string, account string, repository string) bool {
	Logger.Title("Cloning " + account + "/" + repository)

	var githubRepo = "https://github.com/" + account + "/" + repository
	var cloned = Command.GitClone(githubRepo, url)

	return cloned
}

func updateRepository(url string, account string, repo string) bool {
	Logger.Title("Updating " + account + "/" + repo)

	var updated = Command.GitPull(url)

	return updated
}
