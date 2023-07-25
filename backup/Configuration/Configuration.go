package Configuration

import (
	"os"

	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Messages"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Token      string
	BackupPath string
}

const TOKEN = "github.backup.token"
const BACKUP_PATH = "github.backup.path"

func LoadConfiguration() Configuration {
	err := godotenv.Load("./Resources/.env")

	if err != nil {
		Logger.Fatal("Error loading .env file")
	}

	checkCompulsoryParameters()

	return Configuration{
		Token:      os.Getenv(TOKEN),
		BackupPath: os.Getenv(BACKUP_PATH),
	}
}

func checkCompulsoryParameters() {

	if os.Getenv(BACKUP_PATH) == "" {
		Logger.Fatal(Messages.NO_BACKUP_PATH_DEFINED)
	}

}
