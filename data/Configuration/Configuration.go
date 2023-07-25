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

func LoadConfiguration() Configuration {
	err := godotenv.Load("./Resources/.env")

	if err != nil {
		Logger.Error("Error loading .env file")
		os.Exit(1)
	}

	checkCompulsoryParameters()

	return Configuration{
		Username: os.Getenv("github.fetch.username"),
		Token:    os.Getenv("github.fetch.token"),
		JsonPath: os.Getenv("github.json.path"),
	}
}

func checkCompulsoryParameters() {

	if os.Getenv("github.fetch.username") == "" {
		Logger.Error("No Username defined")
		os.Exit(1)
	}

	if os.Getenv("github.json.path") == "" {
		Logger.Error("No JSON path defined")
		os.Exit(1)
	}

}
