package jobs

import (
	"github.com/naueramant/mir/internal/browser"
	"github.com/naueramant/mir/internal/config"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
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
	for i, j := range s.Config.Jobs {
		s.Cron.AddFunc(j.When, func() {
			s.run(i, j)
		})
	}

	logrus.Infof("Scheduled %d job(s)", len(s.Config.Jobs))

	s.Cron.Start()
}

func (s *JobScheduler) Stop() {
	logrus.Infoln("Stopping job schedular")
	s.Cron.Stop()
}

func (s *JobScheduler) run(i int, j config.Job) {
	logrus.Infof("Executing job %d", i)

	switch j.Type {
	case "message":
		s.runFlashMessage(j)
	case "tab":
		s.runTab(j)
	case "command":
		s.runSystemCommand(j)
	}
}
