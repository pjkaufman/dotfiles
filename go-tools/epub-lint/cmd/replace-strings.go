package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var filePaths string
var extraReplacesFilePath string

const (
	FilePathsArgEmpty          = "file-paths must have a non-whitespace value"
	ExtraStringReplaceArgNonMd = "extra-replace-file-path must be a Markdown file"
	ExtraStringReplaceArgEmpty = "extra-replace-file-path must have a non-whitespace value"
)

// replaceStringsCmd represents the replaceStrings command
var replaceStringsCmd = &cobra.Command{
	Use:   "replace-strings",
	Short: "Replaces a list of common strings and the extra strings specified in the extra file for the provide file(s)",
	Long: `Goes and replaces a common set of strings in a file as well as any extra instances that are specified.
	It will print out the successful extra replacements with the number of replacements made followed by warnings
	for any extra strings that it tried to find and replace values for, but did not find any instances to replace.
	
	For example: epub-lint replace-strings -f file-paths -e extra-replace-file-path
	will replace the common strings and the extra strings parsed out of the extra replace file 
	from the provided file(s)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var files = strings.Split(filePaths, ",")
		err := ValidateReplaceStringsFlags(files, extraReplacesFilePath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		for _, filePath := range files {
			filehandler.FileMustExist(filePath, "file-paths")
		}

		filehandler.FileMustExist(extraReplacesFilePath, "extra-replace-file-path")

		var numHits = make(map[string]int)
		var extraTextReplacements = linter.ParseTextReplacements(filehandler.ReadInFileContents(extraReplacesFilePath))
		for _, filePath := range files {
			fileText := filehandler.ReadInFileContents(filePath)
			var newText = linter.CommonStringReplace(fileText)
			newText = linter.ExtraStringReplace(newText, extraTextReplacements, numHits)

			if fileText == newText {
				continue
			}

			filehandler.WriteFileContents(filePath, newText)
		}

		if len(numHits) == 0 {
			logger.WriteWarn("No values were listed as needing replacing")

			return
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
		logger.WriteInfo("")
	},
}

func init() {
	rootCmd.AddCommand(replaceStringsCmd)

	replaceStringsCmd.Flags().StringVarP(&filePaths, "file-paths", "f", "", "the list of files to update in a comma separated list")
	replaceStringsCmd.Flags().StringVarP(&extraReplacesFilePath, "extra-replace-file-path", "e", "", "the path to the file with extra strings to replace")
	replaceStringsCmd.MarkFlagRequired("file-paths")
	replaceStringsCmd.MarkFlagRequired("extra-replace-file-path")
}

func ValidateReplaceStringsFlags(filePaths []string, extraReplaceStringsPath string) error {
	if len(filePaths) == 0 || (len(filePaths) == 1 && strings.TrimSpace(filePaths[0]) == "") {
		return errors.New(FilePathsArgEmpty)
	}

	if strings.TrimSpace(extraReplaceStringsPath) == "" {
		return errors.New(ExtraStringReplaceArgEmpty)
	}

	if !strings.HasSuffix(extraReplaceStringsPath, ".md") {
		return errors.New(ExtraStringReplaceArgNonMd)
	}

	return nil
}
