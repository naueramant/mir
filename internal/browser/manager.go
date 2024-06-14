package browser

import (
	"strconv"
	"time"

	"github.com/naueramant/mir/internal/config"
	"github.com/naueramant/mir/internal/server"
	"github.com/naueramant/mir/internal/utils"
	"github.com/sirupsen/logrus"
)

type BrowserManager struct {
	Browser Browser
	Config  config.Configuration
}

func NewBrowserManager(c config.Configuration) BrowserManager {
	bm := BrowserManager{
		Config: c,
	}

	logrus.Infoln("Spawning chrome browser")

	bm.Browser = NewBrowser()

	if len(c.Tabs) == 0 {
		return bm
	}

	for _, tabCon := range c.Tabs {
		tab := bm.Browser.NewTab()

		if tabCon.Auth.Username != "" && tabCon.Auth.Password != "" {
			tab.NavigateWithBasicAuth(
				tabCon.URL,
				BasicAuthCredentials{
					Username: tabCon.Auth.Username,
					Password: tabCon.Auth.Password,
				},
			)
		} else {
			tab.Navigate(tabCon.URL)
		}

		bm.applyTabExtras(tab, tabCon)
	}

	bm.Browser.Tabs[0].Focus()

	logrus.Infof("Opened %d tab(s)", len(c.Tabs))

	return bm
}

func (bm *BrowserManager) Start() {
	if bm.Config.Syntax == "" {
		bm.showNoConfigScreen()
		logrus.Infoln("No configuration file found")

		return
	}

	if len(bm.Config.Tabs) == 0 {
		bm.showNoTabsScreen()
		logrus.Infoln("No tabs configured")

		return
	}

	bm.startCycle()
}

func (bm *BrowserManager) Close() {
	logrus.Infoln("Closing browser")
	bm.Browser.Close()
}

func (bm *BrowserManager) Pause() {
	// TODO: implement
}

func (bm *BrowserManager) Resume() {
	// TODO: implement
}

func (bm *BrowserManager) applyTabExtras(t *Tab, tc config.Tab) {
	if tc.CSS != "" {
		cssStr, err := utils.ReadFileToString(tc.CSS)
		if err != nil {
			logrus.Errorln("Failed to read CSS file", err)
		} else {
			go t.SetCSS(cssStr)
		}
	}

	if tc.JS != "" {
		jsStr, err := utils.ReadFileToString(tc.JS)
		if err != nil {
			logrus.Errorln("Failed to read JS file", err)
		} else {
			go t.SetJS(jsStr)
		}

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

func (bm *BrowserManager) startCycle() {
	for {
		for i, tab := range bm.Browser.Tabs {
			if bm.Config.Tabs[i].Reload {
				tab.Reload()

				bm.applyTabExtras(tab, bm.Config.Tabs[i])
			}

			tab.Focus()

			if bm.Config.Tabs[i].Duration == 0 {
				return
			}

			delay := time.Duration(bm.Config.Tabs[i].Duration)
			time.Sleep(time.Second * delay)
		}
	}
}
