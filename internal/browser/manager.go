package browser

import (
	"strconv"
	"time"

	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/server"
	"github.com/naueramant/mir/internal/utils"
)

type BrowserManager struct {
	Browser Browser
	Started bool
	Paused  bool
}

func (bm *BrowserManager) Start(c *config.Configuration) {
	if bm.Started {
		return
	}

	bm.Started = true

	bm.Browser = newBrowser()

	if c == nil {
		bm.showNoConfigScreen()
		return
	}

	if len(c.Tabs) == 0 {
		bm.showNoTabsScreen()
		return
	}

	for _, tabCon := range c.Tabs {
		tab := bm.Browser.NewTab()

		if tabCon.Auth.Username != "" && tabCon.Auth.Password != "" {
			tab.NavigateWithBasicAuth(tabCon.URL, BasicAuthCredentials{
				Username: tabCon.Auth.Username,
				Password: tabCon.Auth.Password,
			})
		} else {
			tab.Navigate(tabCon.URL)
		}
	}

	bm.startCycle(c)
}

func (bm *BrowserManager) Stop() {
	if bm.Started {
		bm.Browser.Close()
	}
}

func (bm *BrowserManager) Pause() {
	bm.Paused = true
}

func (bm *BrowserManager) Resume() {
	bm.Paused = false
}

func ApplyTabExtras(t Tab, tc config.Tab) {
	if tc.CSS != "" {
		cssStr, _ := utils.ReadFileToString(tc.CSS)
		go t.AddCSS(cssStr)
	}

	if tc.JS != "" {
		jsStr, _ := utils.ReadFileToString(tc.JS)
		go t.AddJS(jsStr)
	}
}

func (bm *BrowserManager) showNoTabsScreen() {
	t := bm.Browser.NewTab()
	t.Navigate("localhost:" + strconv.Itoa(server.Port) + "/notabs.html?ip=" + utils.GetLocalIp())
}

func (bm *BrowserManager) showNoConfigScreen() {
	t := bm.Browser.NewTab()
	t.Navigate("localhost:" + strconv.Itoa(server.Port) + "/noconfig.html?ip=" + utils.GetLocalIp())
}

func (bm *BrowserManager) startCycle(c *config.Configuration) {
	for {
		for i, tab := range bm.Browser.Tabs {
			if c.Tabs[i].Reload {
				tab.Reload()
			}

			ApplyTabExtras(tab, c.Tabs[i])

			tab.Focus()

			delay := time.Duration(c.Tabs[i].Duration)
			time.Sleep(time.Second * delay)
		}
	}
}
