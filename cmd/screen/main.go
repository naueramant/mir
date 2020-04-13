package main

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/jobs"
	"github.com/naueramant/mir/internal/server"
	"github.com/naueramant/mir/internal/watcher"
)

var (
	c  config.Configuration
	bm browser.BrowserManager
	js jobs.JobScheduler
)

func main() {
	go server.Start()

	go watcher.Watch("./screen.yaml", func() {
		stop()
		start()
	})

	start()

	select {} // Do not terminate
}

func start() {
	c, _ = config.Load()

	bm = browser.NewBrowserManager(c)
	go bm.Start()

	js = jobs.NewJobScheduler(c, bm)
	go js.Start()
}

func stop() {
	bm.Close()
	js.Stop()
}
