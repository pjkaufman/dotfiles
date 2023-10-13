package epub

import (
	"fmt"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func getEpubInfo(dir, epubName string) (string, linter.EpubInfo) {
	opfFiles := filehandler.MustGetAllFilesWithExtInASpecificFolderAndSubFolders(dir, ".opf")
	if len(opfFiles) < 1 {
		logger.WriteError(fmt.Sprintf("did not find opf file for \"%s\"", epubName))
	}

	var opfFile = opfFiles[0]
	opfText := filehandler.ReadInFileContents(opfFile)

	epubInfo, err := linter.ParseOpfFile(opfText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("Failed to parse \"%s\" for \"%s\": %s", opfFile, epubName, err))
	}

	var opfFolder = filehandler.GetFileFolder(opfFile)

	return opfFolder, epubInfo
}

func validateFilesExist(opfFolder string, files map[string]struct{}) {
	for file := range files {
		var filePath = getFilePath(opfFolder, file)

		if !filehandler.FileExists(filePath) {
			logger.WriteError(fmt.Sprintf(`file from manifest not found: "%s" must exist`, filePath))
		}
	}
}
