package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	archivehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/archive-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var lintDir string

// lintEpubV2Cmd represents the lintEpubV2 command
var lintEpubV2Cmd = &cobra.Command{
	Use:   "lint",
	Short: "Takes the opf file of an epub and uses that to lint the files within it",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint lint
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// err := ValidateLintEpubV2Flags(lang)
		// if err != nil {
		// 	logger.WriteError(err.Error())
		// }

		epubs := filehandler.MustGetAllFilesWithExtInASpecificFolder(lintDir, ".epub")
		logger.WriteInfo(strings.Join(epubs, "\n"))
		var epub = epubs[0]
		var src = filehandler.JoinPath(lintDir, epub)
		var dest = filehandler.JoinPath(lintDir, "epub")

		err := archivehandler.UnzipRezipAndRunOperation(src, dest, func() error {
			logger.WriteInfo("Success?")

			opfFiles := filehandler.MustGetAllFilesWithExtInASpecificFolderAndSubFolders(dest, ".opf")
			if len(opfFiles) < 1 {
				logger.WriteError("did not find opf file...")
			}

			var filePath = opfFiles[0]
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
			return nil
		})
		if err != nil {
			logger.WriteError(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(lintEpubV2Cmd)

	lintEpubV2Cmd.Flags().StringVarP(&lintDir, "directory", "d", ".", "the location to run the epub lint logic")
	lintEpubV2Cmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
}

func ValidateLintEpubV2Flags(lang string) error {

	if strings.TrimSpace(lang) == "" {
		return errors.New(LangArgEmpty)
	}

	return nil
}
