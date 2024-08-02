package main

import (
	"fmt"

	handlers "GpioController/internal/handlers/prompt"
	"GpioController/internal/ui"
	"GpioController/types"

	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println(`
 _____       _       _____             _             _ _
|  __ \     (_)     /  __ \           | |           | | |
| |  \/_ __  _  ___ | /  \/ ___  _ __ | |_ _ __ ___ | | | ___ _ __
| | __| '_ \| |/ _ \| |    / _ \| '_ \| __| '__/ _ \| | |/ _ \ '__|
| |_\ \ |_) | | (_) | \__/\ (_) | | | | |_| | | (_) | | |  __/ |
 \____/ .__/|_|\___/ \____/\___/|_| |_|\__|_|  \___/|_|_|\___|_|
      | |
      |_|`)
	fmt.Println("")

	c := types.Controller{}

	err := handlers.GetUserInput(&c)
	if err != nil {
		fmt.Println("Error:", err)
	}

	a := app.New()

	mainWindow := ui.NewMainWindowGui(&c, a)

	mainWindow.ShowAndRun()
}
