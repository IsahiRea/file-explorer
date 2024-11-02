package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	createButton := widget.NewButton("Create File", func() {

		fileNameEntry := widget.NewEntry()
		fileNameEntry.SetPlaceHolder("Enter file name")

		form := dialog.NewForm("Create File", "Create", "Cancel", []*widget.FormItem{widget.NewFormItem("File Name", fileNameEntry)},
			func(confirmed bool) {
				if confirmed {
					fileName := fileNameEntry.Text

					file, err := os.Create(filepath.Join(dir, fileName))
					if err != nil {
						dialog.ShowError(err, window)
					}

					file.Close()
					fileList.Refresh()

					fmt.Printf("File name entered: %v\n", fileName)
				}
			}, window)

		form.Show()
	})

	controls := container.NewHBox(createButton, refreshButton)
	window.SetContent(container.NewBorder(controls, nil, nil, nil, fileList))
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
