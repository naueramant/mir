package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	Context context.Context
	Tabs    []Tab

	Close func()
}

func newBrowser() Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("disable-infobars", true),
	)

	allocCtx, close := chromedp.NewExecAllocator(context.Background(), opts...)

	return Browser{
		Context: allocCtx,
		Close:   close,
	}
}

func (b *Browser) NewTab() Tab {
	ctx, close := chromedp.NewContext(b.Context, chromedp.WithLogf(log.Printf))

	t := Tab{
		Browser: *b,
		Context: ctx,
	}

	t.Close = func() {
		b.removeTab(t)
		close()
	}

	b.Tabs = append(b.Tabs, t)

	return t
}

func (b *Browser) removeTab(t Tab) {
	for i, tab := range b.Tabs {
		if tab.Context == t.Context {
			b.Tabs = append(b.Tabs[:i], b.Tabs[i+1:]...)
			return
		}
	}
}
