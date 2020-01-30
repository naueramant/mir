package watcher

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/radovskyb/watcher"
)

func Watch(file string, changedFunc func()) error {
	w := watcher.New()
	var err error

	w.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-w.Event:
				changedFunc()
			case <-w.Error:
				fmt.Println("No such file " + file)
				os.Exit(1)
			case <-w.Closed:
				return
			}
		}
	}()

	if err = w.Add(file); err != nil {
		return errors.New("No such file " + file)
	}

	return w.Start(time.Millisecond * 100)
}
