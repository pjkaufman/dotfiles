package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

// TODO: add a command for replacing context break logic like what is mentioned here https://www.accessiblepublishing.ca/common-epub-issues/#ContextBreaks
// one for characters, page breaks, and images

// lintEpubCmd represents the lintEpub command
var lintEpubCmd = &cobra.Command{
	Use:   "lint-epub",
	Short: "Takes the opf file of an epub and uses that to lint the files within it",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint lint-epub -f opf-file-path
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateLintEpubFlags(filePath)

		opfText := utils.ReadInFileContents(filePath)

		epubInfo, err := linter.ParseOpfFile(opfText)
		if err != nil {
			utils.WriteError(fmt.Sprintf("Failed to parse \"%s\": %s", filePath, err))
		}

		var opfFolderString = utils.GetFileFolder(filePath)

		validateFilesExist(opfFolderString, epubInfo.HtmlFiles)
		validateFilesExist(opfFolderString, epubInfo.ImagesFiles)
		validateFilesExist(opfFolderString, epubInfo.OtherFiles)

		// fix up all xhtml files first
		for file := range epubInfo.HtmlFiles {
			var filePath = getFilePath(opfFolderString, file)
			fileText := utils.ReadInFileContents(filePath)
			var newText = linter.EnsureEncodingIsPresent(fileText)
			newText = linter.CommonStringReplace(newText)

			// TODO: remove images links that do not exist in the manifest
			// TODO: remove files that exist, but are not in the manifest
			newText = linter.EnsureLanguageIsSet(newText, lang)
			epubInfo.PageIds = linter.GetPageIdsForFile(newText, file, epubInfo.PageIds)

			if fileText == newText {
				continue
			}

			utils.WriteFileContents(filePath, newText)
		}

		updateNavFile(opfFolderString, epubInfo.NavFile, epubInfo.PageIds)
		updateNcxFile(opfFolderString, epubInfo.NcxFile, epubInfo.PageIds)

		// TODO; cleanup TOC file's links
	},
}

func init() {
	rootCmd.AddCommand(lintEpubCmd)

	lintEpubCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the opf file of the epub to lint")
	lintEpubCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	lintEpubCmd.MarkFlagRequired("file-path")
}

func validateLintEpubFlags(filePath string) {
	if !strings.HasSuffix(filePath, ".opf") {
		utils.WriteError(fmt.Sprintf(`file-path: "%s" must be an opf file`, filePath))
	}

	if !utils.FileExists(filePath) {
		utils.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}

func validateFilesExist(opfFolder string, files map[string]struct{}) {
	for file := range files {
		var filePath = getFilePath(opfFolder, file)

		if !utils.FileExists(filePath) {
			utils.WriteError(fmt.Sprintf(`file from manifest not found: "%s" must exist`, filePath))
		}
	}
}

func updateNcxFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := utils.ReadInFileContents(filePath)

	newText, err := linter.CleanupNavMap(fileText)
	if err != nil {
		utils.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNcxFile(newText, pageIds)

	if fileText == newText {
		return
	}

	utils.WriteFileContents(filePath, newText)
}

func updateNavFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := utils.ReadInFileContents(filePath)

	newText, err := linter.RemoveIdsFromNav(fileText)
	if err != nil {
		utils.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNavFile(newText, pageIds)

	if fileText == newText {
		return
	}

	utils.WriteFileContents(filePath, newText)
}

func getFilePath(opfFolder, file string) string {
	return utils.JoinPath(opfFolder, file)
}
