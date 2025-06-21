package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.AddCronJob("*/10 * * * * *", "log-time", func(ctx *gofr.Context) {
		ctx.Logger.Infof("Current time : %v", time.Now())
	})

	app.AddCronJob("0 */5 * * *", "backup-reminder", func(ctx *gofr.Context) {
		ctx.Logger.Infof("BACKUP REMINDER")
	})

	app.Run()
}
