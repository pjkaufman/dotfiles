package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

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
		EpubLint(filePath, lang)
	},
}

func init() {
	rootCmd.AddCommand(lintEpubCmd)

	lintEpubCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the opf file of the epub to lint")
	lintEpubCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	lintEpubCmd.MarkFlagRequired("file-path")
}

func EpubLint(filePath, lang string) {
	validateLintEpubFlags(filePath)

	opfText := filehandler.ReadInFileContents(filePath)

	epubInfo, err := linter.ParseOpfFile(opfText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("Failed to parse \"%s\": %s", filePath, err))
	}

	var opfFolderString = filehandler.GetFileFolder(filePath)

	validateFilesExist(opfFolderString, epubInfo.HtmlFiles)
	validateFilesExist(opfFolderString, epubInfo.ImagesFiles)
	validateFilesExist(opfFolderString, epubInfo.OtherFiles)

	// fix up all xhtml files first
	for file := range epubInfo.HtmlFiles {
		var filePath = getFilePath(opfFolderString, file)
		fileText := filehandler.ReadInFileContents(filePath)
		var newText = linter.EnsureEncodingIsPresent(fileText)
		newText = linter.CommonStringReplace(newText)

		// TODO: remove images links that do not exist in the manifest
		// TODO: remove files that exist, but are not in the manifest
		newText = linter.EnsureLanguageIsSet(newText, lang)
		epubInfo.PageIds = linter.GetPageIdsForFile(newText, file, epubInfo.PageIds)

		if fileText == newText {
			continue
		}

		filehandler.WriteFileContents(filePath, newText)
	}

	updateNavFile(opfFolderString, epubInfo.NavFile, epubInfo.PageIds)
	updateNcxFile(opfFolderString, epubInfo.NcxFile, epubInfo.PageIds)

	// TODO: cleanup TOC file's links
}

func validateLintEpubFlags(filePath string) {
	if !strings.HasSuffix(filePath, ".opf") {
		logger.WriteError(fmt.Sprintf(`file-path: "%s" must be an opf file`, filePath))
	}

	if !filehandler.FileExists(filePath) {
		logger.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}

func validateFilesExist(opfFolder string, files map[string]struct{}) {
	for file := range files {
		var filePath = getFilePath(opfFolder, file)

		if !filehandler.FileExists(filePath) {
			logger.WriteError(fmt.Sprintf(`file from manifest not found: "%s" must exist`, filePath))
		}
	}
}

func updateNcxFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := filehandler.ReadInFileContents(filePath)

	newText, err := linter.CleanupNavMap(fileText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNcxFile(newText, pageIds)

	if fileText == newText {
		return
	}

	filehandler.WriteFileContents(filePath, newText)
}

func updateNavFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := filehandler.ReadInFileContents(filePath)

	newText, err := linter.RemoveIdsFromNav(fileText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNavFile(newText, pageIds)

	if fileText == newText {
		return
	}

	filehandler.WriteFileContents(filePath, newText)
}

func getFilePath(opfFolder, file string) string {
	return filehandler.JoinPath(opfFolder, file)
}
