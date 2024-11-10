package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// CreateFileDialog shows the dialog to create a new file
func CreateFileDialog(window fyne.Window, currentDir string, onSuccess func()) {
	fileNameEntry := widget.NewEntry()
	fileNameEntry.SetPlaceHolder("Enter file name")

	form := dialog.NewForm("Create New File", "Create", "Cancel",
		[]*widget.FormItem{{Text: "File Name", Widget: fileNameEntry}},
		func(confirmed bool) {
			if confirmed {
				file, err := os.Create(filepath.Join(currentDir, fileNameEntry.Text))
				if err != nil {
					dialog.ShowError(err, window)
				} else {
					file.Close()
					onSuccess()
				}
			}
		}, window)

	form.Show()
}

// RenameFileDialog shows the dialog to rename a selected file
func RenameFileDialog(window fyne.Window, currentDir string, selectedFile *string, onSuccess func()) {
	newNameEntry := widget.NewEntry()
	newNameEntry.SetPlaceHolder("Enter new name")

	form := dialog.NewForm("Rename File", "Rename", "Cancel",
		[]*widget.FormItem{{Text: "New Name", Widget: newNameEntry}},
		func(confirmed bool) {
			if confirmed {
				oldPath := filepath.Join(currentDir, *selectedFile)
				newPath := filepath.Join(currentDir, newNameEntry.Text)
				err := os.Rename(oldPath, newPath)
				if err != nil {
					dialog.ShowError(err, window)
				} else {
					onSuccess()
				}
			}
		}, window)

	form.Show()
}

// DeleteFileDialog shows the dialog to delete a selected file
func DeleteFileDialog(window fyne.Window, currentDir string, selectedFile *string, onSuccess func()) {
	form := dialog.NewForm("Delete File", "Delete", "Cancel", nil,
		func(confirmed bool) {
			if confirmed {
				err := os.Remove(filepath.Join(currentDir, *selectedFile))
				if err != nil {
					dialog.ShowError(err, window)
				} else {
					onSuccess()
				}
			}
		}, window)

	form.Show()
}
