package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Every 10 seconds: print timestamp
	app.AddCronJob("*/10 * * * * *", "timestamp-logger", logCurrentTime)

	// Every 5 minutes: show backup reminder
	app.AddCronJob("0 */5 * * *", "remind-backup", showBackupReminder)

	// Every 30 seconds: log heartbeat
	app.AddCronJob("*/30 * * * * *", "heartbeat", logHeartbeat)

	app.Run()
}

// Custom log job: current time
func logCurrentTime(ctx *gofr.Context) {
	now := time.Now().Format("2006-01-02 15:04:05")
	ctx.Logger.Infof("[TimestampLogger] Current time: %v", now)
}

// Custom log job: backup reminder
func showBackupReminder(ctx *gofr.Context) {
	ctx.Logger.Warnf("[BackupReminder] ⏳ Time to check your backup strategy!")
}

// Custom log job: heartbeat signal
func logHeartbeat(ctx *gofr.Context) {
	ctx.Logger.Infof("[Heartbeat] ✅ App is alive @ %v", time.Now().Unix())
}
