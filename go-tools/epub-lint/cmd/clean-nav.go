package cmd

import (
	"fmt"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

// cleanNavCmd represents the cleanNav command
var cleanNavCmd = &cobra.Command{
	Use:   "clean-nav",
	Short: "Cleans up the ncx navigation file",
	Long: `Goes and removes the ids from the navigation content src attributes since they do not seem to play well with the Kindle
	library. It will not remove an id if 2 hrefs are referencing the same file. It also makes sure that the play order increases one at a time.

	For example: epub-lint clean-nav -f file-path
	will remove ids from the navigation links in the file so long as there are not 2 references to the same file and update the play order attributes.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateCleanNavFlags(filePath)
		fileText := utils.ReadInFileContents(filePath)

		newText, err := linter.CleanupNavMap(fileText)
		if err != nil {
			utils.WriteError(fmt.Sprintf("%s: %v", filePath, err))
		}

		if fileText == newText {
			return
		}

		utils.WriteFileContents(filePath, newText)
	},
}

func init() {
	rootCmd.AddCommand(cleanNavCmd)

	cleanNavCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the ncx navigation file to cleanup")
	cleanNavCmd.MarkFlagRequired("file-path")
}

func validateCleanNavFlags(filePath string) {
	filePathExists := utils.FileExists(filePath)

	if !filePathExists {
		utils.WriteError("file-path must exist")
	}
}
