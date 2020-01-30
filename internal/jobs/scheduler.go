package jobs

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/robfig/cron/v3"
)

type JobScheduler struct {
	BrowserManager browser.BrowserManager
	Config         config.Configuration
	Cron           *cron.Cron
}

func NewJobScheduler(config config.Configuration, browserManager browser.BrowserManager) JobScheduler {
	return JobScheduler{
		BrowserManager: browserManager,
		Config:         config,
		Cron:           cron.New(),
	}
}

func (s *JobScheduler) Start() {
	for _, j := range s.Config.Jobs {
		s.schedule(j)
	}

	s.Cron.Start()
}

func (s *JobScheduler) Stop() {
	s.Cron.Stop()
}

func (s *JobScheduler) schedule(j config.Job) {
	s.Cron.AddFunc(j.When, func() {
		s.run(j)
	})
}

func (s *JobScheduler) run(j config.Job) {
	switch j.Type {
	case "message":
		s.runFlashMessage(j)
	case "tab":
		s.runTab(j)
	case "command":
		s.runSystemCommand(j)
	}
}
