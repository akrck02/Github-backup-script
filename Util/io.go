package Util

import (
	"os"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)

	if err != nil {
		return ""
	}

	return string(content)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}
