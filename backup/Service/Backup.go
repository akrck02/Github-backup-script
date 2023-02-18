package Service

import (
	"encoding/json"
	"github.com/akrck02/github-backup-script/Command"
	"github.com/akrck02/github-backup-script/Configuration"
	"github.com/akrck02/github-backup-script/Core"
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Messages"
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
	Logger.Info(Messages.EASTER_EGG)
	Logger.Jump()

	var gitInfoFile = Util.ReadFile("./Resources/git.json")
	var gitInfo []GitBackup

	json.Unmarshal([]byte(gitInfoFile), &gitInfo)
	Logger.FormattedInfo(Messages.GIT_ACCOUNTS, Util.Int2String(len(gitInfo)))

	for _, account := range gitInfo {

		var accountPath = configuration.BackupPath + "/" + account.Username
		Command.NewDirectory(accountPath)

		for _, repo := range account.Repositories {

			var filePath = accountPath + "/" + repo

			if Util.FileExists(filePath) {

				var updated = updateRepository(filePath, account.Username, repo)

				if updated {
					stats.Updated++
					Logger.FormattedInfo(Messages.UPDATED_REPOSITORY, account.Username, repo)
				} else {
					stats.Failed++
					Logger.FormattedError(Messages.FAILED_UPDATE_REPOSITORY, account.Username, repo)
				}
				continue
			}

			var cloned = cloneRepository(accountPath, account.Username, repo)
			if cloned {
				stats.Cloned++
				Logger.FormattedInfo(Messages.CLONED_REPOSITORY, account.Username, repo)
			} else {
				stats.Failed++
				Logger.FormattedError(Messages.FAILED_CLONE_REPOSITORY, account.Username, repo)
			}

		}
	}

	return stats
}

func cloneRepository(url string, account string, repository string) bool {
	Logger.FormattedTitle(Messages.CLONING_REPOSITORY, account, repository)

	var githubRepo = Messages.Format(Messages.GITHUB_REPOSITORY, account, repository)
	var cloned = Command.GitClone(githubRepo, url)

	return cloned
}

func updateRepository(url string, account string, repo string) bool {
	Logger.FormattedTitle(Messages.UPDATING_REPOSITORY, account, repo)
	var updated = Command.GitPull(url)
	return updated
}
