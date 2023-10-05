package cmd

import (
	"fmt"
	"sort"
	"strings"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var stagingDir string

// getSortedListCmd represents the getSortedList command
var getSortedListCmd = &cobra.Command{
	Use:   "get-sorted-list",
	Short: "Gets an alphabetically sorted list of Markdown file names from the provided directory",
	Long: `Gets the list of Markdown files in the working directory provided and sorts them alphabetically
	
	For example: song-converter get-sorted-list -d working-dir
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		SortMarkdownFileList(log, fileHandler, stagingDir)
	},
}

func init() {
	rootCmd.AddCommand(getSortedListCmd)

	getSortedListCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	getSortedListCmd.MarkFlagRequired("working-dir")
}

func SortMarkdownFileList(l logger.Logger, fileManager filehandler.FileManager, stagingDir string) {
	validateGetSortedListFlags(l, fileManager, stagingDir)

	files := fileManager.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")

	sort.Strings(files)

	for i, fileName := range files {
		files[i] = fileManager.JoinPath(stagingDir, fileName)
	}

	l.WriteInfo(strings.Join(files, "\n"))
}

func validateGetSortedListFlags(l logger.Logger, fileManager filehandler.FileManager, stagingDir string) {
	if strings.Trim(stagingDir, " ") == "" {
		l.WriteError("working-dir must have a non-whitespace value")
	}

	if !fileManager.FolderExists(stagingDir) {
		l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}
}
