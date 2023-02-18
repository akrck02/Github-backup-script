package main

import (
	"github.com/akrck02/github-backup-script/Configuration"
	"github.com/akrck02/github-backup-script/Core"
	"github.com/akrck02/github-backup-script/Logger"
	"github.com/akrck02/github-backup-script/Service"
	"github.com/akrck02/github-backup-script/Util"
	"time"
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
	Logger.Info("Elapsed time: " + elapsed.String())

}

func logConfiguration(configuration Configuration.Configuration) {
	Logger.Info("Usernames: " + configuration.Username)
	Logger.Info("Token: " + configuration.Token)
	Logger.Info("Backup path: " + configuration.BackupPath)
}

func logFinalResults(stats Core.ServiceStats) {
	Logger.Line()
	Logger.Info("Finished service.")
	Logger.Jump()
	Logger.Info("Cloned: " + Util.Int2String(stats.Cloned))
	Logger.Info("Updated: " + Util.Int2String(stats.Updated))
	Logger.Info("Failed: " + Util.Int2String(stats.Failed))
}
