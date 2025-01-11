package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"fyne.io/fyne/v2/widget"
)

func GetFiles(currentDir string) []string {
	files, err := os.ReadDir(currentDir)
	if err != nil {
		println(err)
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

func RefreshDirList(dirs *[]string, currentDir string, dirList *widget.List) {
	*dirs = GetDirs(currentDir)
	dirList.Refresh()
}

func isDir(currentDir, fileName string) bool {

	fileName = filepath.Join(currentDir, fileName)

	fileInfo, err := os.Stat(fileName)
	if err != nil {
		println(err)
		return false
	}

	return fileInfo.IsDir()
}

func addDir(currentDir string, dirName string) string {
	return filepath.Join(currentDir, dirName)
}

func findDir(currentDir string, dirName string) string {

	// Split the current directory path
	parts := strings.Split(filepath.ToSlash(currentDir), "/")

	for i, part := range parts {
		if part == dirName {
			parts = parts[:i+1]
			break
		}
	}

	// Join the parts to form the new path
	newPath := filepath.Join(parts...)
	newPath = "/" + newPath

	println(newPath)

	return newPath
}

func GetDirs(currentDir string) []string {

	currentDir = strings.Trim(currentDir, "/")
	dirNames := strings.Split(filepath.ToSlash(currentDir), "/")

	return dirNames
}

func isWSL() bool {
	b, err := exec.Command("uname", "-r").Output()
	if err != nil {
		println(err)
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
