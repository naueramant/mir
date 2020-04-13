package utils

import (
	"time"

	"github.com/radovskyb/watcher"
)

func Watch(file string, changedFunc func()) {
	var w *watcher.Watcher
	var err error

	var start func()
	var recover func()
	var eventLoop func()

	retrigger := false

	recover = func() {
		retrigger = true
		time.Sleep(time.Millisecond * 500)
		start()
	}

	eventLoop = func() {
		for {
			select {
			case <-w.Event:
				go changedFunc()
			case <-w.Error:
				go changedFunc()
				recover()
			case <-w.Closed:
				go changedFunc()
				recover()
			}
		}
	}

	start = func() {
		w = watcher.New()
		w.FilterOps(watcher.Write, watcher.Remove, watcher.Rename, watcher.Move, watcher.Create)

		err = w.Add(file)
		if err != nil {
			recover()
		}

		go eventLoop()

		if retrigger {
			changedFunc()
		}

		err = w.Start(time.Millisecond * 100)
		if err != nil {
			recover()
		}
	}

	start()
}
