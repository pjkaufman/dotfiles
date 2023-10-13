package epub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	lintDir           string
	lang              string
	runCompressImages bool
)

const (
	LintDirArgEmpty = "directory must have a non-whitespace value"
	LangArgEmpty    = "lang must have a non-whitespace value"
)

// compressAndLintCmd represents the compressAndLint command
var compressAndLintCmd = &cobra.Command{
	Use:   "compress-and-lint",
	Short: "Takes the opf file of an epub and uses that to lint the files within it",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint epub compress-and-lint
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCompressAndLintFlags(lintDir, lang)
		if err != nil {
			logger.WriteError(err.Error())
		}

		epubs := filehandler.MustGetAllFilesWithExtInASpecificFolder(lintDir, ".epub")
		logger.WriteInfo(strings.Join(epubs, "\n"))
		var epub = epubs[0]
		var src = filehandler.JoinPath(lintDir, epub)
		var dest = filehandler.JoinPath(lintDir, "epub")

		filehandler.UnzipRunOperationAndRezip(src, dest, func() {
			opfFolder, epubInfo := getEpubInfo(dest, epub)

			validateFilesExist(opfFolder, epubInfo.HtmlFiles)
			validateFilesExist(opfFolder, epubInfo.ImagesFiles)
			validateFilesExist(opfFolder, epubInfo.OtherFiles)

			// fix up all xhtml files first
			for file := range epubInfo.HtmlFiles {
				var filePath = getFilePath(opfFolder, file)
				fileText := filehandler.ReadInFileContents(filePath)
				var newText = linter.EnsureEncodingIsPresent(fileText)
				newText = linter.CommonStringReplace(newText)

				// TODO: remove images links that do not exist in the manifest
				newText = linter.EnsureLanguageIsSet(newText, lang)
				epubInfo.PageIds = linter.GetPageIdsForFile(newText, file, epubInfo.PageIds)

				if fileText == newText {
					continue
				}

				filehandler.WriteFileContents(filePath, newText)
			}

			updateNavFile(opfFolder, epubInfo.NavFile, epubInfo.PageIds)
			updateNcxFile(opfFolder, epubInfo.NcxFile, epubInfo.PageIds)
			//TODO: get all files in the repo and prompt the user whether they want to delete them

			if runCompressImages {
				compressImages(lintDir, opfFolder, epubInfo.ImagesFiles)
			}

			// TODO: cleanup TOC file's links
		})
	},
}

func init() {
	EpubCmd.AddCommand(compressAndLintCmd)

	compressAndLintCmd.Flags().StringVarP(&lintDir, "directory", "d", ".", "the location to run the epub lint logic")
	compressAndLintCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	compressAndLintCmd.Flags().BoolVarP(&runCompressImages, "compress-images", "i", false, "whether or not to also compress images which requires imgp to be installed")
}

func ValidateCompressAndLintFlags(lintDir, lang string) error {
	if strings.TrimSpace(lintDir) == "" {
		return errors.New(LintDirArgEmpty)
	}

	if strings.TrimSpace(lang) == "" {
		return errors.New(LangArgEmpty)
	}

	return nil
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
