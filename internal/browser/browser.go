package browser

import (
	"context"
	"os"
	"path"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/sirupsen/logrus"
)

type Browser struct {
	Context context.Context
	Tabs    []*Tab

	Close func()
}

func NewBrowser() *Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("disable-infobars", true),
		chromedp.Flag("user-data-dir", path.Join(os.TempDir(), "chromium")),
		chromedp.Flag("incognito", true),
	)

	allocCtx, close := chromedp.NewExecAllocator(context.Background(), opts...)

	return &Browser{
		Context: allocCtx,
		Close:   close,
	}
}

func (b *Browser) NewTab() *Tab {
	t := newTab(b)

	chromedp.ListenTarget(t.Context, func(ev interface{}) {
		if e, ok := ev.(*network.EventLoadingFailed); ok {
			if e.Type == network.ResourceTypeDocument {
				logrus.Infof("Tab failed to load, reloading in %v seconds\n", FailedLoadReloadDelay.Seconds())

				go t.delayedReload()
			}
		}

		if _, ok := ev.(*target.EventTargetDestroyed); ok {
			b.removeTab(t)
		}
	})

	t.Close = func() {
		b.removeTab(t)
		t.Close()
	}

	b.Tabs = append(b.Tabs, t)

	return t
}

func (b *Browser) removeTab(t *Tab) {
	for i, tab := range b.Tabs {
		if tab.Context == t.Context {
			b.Tabs = append(b.Tabs[:i], b.Tabs[i+1:]...)
			return
		}
	}
}
