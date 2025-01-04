package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomListItem struct {
	widget.Label
	Icon          *canvas.Image
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

func (i *CustomListItem) CreateRenderer() fyne.WidgetRenderer {
	i.ExtendBaseWidget(i)
	return &CustomListItemRenderer{
		item:    i,
		objects: []fyne.CanvasObject{i.Icon, &i.Label},
	}
}

func NewCustomListItem(text string, icon fyne.Resource, onTapped func(), onDoubleClick func()) *CustomListItem {
	iconImg := canvas.NewImageFromResource(icon)
	iconImg.SetMinSize(fyne.NewSize(20, 20))

	item := &CustomListItem{
		Label:         *widget.NewLabel(text),
		Icon:          iconImg,
		onTapped:      onTapped,
		onDoubleClick: onDoubleClick,
	}

	item.ExtendBaseWidget(item)
	return item
}

type CustomListItemRenderer struct {
	item    *CustomListItem
	objects []fyne.CanvasObject
}

func (r *CustomListItemRenderer) Layout(size fyne.Size) {
	r.item.Icon.Resize(fyne.NewSize(20, 20))
	r.item.Icon.Move(fyne.NewPos(0, (size.Height-20)/2))

	r.item.Label.Resize(fyne.NewSize(size.Width-25, size.Height))
	r.item.Label.Move(fyne.NewPos(25, 0))
}

func (r *CustomListItemRenderer) MinSize() fyne.Size {
	return fyne.NewSize(100, 20) // Example minimum size
}

func (r *CustomListItemRenderer) Refresh() {
	canvas.Refresh(r.item.Icon)
	r.item.Label.Refresh()
}

func (r *CustomListItemRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *CustomListItemRenderer) Destroy() {}

// ----------------------------------------------------------------------------------
func CreateFileList(files *[]string, selectedFile *string, renameButton, deleteButton *widget.Button, window fyne.Window) *widget.List {
	return widget.NewList(
		func() int {
			return len(*files)
		},
		func() fyne.CanvasObject {
			return NewCustomListItem("", nil, func() {}, func() {})
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			fileName := (*files)[id]
			item := co.(*CustomListItem)

			item.SetText(fileName)

			// Create Icon
			item.Icon.Resource = theme.DocumentIcon()

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

func CreateDirList(currentDir string, dirs *[]string, files *[]string, fileList *widget.List, window fyne.Window) *widget.List {
	return widget.NewList(
		func() int {
			return len(*dirs)
		},
		func() fyne.CanvasObject {
			return NewCustomListItem("", nil, func() {}, func() {})
		},
		func(id widget.ListItemID, co fyne.CanvasObject) {
			dirName := (*dirs)[id]
			item := co.(*CustomListItem)

			item.SetText(dirName)

			// Create Icon
			item.Icon.Resource = theme.FolderIcon()

			item.onTapped = func() {
				//TODO Grab the valid directory path
				/*
					Change the currentDir Var (Done)
					Try passing in the var to this function
					Change the var to the proper path
					Update the fileList
				*/

				fmt.Println(dirName)
				currentDir = dirName
				path := compareDir(currentDir, dirName)
				//os.Chdir(dirName)
				RefreshFileList(files, path, fileList)
			}
		},
	)
}
