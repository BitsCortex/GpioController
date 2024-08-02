package ui

import (
	"GpioController/types"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindowGui(c *types.Controller, app fyne.App) fyne.Window {
	w := app.NewWindow("GpioControler GUI")
	w.Resize(fyne.NewSize(800, 480))

	str := fmt.Sprintf("%d", c.Pin)

	label := widget.NewLabel("Click to show Gpio pin!")

	button := widget.NewButton("Click Me", func() {
		label.SetText(str)
	})

	wc := container.NewVBox(label, button)

	w.SetContent(wc)

	return w
}
