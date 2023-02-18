package Command

import (
	"github.com/akrck02/github-backup-script/Core"
	"os"
)

func NewDirectory(path string) *Core.Error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return Core.NewError(Core.UNEXPECTED_ERROR, "Failed to create directory: "+path+" System tells \n "+err.Error())
	}

	return Core.NewError(Core.NO_ERROR, "")
}
