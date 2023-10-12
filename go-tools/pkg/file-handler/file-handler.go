package filehandler

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func FileExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists: %s`, path, err))
	}

	return true
}

func FileMustExist(path, name string) {
	if strings.TrimSpace(path) == "" {
		logger.WriteError(fmt.Sprintf("%s must have a non-whitespace value", name))
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.WriteError(fmt.Sprintf("%s: \"%s\" must exist", name, path))
		}

		logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists: %s`, path, err))
	}
}

func FolderExists(path string) bool {
	if strings.TrimSpace(path) == "" {
		return false
	}

	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists and is a directory: %s`, path, err))
	}

	if !folderInfo.IsDir() {
		return false
	}

	return true
}

func FolderMustExist(path, name string) {
	if strings.TrimSpace(path) == "" {
		logger.WriteError(fmt.Sprintf("%s must have a non-whitespace value", name))
	}

	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			logger.WriteError(fmt.Sprintf("%s: \"%s\" must exist", name, path))
		}

		logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists and is a directory: %s`, path, err))
	}

	if !folderInfo.IsDir() {
		logger.WriteError(fmt.Sprintf("%s: \"%s\" must be a folder", name, path))
	}
}

func GetFoldersInCurrentFolder(path string) []string {
	if strings.TrimSpace(path) == "" {
		return nil
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`could not get files/folders in "%s": %s`, path, err))
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

func GetFileFolder(filePath string) string {
	if strings.TrimSpace(filePath) == "" {
		return ""
	}

	return path.Join(filePath, "..")
}

func JoinPath(elements ...string) string {
	return path.Join(elements...)
}

func ReadInFileContents(path string) string {
	if strings.TrimSpace(path) == "" {
		return ""
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`could not read in file contents for "%s": %s`, path, err))
	}

	return string(fileBytes)
}

func WriteFileContents(path, content string) {
	if strings.TrimSpace(path) == "" {
		return
	}

	var fileMode fs.FileMode

	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fileMode = fs.ModePerm
		} else {
			logger.WriteError(fmt.Sprintf(`could not read in existing file info to retain existing permission for "%s": %s`, path, err))
		}
	} else {
		fileMode = fileInfo.Mode()
	}

	err = os.WriteFile(path, []byte(content), fileMode)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`could not write to file "%s": %s`, path, err))
	}
}

func MustGetAllFilesWithExtInASpecificFolder(dir, ext string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		logger.WriteError(fmt.Sprintf(`failed to read in folder "%s": %s`, dir, err))
	}

	var fileList []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			fileList = append(fileList, f.Name())
		}
	}

	return fileList
}
