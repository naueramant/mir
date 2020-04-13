package main

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/jobs"
	"github.com/naueramant/mir/internal/server"
	"github.com/naueramant/mir/internal/utils"
	"github.com/sirupsen/logrus"
)

var (
	c  config.Configuration
	bm browser.BrowserManager
	js jobs.JobScheduler
)

func main() {
	go server.Start()

	go utils.Watch("./screen.yaml", func() {
		logrus.Infoln("Configuration file changed")

		stop()
		start()
	})

	start()

	select {} // Do not terminate
}

func start() {
	c, err := config.Load()
	if err != nil {
		logrus.Error(err)
	}

	bm = browser.NewBrowserManager(c)
	go bm.Start()

	js = jobs.NewJobScheduler(c, bm)
	go js.Start()
}

func stop() {
	bm.Close()
	js.Stop()
}
