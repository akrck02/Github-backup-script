package Configuration

import (
	"os"

	"github.com/akrck02/github-data-fetch/Logger"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Username string
	Token    string
	JsonPath string
}

const USERNAMES = "github.fetch.usernames"
const TOKEN = "github.fetch.token"
const JSON_PATH = "github.fetch.json.path"

func LoadConfiguration() Configuration {
	err := godotenv.Load("./Resources/.env")

	if err != nil {
		Logger.Error("Error loading .env file")
		os.Exit(1)
	}

	checkCompulsoryParameters()

	return Configuration{
		Username: os.Getenv(USERNAMES),
		Token:    os.Getenv(TOKEN),
		JsonPath: os.Getenv(JSON_PATH),
	}
}

func checkCompulsoryParameters() {

	if os.Getenv(USERNAMES) == "" {
		Logger.Error("No Username defined")
		os.Exit(1)
	}

	if os.Getenv(JSON_PATH) == "" {
		Logger.Error("No JSON path defined")
		os.Exit(1)
	}

}
