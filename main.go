package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	os.Setenv("LANG", "en_US.UTF-8")
	os.Setenv("LC_ALL", "en_US.UTF-8")

	application := app.New()
	window := application.NewWindow("File Explorer")

	dir := "./"
	fileList := widget.NewList(
		func() int {
			files, _ := os.ReadDir(dir)
			return len(files)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("File name")
		},
		func(i widget.ListItemID, co fyne.CanvasObject) {
			files, _ := os.ReadDir(dir)
			co.(*widget.Label).SetText(files[i].Name())
		},
	)

	refreshButton := widget.NewButton("Refresh", func() {
		fileList.Refresh()
	})

	window.SetContent(container.NewBorder(refreshButton, nil, nil, nil, fileList))
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
