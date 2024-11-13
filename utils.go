package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

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

func isWSL() bool {
	b, err := exec.Command("uname", "-r").Output()
	if err != nil {
		return false
	}

	return strings.Contains(string(b), "microsoft")
}

func openFile(fileName string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", fileName)
	case "darwin":
		cmd = exec.Command("open", fileName)
	case "linux":
		// Check for WSL
		if isWSL() {
			cmd = exec.Command("explorer.exe", fileName)
		} else {
			cmd = exec.Command("xdg-open", fileName)
		}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
