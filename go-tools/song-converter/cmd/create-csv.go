package cmd

import (
	"fmt"
	"sort"
	"strings"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/spf13/cobra"
)

var outputFile string

// createCsvCmd represents the createCsv command
var createCsvCmd = &cobra.Command{
	Use:   "create-csv",
	Short: `Creates a "|" delimited csv file that includes metadata about songs like whether they are in the church or copyrighted`,
	Long: `Gets the list of Markdown files in the working directory provided and sorts them alphabetically
	
	For example: song-converter create-csv -d working-dir -o churchSongs.csv
	Iterates over all of the Markdown files in the specified directory and pulls out metadata
	like the author, book location, and copyright info to put in the csv file specified.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		CreateCsv(log, fileHandler, stagingDir, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(createCsvCmd)

	createCsvCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	createCsvCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "the file to write the csv to")
	createCsvCmd.MarkFlagRequired("working-dir")
}

func CreateCsv(l logger.Logger, fileManager filehandler.FileManager, stagingDir, outputFile string) {
	validateCreateCsvFlags(l, fileManager, stagingDir)

	l.WriteInfo("Converting Markdown files to csv")

	files := fileManager.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")

	sort.Strings(files)

	var csvContents = strings.Builder{}
	csvContents.WriteString("Song|Location|Author|Copyright\n")

	for _, fileName := range files {
		var filePath = fileManager.JoinPath(stagingDir, fileName)

		converter.ConvertMdToCsv(l, fileManager, fileName, filePath, csvContents)
	}

	var outputCsv = csvContents.String()
	outputCsv = strings.ReplaceAll(outputCsv, "&nbsp;", "")
	writeToFileOrStdOut(l, fileManager, outputCsv, outputFile)

	l.WriteInfo("Finished converting Markdown files to csv")
}

func validateCreateCsvFlags(l logger.Logger, fileManager filehandler.FileManager, stagingDir string) {
	if strings.Trim(stagingDir, " ") == "" {
		l.WriteError("working-dir must have a non-whitespace value")
	}

	if !fileManager.FolderExists(stagingDir) {
		l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}
}
