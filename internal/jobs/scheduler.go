package jobs

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
)

type JobScheduler struct {
	BrowserManager browser.BrowserManager
}

func (s *JobScheduler) Start(c config.Configuration) {

}

func (s *JobScheduler) Stop() {

}
