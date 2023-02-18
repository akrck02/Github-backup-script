package Command

import (
	"github.com/akrck02/github-backup-script/Logger"
	"os/exec"
)

func GitClone(url string, path string) bool {
	cmd := exec.Command("git", "clone", url)
	cmd.Dir = path
	var err = cmd.Run()

	if err != nil {
		Logger.Error(err.Error())
		return false
	}

	return true
}

func GitPull(path string) bool {
	cmd := exec.Command("git", "pull")
	cmd.Dir = path
	var err = cmd.Run()
	if err != nil {
		Logger.Error(err.Error())
		return false
	}

	return true
}
