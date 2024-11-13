package main

import (
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

				//FIXME Figure out how to open a file
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
