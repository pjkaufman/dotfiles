package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var filePath string

// lintFileCmd represents the lintFile command
var lintFileCmd = &cobra.Command{
	Use:   "lint-file",
	Short: "Lints the xhtml, htm, or html file provided making sure to fix several different issues that may show up",
	Long: `Lints the xhtml, htm, or html file provided by doing the following:
	- ensure that there is an encoding set which will be set to utf-8 if it is not already set
	- replace common strings in the file with its desired equivalent
	
	For example: epub-lint lint-file -f file-path
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateLintFileFlags(filePath)

		fileText := utils.ReadInFileContents(filePath)
		var newText = linter.EnsureEncodingIsPresent(fileText)
		newText = linter.CommonStringReplace(newText)
		newText, err := linter.RemoveIdsFromNav(newText)
		if err != nil {
			utils.WriteError(fmt.Sprintf("%s: %v", filePath, err))
		}

		// TODO: handle lang="en" xml:lang="en" in html tag in non-opf files

		if fileText == newText {
			return
		}

		utils.WriteFileContents(filePath, newText)
	},
}

func init() {
	rootCmd.AddCommand(lintFileCmd)

	lintFileCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the xhtml, htm, or html file to lint")
	lintFileCmd.MarkFlagRequired("file-path")
}

func validateLintFileFlags(filePath string) {
	if !strings.HasSuffix(filePath, ".html") && !strings.HasSuffix(filePath, ".xhtml") && !strings.HasSuffix(filePath, ".htm") {
		utils.WriteError(fmt.Sprintf(`file-path: "%s" must be an html, htm, or xhtml file`, filePath))
	}

	if !utils.FileExists(filePath) {
		utils.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}
