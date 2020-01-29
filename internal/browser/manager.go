package browser

import "github.com/naueramant/mir/internal/config"

import "time"

type BrowserManager struct {
	Browser Browser
}

func (bm *BrowserManager) Start(c *config.Configuration) {
	bm.Browser = newBrowser()

	if c == nil {
		bm.showNoValidConfigurationScreen()
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
	bm.Browser.Close()
}

func (bm *BrowserManager) showNoValidConfigurationScreen() {

}

func (bm *BrowserManager) showNoTabsScreen() {

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
