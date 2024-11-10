package main

import (
	"os"

	"fyne.io/fyne/v2/widget"
)

func GetFiles(currentDir string) []string {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		return nil
	}

	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func RefreshFileList(files *[]string, currentDir string, fileList *widget.List) {
	*files = GetFiles(currentDir)
	fileList.Refresh()
}
