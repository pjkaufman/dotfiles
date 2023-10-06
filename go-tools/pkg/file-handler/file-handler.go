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

type FileManager interface {
	FileExists(path string) bool
	FolderExists(path string) bool
	GetFoldersInCurrentFolder(path string) []string
	WriteFileContents(path, content string)
	ReadInFileContents(path string) string
	MustGetAllFilesWithExtInASpecificFolder(dir, ext string) []string
	GetFileFolder(filePath string) string
	JoinPath(elements ...string) string
}

type FileHandler struct {
	logger logger.Logger
}

func NewFileHandler(logger logger.Logger) *FileHandler {
	return &FileHandler{
		logger: logger,
	}
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

		fh.logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists: %s`, path, err))
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

		fh.logger.WriteError(fmt.Sprintf(`could not verify that "%s" exists and is a directory: %s`, path, err))
	}

	if !folderInfo.IsDir() {
		return false
	}

	return true
}

func (fh *FileHandler) GetFoldersInCurrentFolder(path string) []string {
	if strings.Trim(path, " ") == "" {
		return nil
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		fh.logger.WriteError(fmt.Sprintf(`could not get files/folders in "%s": %s`, path, err))
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

func (fh *FileHandler) GetFileFolder(filePath string) string {
	if strings.Trim(filePath, " ") == "" {
		return ""
	}

	return path.Join(filePath, "..")
}

func (fh *FileHandler) JoinPath(elements ...string) string {
	return path.Join(elements...)
}

func (fh *FileHandler) ReadInFileContents(path string) string {
	if strings.Trim(path, " ") == "" {
		return ""
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fh.logger.WriteError(fmt.Sprintf(`could not read in file contents for "%s": %s`, path, err))
	}

	return string(fileBytes)
}

func (fh *FileHandler) WriteFileContents(path, content string) {
	if strings.Trim(path, " ") == "" {
		return
	}

	var fileMode fs.FileMode

	fileInfo, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fileMode = fs.ModePerm
		} else {
			fh.logger.WriteError(fmt.Sprintf(`could not read in existing file info to retain existing permission for "%s": %s`, path, err))
		}
	} else {
		fileMode = fileInfo.Mode()
	}

	err = os.WriteFile(path, []byte(content), fileMode)
	if err != nil {
		fh.logger.WriteError(fmt.Sprintf(`could not write to file "%s": %s`, path, err))
	}
}

func (fh *FileHandler) MustGetAllFilesWithExtInASpecificFolder(dir, ext string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		fh.logger.WriteError(fmt.Sprintf(`failed to read in folder "%s": %s`, dir, err))
	}

	var fileList []string
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ext) {
			fileList = append(fileList, f.Name())
		}
	}

	return fileList
}
