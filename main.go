package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
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

	if err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.xkcd.com/`),
	); err != nil {
		log.Fatal(err)
	}

	time.Sleep(100 * time.Second)
}
