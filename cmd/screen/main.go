package main

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/server"
)

func main() {
	go server.Start()
	go initBrowserManager()

	select {} // Do not terminate
}

func initBrowserManager() browser.BrowserManager {
	bm := browser.BrowserManager{}

	c, _ := config.Load()
	bm.Start(c)

	return bm
}
