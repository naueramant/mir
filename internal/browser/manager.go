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
}

func (bm *BrowserManager) Start(c *config.Configuration) {
	bm.Browser = newBrowser()

	if c == nil || len(c.Tabs) == 0 {
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
	bm.Browser.Close()
}

func (bm *BrowserManager) showNoTabsScreen() {
	t := bm.Browser.NewTab()
	t.Navigate("localhost:" + strconv.Itoa(server.Port) + "/notabs.html?ip=" + utils.GetLocalIp())
}

func (bm *BrowserManager) startCycle(c *config.Configuration) {
	for {
		for i, tab := range bm.Browser.Tabs {
			tabCon := c.Tabs[i]

			if tabCon.Reload {
				tab.Reload()
			}

			tab.Focus()

			delay := time.Duration(tabCon.Duration)
			time.Sleep(time.Second * delay)
		}
	}
}
