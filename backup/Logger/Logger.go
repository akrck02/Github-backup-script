package Logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/akrck02/github-backup-script/Messages"
	"github.com/withmandala/go-log"
)

const titleCharNum = 67
const titleChar = "-"

var Logger = log.New(os.Stderr)

func ShowLogAppTitle() {
	Line()
	fmt.Println(Messages.SCRIPT_TITLE)
	Line()
}

func Log(msg string) {
	Logger.Info(msg)
}

func Error(msg string) {
	Logger.Error(msg)
}

func FormattedError(msg string, args ...string) {
	Logger.Error(Messages.Format(msg, args...))
}

func Info(msg string) {
	Logger.Info(msg)
}

func FormattedInfo(msg string, args ...string) {
	Logger.Info(Messages.Format(msg, args...))
}

func Jump() {
	Log("")
}

func Line() {
	fmt.Println(strings.Repeat(titleChar, titleCharNum))
}

func Title(title string) {
	Jump()
	Line()
	fmt.Println("   " + title)
	Line()
}

func FormattedTitle(msg string, args ...string) {
	Title(Messages.Format(msg, args...))
}

func Fatal(msg string) {
	Logger.Fatal(msg)
}
