package main

import (
	"time"

	"github.com/labstack/gommon/log"
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/server"
	"github.com/radovskyb/watcher"
)

func main() {
	go server.Start()

	w := watcher.New()

	w.FilterOps(watcher.Write, watcher.Move, watcher.Create, watcher.Remove, watcher.Rename)

	go func() {
		bm := initBrowserManager()

		for {
			select {
			case <-w.Event:
				bm.Stop()
				bm = initBrowserManager()
			case err := <-w.Error:
				log.Fatal(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(config.Filename); err != nil {
		log.Fatal(err)
	}

	if err := w.Start(time.Second * 1); err != nil {
		log.Fatal(err)
	}
}

func initBrowserManager() browser.BrowserManager {
	bm := browser.BrowserManager{}

	c, _ := config.Load()
	bm.Start(c)

	return bm
}
