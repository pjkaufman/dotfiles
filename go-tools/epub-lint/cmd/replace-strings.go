package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var filePaths string
var extraReplacesFilePath string

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
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		ReplaceExtraStrings(log, fileHandler, filePaths, extraReplacesFilePath)
	},
}

func init() {
	rootCmd.AddCommand(replaceStringsCmd)

	replaceStringsCmd.Flags().StringVarP(&filePaths, "file-paths", "f", "", "the list of files to update in a comma separated list")
	replaceStringsCmd.Flags().StringVarP(&extraReplacesFilePath, "extra-replace-text", "e", "", "the path to the file with extra strings to replace")
	replaceStringsCmd.MarkFlagRequired("file-paths")
	replaceStringsCmd.MarkFlagRequired("extra-replace-text")
}

func ReplaceExtraStrings(l logger.Logger, fileManager filehandler.FileManager, filePaths, extraReplacesFilePath string) {
	var files = strings.Split(filePaths, ",")
	validateReplaceStringsFlags(l, fileManager, files, extraReplacesFilePath)

	var numHits = make(map[string]int)
	var extraTextReplacements = linter.ParseTextReplacements(l, fileManager.ReadInFileContents(extraReplacesFilePath))
	for _, filePath := range files {
		fileText := fileManager.ReadInFileContents(filePath)
		var newText = linter.CommonStringReplace(fileText)
		newText = linter.ExtraStringReplace(newText, extraTextReplacements, numHits)

		if fileText == newText {
			continue
		}

		fileManager.WriteFileContents(filePath, newText)
	}

	if len(numHits) == 0 {
		l.WriteWarn("No values were listed as needing replacing")

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

	l.WriteInfo("Successful Replaces:")
	for _, successfulReplace := range successfulReplaces {
		l.WriteInfo(successfulReplace)
	}

	if len(failedReplaces) == 0 {
		return
	}

	l.WriteInfo("")
	l.WriteWarn("Failed Replaces:")
	for i, failedReplace := range failedReplaces {
		l.WriteWarn(fmt.Sprintf("%d. %s", i+1, failedReplace))
	}
	l.WriteInfo("")
}

func validateReplaceStringsFlags(l logger.Logger, fileManager filehandler.FileManager, filePaths []string, extraReplaceStringsPath string) {
	for _, filePath := range filePaths {
		filePathExists := fileManager.FileExists(filePath)

		if !filePathExists {
			l.WriteError(fmt.Sprintf(`file-paths: "%s" must exist`, filePath))
		}
	}

	if strings.Trim(extraReplacesFilePath, " ") == "" {
		l.WriteError("extra-replace-file-path must have a non-whitespace value")
	}

	extraReplacesFilePathExists := fileManager.FileExists(extraReplacesFilePath)
	if !extraReplacesFilePathExists {
		l.WriteError("extra-replace-file-path must exist")
	}
}
