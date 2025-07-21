package initialize

import "github.com/robfig/cron/v3"

func Startup() {
	cron.New(cron.WithSeconds())
}
