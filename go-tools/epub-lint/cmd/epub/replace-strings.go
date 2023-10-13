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

var epubFile string
var extraReplacesFilePath string

const (
	EpubPathArgEmpty           = "epub-file must have a non-whitespace value"
	EpubPathArgNonEpub         = "epub-file must be an Epub file"
	ExtraStringReplaceArgNonMd = "extra-replace-file must be a Markdown file"
	ExtraStringReplaceArgEmpty = "extra-replace-file must have a non-whitespace value"
)

// replaceStringsCmd represents the replaceStrings command
var replaceStringsCmd = &cobra.Command{
	Use:   "replace-strings",
	Short: "Replaces a list of common strings and the extra strings for all content/xhtml files in the provided epub",
	Long: `Uses the provided epub and extra replace Markdown file to replace a common set of strings and any extra
	instances specified in the extra file replace. After all replacements are made, the original epub will be moved
	to a .original file and the new file will take the place of the old file. It will also print out the successful
	extra replacements with the number 	of replacements made followed by warnings for any extra strings that it tried
	to find and replace values for, but did not find any instances to replace.
	Note: it only replaces strings in content/xhtml files listed in the opf file.

	For example: epub-lint epub replace-strings -f test.epub -e replacements.md
	will replace the common strings and the extra strings parsed out of replacements.md
	in content/xhtml files located in test.epub. The original test.epub will be moved to test.epub.original and
	test.epub will have the updated files.

	replacements.md is expected to be in the following format:
    | Text to replace | Text to replace with |
    | --------------- | -------------------- |
    | I am typo | I the correct value |
	...
	| I am another issue to correct | the correction |
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateReplaceStringsFlags(epubFile, extraReplacesFilePath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		filehandler.FileMustExist(epubFile, "epub-file")
		filehandler.FileMustExist(extraReplacesFilePath, "extra-replace-file")

		logger.WriteInfo("Starting epub string replacement...\n")

		var numHits = make(map[string]int)
		var extraTextReplacements = linter.ParseTextReplacements(filehandler.ReadInFileContents(extraReplacesFilePath))

		var epubFolder = filehandler.GetFileFolder(epubFile)
		var dest = filehandler.JoinPath(epubFolder, "epub")
		err = filehandler.UnzipRunOperationAndRezip(epubFile, dest, func() {
			opfFolder, epubInfo := getEpubInfo(dest, epubFile)
			validateFilesExist(opfFolder, epubInfo.HtmlFiles)

			for file := range epubInfo.HtmlFiles {
				var filePath = getFilePath(opfFolder, file)
				fileText := filehandler.ReadInFileContents(filePath)

				var newText = linter.CommonStringReplace(fileText)
				newText = linter.ExtraStringReplace(newText, extraTextReplacements, numHits)

				if fileText == newText {
					continue
				}

				filehandler.WriteFileContents(filePath, newText)
			}

			var successfulReplaces []string
			var failedReplaces []string
			for searchText, hits := range numHits {
				if hits == 0 {
					failedReplaces = append(failedReplaces, searchText)
				} else {
					var timeText = "time"
					if hits > 1 {
						timeText += "s"
					}

					successfulReplaces = append(successfulReplaces, fmt.Sprintf("`%s` was replaced %d %s", searchText, hits, timeText))
				}
			}

			logger.WriteInfo("Successful Replaces:")
			for _, successfulReplace := range successfulReplaces {
				logger.WriteInfo(successfulReplace)
			}

			if len(failedReplaces) == 0 {
				return
			}

			logger.WriteInfo("")
			logger.WriteWarn("Failed Replaces:")
			for i, failedReplace := range failedReplaces {
				logger.WriteWarn(fmt.Sprintf("%d. %s", i+1, failedReplace))
			}
		})
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteInfo("\nFinished epub string replacement...")
	},
}

func init() {
	EpubCmd.AddCommand(replaceStringsCmd)

	replaceStringsCmd.Flags().StringVarP(&epubFile, "epub-file", "f", "", "the epub file to do the string replacements in")
	replaceStringsCmd.Flags().StringVarP(&extraReplacesFilePath, "extra-replace-file", "e", "", "the path to the file with extra strings to replace")
	replaceStringsCmd.MarkFlagRequired("epub-file")
	replaceStringsCmd.MarkFlagRequired("extra-replace-file")
}

func ValidateReplaceStringsFlags(epubPath, extraReplaceStringsPath string) error {
	if strings.TrimSpace(epubPath) == "" {
		return errors.New(EpubPathArgEmpty)
	}

	if !strings.HasSuffix(epubPath, ".epub") {
		return errors.New(EpubPathArgNonEpub)
	}

	if strings.TrimSpace(extraReplaceStringsPath) == "" {
		return errors.New(ExtraStringReplaceArgEmpty)
	}

	if !strings.HasSuffix(extraReplaceStringsPath, ".md") {
		return errors.New(ExtraStringReplaceArgNonMd)
	}

	return nil
}
