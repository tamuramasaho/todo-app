package main

import (
	"github.com/tamuramasaho/todo-app/router"
	// "github.com/robfig/cron/v3"
	// "github.com/tamuramasaho/todo-app/usecase"
	// "os"
)



func main() {
	r := router.New()
	// job := func() { usecase.NewTodo().SendReminder() }
	// cron := cron.New()
	// cron.AddFunc("CRON_TZ=Asia/Tokyo 00 9 * * *", job)
	// cron.AddFunc("CRON_TZ=Asia/Tokyo 00 12 * * *", job)	
	// cron.AddFunc("CRON_TZ=Asia/Tokyo 00 20 * * *", job)
	// cron.Start()
  r.Logger.Fatal(r.Start(":8000"))
}
