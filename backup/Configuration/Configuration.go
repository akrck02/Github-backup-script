package Configuration

import (
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Messages"
	"github.com/joho/godotenv"
	"os"
)

type Configuration struct {
	Username   string
	Token      string
	BackupPath string
}

const USERNAME = "username"
const TOKEN = "token"
const BACKUP_PATH = "backup.path"

func LoadConfiguration() Configuration {
	err := godotenv.Load("./Resources/.env")

	if err != nil {
		Logger.Fatal("Error loading .env file")
	}

	checkCompulsoryParameters()

	return Configuration{
		Username:   os.Getenv(USERNAME),
		Token:      os.Getenv(TOKEN),
		BackupPath: os.Getenv(BACKUP_PATH),
	}
}

func checkCompulsoryParameters() {

	if os.Getenv(USERNAME) == "" {
		Logger.Fatal(Messages.NO_USERNAME_DEFINED)

	}

	if os.Getenv(TOKEN) == "" {
		Logger.Fatal(Messages.NO_TOKEN_DEFINED)
	}

	if os.Getenv(BACKUP_PATH) == "" {
		Logger.Fatal(Messages.NO_BACKUP_PATH_DEFINED)
	}

}
