package main

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/jobs"
	"github.com/naueramant/mir/internal/server"
)

var (
	c  config.Configuration
	bm browser.BrowserManager
	js jobs.JobScheduler
)

func main() {
	go server.Start()

	initAll()

	select {} // Do not terminate
}

func initAll() {
	c, _ = config.Load()

	initBrowserManager()
	initJobSchedular()
}

func initBrowserManager() {
	bm = browser.NewBrowserManager(c)
	bm.Start()
}

func initJobSchedular() {
	js = jobs.NewJobScheduler(c, bm)
	js.Start()
}
