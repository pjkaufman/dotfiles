package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// TODO: add a command for replacing context break logic like what is mentioned here https://www.accessiblepublishing.ca/common-epub-issues/#ContextBreaks
// one for characters, page breaks, and images
/**
Steps for replacing the character breaks:
- Identify the symbol that indicates this
- Find all lines that contain this symbol
- Verify that it needs replacing (manual confirmation)
- Append css to the proper css file
*/

var (
	filePath string
	lang     string
)

// lintEpubCmd represents the lintEpub command
var lintEpubCmd = &cobra.Command{
	Use:   "lint-epub",
	Short: "Takes the opf file of an epub and uses that to lint the files within it",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint lint-epub -f opf-file-path
	`,
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		EpubLint(log, fileHandler, filePath, lang)
	},
}

func init() {
	rootCmd.AddCommand(lintEpubCmd)

	lintEpubCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the opf file of the epub to lint")
	lintEpubCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	lintEpubCmd.MarkFlagRequired("file-path")
}

func EpubLint(l logger.Logger, fileManager filehandler.FileManager, filePath, lang string) {
	validateLintEpubFlags(l, fileManager, filePath)

	opfText := fileManager.ReadInFileContents(filePath)

	epubInfo, err := linter.ParseOpfFile(opfText)
	if err != nil {
		l.WriteError(fmt.Sprintf("Failed to parse \"%s\": %s", filePath, err))
	}

	var opfFolderString = fileManager.GetFileFolder(filePath)

	validateFilesExist(l, fileManager, opfFolderString, epubInfo.HtmlFiles)
	validateFilesExist(l, fileManager, opfFolderString, epubInfo.ImagesFiles)
	validateFilesExist(l, fileManager, opfFolderString, epubInfo.OtherFiles)

	// fix up all xhtml files first
	for file := range epubInfo.HtmlFiles {
		var filePath = getFilePath(fileManager, opfFolderString, file)
		fileText := fileManager.ReadInFileContents(filePath)
		var newText = linter.EnsureEncodingIsPresent(fileText)
		newText = linter.CommonStringReplace(newText)

		// TODO: remove images links that do not exist in the manifest
		// TODO: remove files that exist, but are not in the manifest
		newText = linter.EnsureLanguageIsSet(newText, lang)
		epubInfo.PageIds = linter.GetPageIdsForFile(l, newText, file, epubInfo.PageIds)

		if fileText == newText {
			continue
		}

		fileManager.WriteFileContents(filePath, newText)
	}

	updateNavFile(l, fileManager, opfFolderString, epubInfo.NavFile, epubInfo.PageIds)
	updateNcxFile(l, fileManager, opfFolderString, epubInfo.NcxFile, epubInfo.PageIds)

	// TODO; cleanup TOC file's links
}

func validateLintEpubFlags(l logger.Logger, fileManager filehandler.FileManager, filePath string) {
	if !strings.HasSuffix(filePath, ".opf") {
		l.WriteError(fmt.Sprintf(`file-path: "%s" must be an opf file`, filePath))
	}

	if !fileManager.FileExists(filePath) {
		l.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}

func validateFilesExist(l logger.Logger, fileManager filehandler.FileManager, opfFolder string, files map[string]struct{}) {
	for file := range files {
		var filePath = getFilePath(fileManager, opfFolder, file)

		if !fileManager.FileExists(filePath) {
			l.WriteError(fmt.Sprintf(`file from manifest not found: "%s" must exist`, filePath))
		}
	}
}

func updateNcxFile(l logger.Logger, fileManager filehandler.FileManager, opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(fileManager, opfFolder, file)
	fileText := fileManager.ReadInFileContents(filePath)

	newText, err := linter.CleanupNavMap(fileText)
	if err != nil {
		l.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNcxFile(newText, pageIds)

	if fileText == newText {
		return
	}

	fileManager.WriteFileContents(filePath, newText)
}

func updateNavFile(l logger.Logger, fileManager filehandler.FileManager, opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(fileManager, opfFolder, file)
	fileText := fileManager.ReadInFileContents(filePath)

	newText, err := linter.RemoveIdsFromNav(fileText)
	if err != nil {
		l.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNavFile(newText, pageIds)

	if fileText == newText {
		return
	}

	fileManager.WriteFileContents(filePath, newText)
}

func getFilePath(fileManager filehandler.FileManager, opfFolder, file string) string {
	return fileManager.JoinPath(opfFolder, file)
}
