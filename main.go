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

	currentDir, _ := os.Getwd()
	println(currentDir)

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
	fileList = CreateFileList(&currentDir, &files, &selectedFile, renameButton, deleteButton, window)
	dirList = CreateDirList(&currentDir, &dirs, &files, window)

	// Go routine to refresh the file list and directory list
	go func() {
		for {
			RefreshFileList(&files, currentDir, fileList)
			RefreshDirList(&dirs, currentDir, dirList)
		}
	}()

	// Layout
	path := container.NewHSplit(dirList, fileList)
	controls := container.NewHBox(createButton, renameButton, deleteButton)

	layout := container.NewBorder(nil, controls, nil, nil, path)

	window.SetContent(layout)
	window.Resize(fyne.NewSize(600, 400))
	window.ShowAndRun()
}
