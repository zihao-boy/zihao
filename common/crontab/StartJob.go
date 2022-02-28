package crontab

func StartCrontab()  {
	var (
		monitorJob MonitorJob
		backUpJob BackUpJob
	)
	monitorJob.Restart()
	backUpJob.Restart()
}
