package tui

import (
	"github.com/rivo/tview"
)

func HelloWorld() {
	app := tview.NewApplication()

	// Create a TextView
	textView := tview.NewTextView().
		SetText("hello world!").
		SetTextAlign(tview.AlignCenter). // Center horizontally
		SetDynamicColors(true).
		SetWrap(true)

	// Create a Flex layout to center vertically and horizontally
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).   // Stack items vertically
		AddItem(nil, 0, 1, false).     // Top spacer
		AddItem(textView, 0, 1, true). // Centered TextView
		AddItem(nil, 0, 1, false)      // Bottom spacer

	// Run the application
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
