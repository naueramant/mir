package browser

var (
	browser Browser
)

func StartManager() {
	browser = NewBrowser()

	t := browser.NewTab()

	t.Navigate("https://google.com")

	for true {

	}
}
