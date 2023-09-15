package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var filePaths string
var extraReplacesFilePath string

// replaceStringsCmd represents the replaceStrings command
var replaceStringsCmd = &cobra.Command{
	Use:   "replace-strings",
	Short: "Replaces a common set of strings in a file",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint replace-strings -f file-paths -e extra-replace-file-path
	will replace the common strings and the extra strings parsed out of the extra replace file 
	from the provided file(s)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var files = strings.Split(filePaths, ",")
		validateReplaceStringsFlags(files, extraReplacesFilePath)

		var numHits = make(map[string]int)
		var extraTextReplacements = linter.ParseTextReplacements(utils.ReadInFileContents(extraReplacesFilePath))
		for _, filePath := range files {
			fileText := utils.ReadInFileContents(filePath)
			var newText = linter.CommonStringReplace(fileText)
			newText = linter.ExtraStringReplace(newText, extraTextReplacements, numHits)

			if fileText == newText {
				continue
			}

			utils.WriteFileContents(filePath, newText)
		}

		if len(numHits) == 0 {
			utils.WriteWarn("No values were listed as needing replacing")
		}

		for searchText, hits := range numHits {
			if hits == 0 {
				utils.WriteWarn(fmt.Sprintf("Did not find any replacements for `%s`", searchText))
			} else {
				utils.WriteInfo(fmt.Sprintf("`%s` was replaced %d time(s)", searchText, hits))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceStringsCmd)

	replaceStringsCmd.Flags().StringVarP(&filePaths, "file-paths", "f", "", "the list of files to update in a comma separated list")
	replaceStringsCmd.Flags().StringVarP(&extraReplacesFilePath, "extra-replace-text", "e", "", "the path to the file with extra strings to replace")
	replaceStringsCmd.MarkFlagRequired("file-paths")
	replaceStringsCmd.MarkFlagRequired("extra-replace-text")
}

func validateReplaceStringsFlags(filePaths []string, extraReplaceStringsPath string) {
	for _, filePath := range filePaths {
		filePathExists := utils.FileExists(filePath)

		if !filePathExists {
			utils.WriteError(fmt.Sprintf(`file-paths: "%s" must exist`, filePath))
		}
	}

	if strings.Trim(extraReplacesFilePath, " ") == "" {
		utils.WriteError("extra-replace-file-path must have a non-whitespace value")
	}

	extraReplacesFilePathExists := utils.FileExists(extraReplacesFilePath)
	if !extraReplacesFilePathExists {
		utils.WriteError("extra-replace-file-path must exist")
	}
}
