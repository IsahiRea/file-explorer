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

	currentDir := "./"

	var selectedFile string
	var fileList *widget.List
	var dirList *widget.List

	dirs := GetDirs(currentDir)
	files := GetFiles(currentDir)

	// Define Buttons
	createButton := widget.NewButton("Create File", func() { CreateFileDialog(window, currentDir, func() { RefreshFileList(&files, currentDir, fileList) }) })
	renameButton := widget.NewButton("Rename File", func() {
		RenameFileDialog(window, currentDir, &selectedFile, func() { RefreshFileList(&files, currentDir, fileList) })
	})
	deleteButton := widget.NewButton("Delete File", func() {
		DeleteFileDialog(window, currentDir, &selectedFile, func() { RefreshFileList(&files, currentDir, fileList) })
	})

	renameButton.Disable()
	deleteButton.Disable()

	// File list with selection handling
	fileList = CreateFileList(&files, &selectedFile, renameButton, deleteButton, window)

	// TODO: Directory navigation
	dirList = CreateDirList(&dirs, &selectedFile, window)

	// Layout
	controls := container.NewHBox(createButton, renameButton, deleteButton)

	window.SetContent(container.NewBorder(dirList, controls, nil, nil, fileList))
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
