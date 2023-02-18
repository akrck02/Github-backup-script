package Configuration

import (
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/joho/godotenv"
	"os"
)

type Configuration struct {
	Username   string
	Token      string
	BackupPath string
}

func LoadConfiguration() Configuration {
	err := godotenv.Load("./Resources/.env")

	if err != nil {
		Logger.Error("Error loading .env file")
		os.Exit(1)
	}

	checkCompulsoryParameters()

	return Configuration{
		Username:   os.Getenv("username"),
		Token:      os.Getenv("token"),
		BackupPath: os.Getenv("backup.path"),
	}
}

func checkCompulsoryParameters() {

	if os.Getenv("username") == "" {
		Logger.Error("No Username defined")
		os.Exit(1)
	}

	if os.Getenv("token") == "" {
		Logger.Error("No Token defined")
		os.Exit(1)
	}

	if os.Getenv("backup.path") == "" {
		Logger.Error("No Backup path defined")
		os.Exit(1)
	}

}
