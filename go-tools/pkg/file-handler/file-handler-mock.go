//go:build unit

package filehandler

import (
	"fmt"
	"path"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

type FileHandlerMock struct {
	logger          logger.Logger
	existingFiles   map[string]struct{}
	existingFolders map[string]struct{}
	pathToFolders   map[string][]string
	fileContents    map[string]string
	fileWriteErr    error
}

func NewMockFileHandler(logger logger.Logger, existingFiles, existingFolders map[string]struct{}, pathToFolders map[string][]string, fileContents map[string]string) *FileHandlerMock {
	return &FileHandlerMock{
		logger:          logger,
		existingFiles:   existingFiles,
		existingFolders: existingFolders,
		pathToFolders:   pathToFolders,
		fileContents:    fileContents,
	}
}

func (fh *FileHandlerMock) FileExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	if _, ok := fh.existingFiles[path]; ok {
		return true
	}

	return false
}

func (fh *FileHandlerMock) FolderExists(path string) bool {
	if strings.Trim(path, " ") == "" {
		return false
	}

	if _, ok := fh.existingFolders[path]; ok {
		return true
	}

	return false
}

func (fh *FileHandlerMock) GetFoldersInCurrentFolder(path string) []string {
	if strings.Trim(path, " ") == "" {
		return nil
	}

	if folders, ok := fh.pathToFolders[path]; ok {
		return folders
	}

	fh.logger.WriteError(fmt.Sprintf(`could not get files/folders in "%s": %s`, path, "path not found"))

	return nil
}

func (fh *FileHandlerMock) ReadInFileContents(path string) string {
	if strings.Trim(path, " ") == "" {
		return ""
	}

	if fileContent, ok := fh.fileContents[path]; ok {
		return fileContent
	}

	fh.logger.WriteError(fmt.Sprintf(`could not read in file contents for "%s": %s`, path, "path not found"))

	return ""
}

func (fh *FileHandlerMock) WriteFileContents(path, content string) {
	if strings.Trim(path, " ") == "" {
		return
	}

	if fh.fileWriteErr != nil {
		fh.logger.WriteError(fmt.Sprintf(`could not write to file "%s": %s`, path, fh.fileWriteErr))
	}
}

func (fh *FileHandlerMock) MustGetAllFilesWithExtInASpecificFolder(dir, ext string) []string {
	var fileList []string
	for _, f := range fh.pathToFolders[dir] {
		if strings.HasSuffix(f, ext) {
			fileList = append(fileList, f)
		}
	}

	return fileList
}

func (fh *FileHandlerMock) GetFileFolder(filePath string) string {
	if strings.Trim(filePath, " ") == "" {
		return ""
	}

	return path.Join(filePath, "..")
}

func (fh *FileHandlerMock) JoinPath(elements ...string) string {
	return path.Join(elements...)
}
