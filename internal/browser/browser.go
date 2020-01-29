package browser

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	Context context.Context
}

func NewBrowser() *Browser {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("start-fullscreen", true),
		chromedp.Flag("disable-infobars", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	// create context
	ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))

	return &Browser{
		Context: ctx,
	}
}
