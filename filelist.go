package main

import (
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
	r.item.Icon.Move(fyne.NewPos(0, 0))

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

			// Create Icon for Files
			item.Icon = canvas.NewImageFromResource(theme.DocumentIcon())
			item.Icon.SetMinSize(fyne.NewSize(20, 20))
			item.Icon.Show()

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

func CreateDirList(dirs *[]string, files *[]string, selectedFile *string, fileList *widget.List, window fyne.Window) *widget.List {
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
			item.onDoubleClick = func() {
				*selectedFile = dirName

				//TODO: Change the Directory
				// err := os.Chdir("/" + dirName)
				// if err != nil {
				// 	dialog.ShowError(err, window)
				// }

				// Or Call RefreshFileList
				RefreshFileList(files, dirName, fileList)
			}
		},
	)
}
