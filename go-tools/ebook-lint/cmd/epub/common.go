package epub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/ebook-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

const (
	EpubPathArgEmpty   = "epub-file must have a non-whitespace value"
	EpubPathArgNonEpub = "epub-file must be an Epub file"
	cliLineSeparator   = "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"
)

var epubFile string

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

func validateCommonEpubFlags(epubPath string) error {
	if strings.TrimSpace(epubPath) == "" {
		return errors.New(EpubPathArgEmpty)
	}

	if !strings.HasSuffix(epubPath, ".epub") {
		return errors.New(EpubPathArgNonEpub)
	}

	return nil
}

// func compressImages(destFolder, opfFolder string, images map[string]struct{}) {
// 	for imagePath := range images {
// 		if !isCompressableImage(imagePath) {
// 			continue
// 		}

// 		var params = fmt.Sprintf("%s %s %s", imgComperssionProgramName, strings.Join(compressionParams, " "), filehandler.JoinPath(opfFolder, imagePath))
// 		commandhandler.MustRunCommand("bash", fmt.Sprintf(`failed to compress "%s"`, imagePath), []string{"-c", params}...)

// 		// TODO: see if I can figure out why the following does not work
// 		// var params = append(compressionParams, "\""+filehandler.JoinPath(opfFolder, imagePath)+"\"")
// 		// fmt.Println(commandhandler.MustGetCommandOutput(imgComperssionProgramName, fmt.Sprintf(`failed to compress "%s"`, imagePath), params...))
// 	}
// }

// func isCompressableImage(imagePath string) bool {
// 	for _, ext := range compressableImageExts {
// 		if strings.HasSuffix(strings.ToLower(imagePath), ext) {
// 			return true
// 		}
// 	}

// 	return false
// }
