package cmd

import (
	"errors"
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
		err := ValidateCreateCsvFlags(stagingDir)
		if err != nil {
			logger.WriteError(err.Error())
		}

		filehandler.FolderMustExist(stagingDir, "working-dir")

		logger.WriteInfo("Converting Markdown files to csv")

		files := filehandler.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")
		sort.Strings(files)

		var mdInfo = make([]converter.MdFileInfo, len(files))

		for i, fileName := range files {
			var filePath = filehandler.JoinPath(stagingDir, fileName)
			var contents = filehandler.ReadInFileContents(filePath)

			mdInfo[i] = converter.MdFileInfo{
				FilePath:     filePath,
				FileName:     fileName,
				FileContents: contents,
			}
		}

		csvFile, err := converter.BuildCsv(mdInfo)
		if err != nil {
			logger.WriteError(err.Error())
		}

		writeToFileOrStdOut(csvFile, outputFile)

		logger.WriteInfo("Finished converting Markdown files to csv")
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
