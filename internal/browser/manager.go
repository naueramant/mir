package browser

import "fmt"

var (
	browser *Browser
)

func StartManager() {
	browser = NewBrowser()
	tab := browser.NewTab(TabConfig{})

	fmt.Println(tab)

	for true {

	}
}
