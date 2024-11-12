package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

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

func openFile(fileName string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", "start", fileName)
	case "darwin":
		cmd = exec.Command("open", fileName)
	case "linux":
		cmd = exec.Command("xdg-open", fileName)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
