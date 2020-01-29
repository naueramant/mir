package browser

import (
	"context"
	"encoding/base64"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type Tab struct {
	Browser Browser
	Context context.Context

	Close func()
}

type BasicAuthCredentials struct {
	Username string
	Password string
}

func (t *Tab) Focus() {
	chromedp.Run(t.Context, page.BringToFront())
}

func (t *Tab) Reload() {
	chromedp.Run(t.Context, chromedp.Reload())
}

func (t *Tab) Navigate(url string) {
	chromedp.Run(t.Context, chromedp.Navigate(url))
}

func (t *Tab) NavigateWithBasicAuth(url string, creds BasicAuthCredentials) {
	x := base64.StdEncoding.EncodeToString([]byte(creds.Username + ":" + creds.Password))
	headers := network.Headers{"Authorization": "Basic " + x}

	tasks := chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(headers)),
		chromedp.Navigate(url),
	}

	chromedp.Run(t.Context, tasks)
}
