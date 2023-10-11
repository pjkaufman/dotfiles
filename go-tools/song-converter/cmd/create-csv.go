package cmd

import (
	"errors"
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
		var l = logger.NewLoggerHandler()
		var fileManager = filehandler.NewFileHandler(l)

		err := ValidateCreateCsvFlags(stagingDir)
		if err != nil {
			l.WriteError(err.Error())
		}

		if !fileManager.FolderExists(stagingDir) {
			l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
		}

		l.WriteInfo("Converting Markdown files to csv")

		files := fileManager.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")
		sort.Strings(files)

		var mdInfo = make([]converter.MdFileInfo, len(files))

		// var csvContents = strings.Builder{}
		// csvContents.WriteString("Song|Location|Author|Copyright\n")

		for i, fileName := range files {
			var filePath = fileManager.JoinPath(stagingDir, fileName)
			var contents = fileManager.ReadInFileContents(filePath)

			mdInfo[i] = converter.MdFileInfo{
				FilePath:     filePath,
				FileName:     fileName,
				FileContents: contents,
			}

			// csvString, err := converter.ConvertMdToCsv(fileName, filePath, contents)
			// if err != nil {
			// 	l.WriteError(err.Error())
			// }

			// csvContents.WriteString(csvString)
		}

		csvFile, err := converter.BuildCsv(mdInfo)
		if err != nil {
			l.WriteError(err.Error())
		}

		writeToFileOrStdOut(l, fileManager, csvFile, outputFile)

		l.WriteInfo("Finished converting Markdown files to csv")
	},
}

func init() {
	rootCmd.AddCommand(createCsvCmd)

	createCsvCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	createCsvCmd.Flags().StringVarP(&outputFile, "output-file", "o", "", "the file to write the csv to")
	createCsvCmd.MarkFlagRequired("working-dir")
}

func ValidateCreateCsvFlags(stagingDir string) error {
	if strings.TrimSpace(stagingDir) == "" {
		return errors.New(StagingDirArgEmpty)
	}

	return nil
}
