package Logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/withmandala/go-log"
)

const titleCharNum = 67
const titleChar = "-"

var Logger = log.New(os.Stderr)

func ShowLogAppTitle() {
	Line()
	fmt.Println(`
		Github backup script by @akrck02																		
	`)
	Line()
}

func Log(msg string) {
	Logger.Info(msg)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Info(msg string) {
	Logger.Info(msg)
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
