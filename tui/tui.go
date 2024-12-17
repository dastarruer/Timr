package tui

import (
	"github.com/rivo/tview"
)

func HelloWorld() {
	flex := tview.NewFlex()

	// input := tview.NewInputField()

	app := tview.NewApplication()

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}