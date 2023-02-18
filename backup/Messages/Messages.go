package Messages

import (
	"strconv"
	"strings"
)

// App messages
const SCRIPT_TITLE = "Github backup script by @akrck02"
const EASTER_EGG = "She doesn't love me, let's code T.T"

// Service messages
const ELAPSED_TIME = "Elapsed time: "
const FINISHED_SERVICE = "Finished service."
const CLONED = "Cloned: ${0}"
const UPDATED = "Updated: ${0}"
const FAILED = "Failed: ${0}"
const USERNAMES = "Usernames: ${0}"
const TOKEN = "Token: ${0}"
const BACKUP_PATH = "Backup path: "

// Github messages
const GIT_ACCOUNTS = "Git accounts: ${0}"
const UPDATED_REPOSITORY = "Updated repository: ${0} / ${1}"
const FAILED_UPDATE_REPOSITORY = "Failed to update repository: ${0} / ${1}"
const CLONED_REPOSITORY = "Cloned repository: ${0} / ${1}"
const FAILED_CLONE_REPOSITORY = "Failed to clone repository: ${0} / ${1}"
const CLONING_REPOSITORY = "Cloning repository: ${0}/${1}"
const UPDATING_REPOSITORY = "Updating repository: ${0}/${1}"
const GITHUB_REPOSITORY = "https://github.com/${0}/${1}"

// Configuration messages
const ERROR_LOADING_ENV_FILE = "Error loading .env file"
const NO_USERNAME_DEFINED = "No Username defined"
const NO_TOKEN_DEFINED = "No Token defined"
const NO_BACKUP_PATH_DEFINED = "No Backup path defined"

func Format(message string, args ...string) string {
	var i int

	for i = 0; i < len(args); i++ {
		message = strings.Replace(message, "${"+int2String(i)+"}", args[i], -1)
	}

	return message
}

func int2String(num int) string {
	return strconv.FormatInt(int64(num), 10)
}
