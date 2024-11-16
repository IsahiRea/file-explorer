package main

import (
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type CustomListItem struct {
	widget.Label
	onTapped      func()
	onDoubleClick func()
	lastTap       time.Time
}

func (i *CustomListItem) Tapped(_ *fyne.PointEvent) {
	now := time.Now()

	if now.Sub(i.lastTap) < 300*time.Millisecond {
		i.onDoubleClick()
	} else {
		i.lastTap = now
		i.onTapped()
	}
}

func NewCustomListItem(text string, onTapped func(), onDoubleClick func()) *CustomListItem {
	item := &CustomListItem{
		Label:         *widget.NewLabel(text),
		onTapped:      onTapped,
		onDoubleClick: onDoubleClick,
	}

	item.ExtendBaseWidget(item)
	return item
}

func CreateFileList(files *[]string, selectedFile *string, renameButton, deleteButton *widget.Button, window fyne.Window) *widget.List {
	return widget.NewList(
		func() int {
			return len(*files)
		},
		func() fyne.CanvasObject {
			return NewCustomListItem("", func() {}, func() {})
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			fileName := (*files)[id]
			item := co.(*CustomListItem)

			item.SetText(fileName)
			item.onDoubleClick = func() {
				*selectedFile = fileName

				err := openFile(*selectedFile)
				if err != nil {
					dialog.ShowError(err, window)
				}

			}

			item.onTapped = func() {
				*selectedFile = fileName
				renameButton.Enable()
				deleteButton.Enable()
			}
		},
	)
}

func CreateDirList(dirs *[]string, selectedFile *string, window fyne.Window) *widget.List {
	return widget.NewList(
		func() int {
			return len(*dirs)
		},
		func() fyne.CanvasObject {
			return NewCustomListItem("", func() {}, func() {})
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			dirName := (*dirs)[id]
			item := co.(*CustomListItem)

			item.SetText(dirName)
			item.onDoubleClick = func() {
				*selectedFile = dirName

				//TODO: Change the Directory
				err := os.Chdir(dirName)
				if err != nil {
					dialog.ShowError(err, window)
				}
			}
		},
	)
}
