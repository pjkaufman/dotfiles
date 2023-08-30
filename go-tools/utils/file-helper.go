package utils

import (
	"fmt"
	"os"
	"strings"
)

func FileExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		WriteError(fmt.Sprintf(`could not verify that "%s" exists: %s`, path, err))
	}

	return true
}

func FolderExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		WriteError(fmt.Sprintf(`could not verify that "%s" exists and is a directory: %s`, path, err))
	}

	if !folderInfo.IsDir() {
		return false
	}

	return true
}

func GetFoldersInCurrentFolder(path string) []string {
	if strings.Trim(path, " ") == "" {
		return nil
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		WriteError(fmt.Sprintf(`could not get files/folders in "%s": %s`, path, err))
	}

	var actualDirs []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		actualDirs = append(actualDirs, dir.Name())
	}

	return actualDirs
}
