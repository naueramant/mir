package jobs

import (
	"fmt"
	"net/url"
	"os/exec"
	"strconv"
	"time"

	"github.com/naueramant/mir/internal/config"
)

func (s *JobScheduler) runFlashMessage(j config.Job) {
	s.BrowserManager.Pause()
	tab := s.BrowserManager.Browser.NewTab()

	u, _ := url.Parse(fmt.Sprintf(
		"%s/static/message.html",
		s.AssetsServer.Host(),
	))
	q, _ := url.ParseQuery(u.RawQuery)

	q.Add("msg", j.Options.Message)
	if j.Options.FontSize != 0 {
		q.Add("fontsize", strconv.Itoa(int(j.Options.FontSize)))
	}
	if j.Options.TextColor != "" {
		q.Add("textcolor", j.Options.TextColor)
	}
	if j.Options.BackgroundColor != "" {
		q.Add("bgcolor", j.Options.BackgroundColor)
	}
	q.Add("blink", strconv.FormatBool(j.Options.Blink))

	u.RawQuery = q.Encode()

	tab.Navigate(u.String())

	time.Sleep(time.Second * time.Duration(int(j.Options.Duration)))

	tab.Close()
	s.BrowserManager.Resume()
}

func (s *JobScheduler) runTab(j config.Job) {
	s.BrowserManager.Pause()
	tab := s.BrowserManager.Browser.NewTab()

	tab.Navigate(j.Options.URL)

	time.Sleep(time.Second * time.Duration(int(j.Options.Duration)))

	tab.Close()
	s.BrowserManager.Resume()
}

func (s *JobScheduler) runSystemCommand(j config.Job) {
	exec.Command(j.Options.Command, j.Options.Args...).Run()
}
