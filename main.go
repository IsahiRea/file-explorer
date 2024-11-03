package main

import (
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
	files, _ := os.ReadDir(dir)
	var selectedFile string

	fileList := widget.NewList(
		func() int {
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

	var createButton, deleteButton *widget.Button

	refreshTotal := func() {
		files, _ = os.ReadDir(dir)
		fileList.Refresh()
		selectedFile = ""
		deleteButton.Disable()
	}

	fileList.OnSelected = func(id widget.ListItemID) {
		selectedFile = files[id].Name()
		deleteButton.Enable()
	}

	// -------------------------------------------------------------------------------------
	refreshButton := widget.NewButton("Refresh", func() {
		refreshTotal()
	})

	createButton = widget.NewButton("Create File", func() {

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
					refreshTotal()
				}
			}, window)

		form.Show()
	})

	deleteButton = widget.NewButton("Delete File", func() {

		if selectedFile == "" {
			dialog.ShowInformation("No file selected", "Please select a file", window)

			return
		}

		form := dialog.NewForm("Delete File", "Delete", "Cancel", nil,
			func(confirmed bool) {
				if confirmed {

					if err := os.Remove(filepath.Join(dir, selectedFile)); err != nil {
						dialog.ShowError(err, window)
					}

					refreshTotal()
				}
			}, window)

		form.Show()
	})

	deleteButton.Disable()

	controls := container.NewHBox(createButton, deleteButton, refreshButton)
	window.SetContent(container.NewBorder(controls, nil, nil, nil, fileList))
	window.Resize(fyne.NewSize(400, 600))
	window.ShowAndRun()
}
