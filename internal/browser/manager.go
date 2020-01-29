package browser

import (
	"log"

	"github.com/chromedp/chromedp"
)

var (
	browser *Browser
)

func StartManager() {
	browser = NewBrowser()

	if err := chromedp.Run(browser.Context,
		chromedp.Navigate(`https://www.xkcd.com/`),
	); err != nil {
		log.Fatal(err)
	}

	for true {

	}
}
