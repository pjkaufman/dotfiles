package filehandler

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
)

type FileManager interface {
	FileExists(path string) bool
	FolderExists(path string) bool
}

type FileHandler struct {
}

func (fh *FileHandler) FileExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		utils.WriteError(fmt.Sprintf(`could not verify that "%s" exists: %s`, path, err))
	}

	return true
}

func (fh *FileHandler) FolderExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}

		utils.WriteError(fmt.Sprintf(`could not verify that "%s" exists and is a directory: %s`, path, err))
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
		utils.WriteError(fmt.Sprintf(`could not get files/folders in "%s": %s`, path, err))
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
	if strings.Trim(filePath, " ") == "" {
		return ""
	}

	return path.Join(filePath, "..")
}

func JoinPath(elements ...string) string {
	return path.Join(elements...)
}

func ReadInFileContents(path string) string {
	if strings.Trim(path, " ") == "" {
		return ""
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		utils.WriteError(fmt.Sprintf(`could not read in file contents for "%s": %s`, path, err))
	}

	return string(fileBytes)
}

func WriteFileContents(path, content string) {
	if strings.Trim(path, " ") == "" {
		return
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		utils.WriteError(fmt.Sprintf(`could not read in existing file info to retain existing permission for "%s": %s`, path, err))
	}

	err = os.WriteFile(path, []byte(content), fileInfo.Mode())
	if err != nil {
		utils.WriteError(fmt.Sprintf(`could not write to file "%s": %s`, path, err))
	}
}

func MustGetAllFilesWithExtInASpecificFolder(dir, ext string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		utils.WriteError(fmt.Sprintf(`failed to read in folder "%s": %s`, dir, err))
	}

	var fileList []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			fileList = append(fileList, f.Name())
		}
	}

	return fileList
}
