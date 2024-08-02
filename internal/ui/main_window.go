package ui

import (
	handlers "GpioController/internal/handlers/prompt"
	"GpioController/types"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindowGui(c *types.Controller, app fyne.App) fyne.Window {
	w := app.NewWindow("GpioControler GUI")
	w.Resize(fyne.NewSize(800, 480))

	intro := canvas.NewText("Click Show!", color.Opaque)
	intro.TextSize = 32

	pin := fmt.Sprintf("Pin: %d", c.Pin)
	time := fmt.Sprintf("Time: %d sec", int(c.Ticker.Seconds()))

	pinW := widget.NewLabel("")
	timeW := widget.NewLabel("")

	wGrid := container.New(layout.NewGridLayout(2), pinW, timeW)

	button := widget.NewButton("Show", func() {
		pinW.SetText(pin)
		timeW.SetText(time)
	})

	newSetupButton := widget.NewButton("Setup new pin", func() {
		w.Hide()
		err := handlers.GetUserInput(c)
		if err != nil {
			fmt.Println("Error:", err)
		}

		w = NewMainWindowGui(c, app)
		w.Show()
	})

	closeButton := widget.NewButton("Close GUI", func() {
		c := widget.NewLabel("Do you want to leave?")

		pursueB := widget.NewButton("Yes", func() {
			app.Quit()
		})

		buttons := container.NewGridWithColumns(3,
			layout.NewSpacer(),
			pursueB,
			layout.NewSpacer(),
		)

		confirm := dialog.NewCustom("Alert!", "Cancel", container.NewGridWithRows(2, c, buttons), w)

		confirm.Resize(fyne.NewSize(350, 200))

		confirm.Show()
	})

	buttonGrid := container.New(layout.NewGridLayout(3), button, newSetupButton, closeButton)

	grid := container.New(layout.NewGridLayout(1), intro, wGrid, buttonGrid)
	center := container.New(layout.NewCenterLayout(), grid)

	w.SetContent(center)

	return w
}
