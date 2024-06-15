package main

import (
	"flag"

	"github.com/naueramant/mir/internal/assets"
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/jobs"
	"github.com/naueramant/mir/internal/utils"
	"github.com/sirupsen/logrus"
)

var (
	c  *config.Configuration
	bm *browser.BrowserManager
	js *jobs.JobScheduler

	configPath = flag.String("config", "screen.yaml", "path to screen configuration")
)

func main() {
	flag.Parse()

	logrus.Infof("Using configuration file %s", *configPath)

	go utils.Watch(*configPath, func() {
		logrus.Infoln("Configuration file changed")

		stop()
		start()
	})

	start()

	select {} // Do not terminate
}

func start() {
	c, err := config.Load(*configPath)
	if err != nil {
		logrus.Error(err)
	}

	as, err := assets.NewServer()
	if err != nil {
		logrus.Error(err)
	}

	go func() {
		if err := as.Start(); err != nil {
			logrus.WithError(err).Fatal("Failed to start assets server")
		}
	}()

	bm = browser.NewBrowserManager(c, as)
	go bm.Start()

	js = jobs.NewJobScheduler(c, bm, as)
	go js.Start()
}

func stop() {
	bm.Close()
	js.Stop()
}
