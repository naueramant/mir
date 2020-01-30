package jobs

import (
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/server"
)

func (s *JobScheduler) runFlashMessage(j config.Job) {
	s.BrowserManager.Pause()
	tab := s.BrowserManager.Browser.NewTab()

	u, _ := url.Parse("localhost:" + strconv.Itoa(server.Port) + "/message.html")
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("msg", j.Data.Message)
	if j.Data.FontSize != 0 {
		q.Add("fontsize", strconv.Itoa(int(j.Data.FontSize)))
	}
	if j.Data.TextColor != "" {
		q.Add("textcolor", j.Data.TextColor)
	}
	if j.Data.BackgroundColor != "" {
		q.Add("bgcolor", j.Data.BackgroundColor)
	}
	q.Add("blink", strconv.FormatBool(j.Data.Blink))

	u.RawQuery = q.Encode()

	tab.Navigate(u.String())

	time.Sleep(time.Second * time.Duration(int(j.Data.Duration)))

	tab.Close()
	s.BrowserManager.Resume()
}

func (s *JobScheduler) runTab(j config.Job) {
	s.BrowserManager.Pause()
	tab := s.BrowserManager.Browser.NewTab()

	tab.Navigate(j.Data.URL)

	time.Sleep(time.Second * time.Duration(int(j.Data.Duration)))

	tab.Close()
	s.BrowserManager.Resume()
}

func (s *JobScheduler) runSystemCommand(j config.Job) {
	exec.Command(j.Data.Command, j.Data.Args...).Run()
}
