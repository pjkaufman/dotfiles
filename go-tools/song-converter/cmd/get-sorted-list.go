package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var stagingDir string

// getSortedListCmd represents the getSortedList command
var getSortedListCmd = &cobra.Command{
	Use:   "get-sorted-list",
	Short: "Gets an alphabetically sorted list of Markdown file names from the provided directory",
	Long: `Gets the list of Markdown files in the working directory provided and sorts them alphabetically
	
	For example: song-converter get-sorted-list -f working-dir
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateGetSortedListFlags(stagingDir)

		files := utils.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")

		sort.Strings(files)

		for i, fileName := range files {
			files[i] = utils.JoinPath(stagingDir, fileName)
		}

		utils.WriteInfo(strings.Join(files, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(getSortedListCmd)

	getSortedListCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	getSortedListCmd.MarkFlagRequired("working-dir")
}

func validateGetSortedListFlags(stagingDir string) {
	if strings.Trim(stagingDir, " ") == "" {
		utils.WriteError("working-dir must have a non-whitespace value")
	}

	if !utils.FileExists(stagingDir) {
		utils.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}
}
