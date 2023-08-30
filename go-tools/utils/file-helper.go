package utils

import (
	"os"
	"strings"
)

func FileExists(path string) (bool, error) {
	if strings.Trim(path, " ") == "" {
		return false, nil
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func FolderExists(path string) (bool, error) {
	if strings.Trim(path, " ") == "" {
		return false, nil
	}

	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if !folderInfo.IsDir() {
		return false, nil
	}

	return true, nil
}
