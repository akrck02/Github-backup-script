package main

import (
	"time"

	"github.com/akrck02/github-backup-script/Configuration"
	"github.com/akrck02/github-backup-script/Core"
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Messages"
	"github.com/akrck02/github-backup-script/Service"
	"github.com/akrck02/github-backup-script/Util"
)

func main() {

	start := time.Now()

	Logger.ShowLogAppTitle()

	var stats = Core.ServiceStats{
		Failed:  0,
		Updated: 0,
		Cloned:  0,
	}

	var configuration = Configuration.LoadConfiguration()

	logConfiguration(configuration)
	stats = Service.Backup(configuration, stats)
	logFinalResults(stats)

	end := time.Now()
	elapsed := end.Sub(start)

	Logger.Line()
	Logger.Info(Messages.ELAPSED_TIME + elapsed.String())

}

func logConfiguration(configuration Configuration.Configuration) {
	Logger.FormattedInfo(Messages.TOKEN, configuration.Token)
	Logger.FormattedInfo(Messages.BACKUP_PATH, configuration.BackupPath)
}

func logFinalResults(stats Core.ServiceStats) {
	Logger.Line()
	Logger.Info(Messages.FINISHED_SERVICE)
	Logger.Jump()
	Logger.FormattedInfo(Messages.CLONED, Util.Int2String(stats.Cloned))
	Logger.FormattedInfo(Messages.UPDATED, Util.Int2String(stats.Updated))
	Logger.FormattedInfo(Messages.FAILED, Util.Int2String(stats.Failed))
}
