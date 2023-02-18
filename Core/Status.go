package Core

import (
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Util"
)

const NO_ERROR = 0
const UNEXPECTED_ERROR = 1

type Error struct {
	Message string
	Code    int
}

type ServiceStats struct {
	Failed  int
	Updated int
	Cloned  int
}

func NewError(code int, message string) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func Stop(error *Error) {
	Logger.Error(formatError(error))
}

func formatError(error *Error) string {
	return "#" + Util.Int2String(error.Code) + " | " + error.Message
}
