package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Hola Fyne")

	myWindow.SetContent(widget.NewLabel("Hello world"))
	myWindow.ShowAndRun()
}
