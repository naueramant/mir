package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	Context *context.Context
}

func NewBrowser() *Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("disable-infobars", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	return &Browser{
		Context: &allocCtx,
	}
}

func (b *Browser) NewTab(c TabConfig) *Tab {
	ctx, _ := chromedp.NewContext(*b.Context, chromedp.WithLogf(log.Printf))

	chromedp.Run(ctx, chromedp.Navigate("www.google.com"))

	return &Tab{
		Context: &ctx,
	}
}
