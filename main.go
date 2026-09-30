package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var version = "dev"

func main() {
	a := app.New()

	w := a.NewWindow("Hello GUI " + version)

	output := widget.NewLabel("Нажмите кнопку ниже")

	greetButton := widget.NewButton("Поздороваться", func() {
		output.SetText(Greeting("GitHub"))
	})

	quitButton := widget.NewButton("Выход", func() {
		a.Quit()
	})

	content := container.NewVBox(
		widget.NewLabel("Hello from Go GUI!"),
		widget.NewSeparator(),
		greetButton,
		output,
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("Version: %s", version)),
		widget.NewLabel(fmt.Sprintf("Sum 1..10 = %d", SumRange(1, 10))),
		widget.NewSeparator(),
		quitButton,
	)

	w.SetContent(content)
	w.Resize(content.MinSize())
	w.ShowAndRun()
}