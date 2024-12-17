package tui

import (
	"github.com/rivo/tview"
)

func HelloWorld() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	
	app := tview.NewApplication()

	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}