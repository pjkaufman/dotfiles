package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
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

		contents := fileManager.ReadInFileContents(filePath)
		var metadata SongMetadata
		_, err := frontmatter.Parse(strings.NewReader(contents), &metadata)
		if err != nil {
			l.WriteError(fmt.Sprintf(`There was an error getting the frontmatter for file '%s': %s`, filePath, err))
		}

		csvContents.WriteString(fileName + "|" + buildMetadataCsv(&metadata) + "\n")
	}

	if strings.Trim(outputFile, " ") != "" {
		fileManager.WriteFileContents(outputFile, csvContents.String())
	} else {
		l.WriteInfo(csvContents.String())
	}
}

func validateCreateCsvFlags(l logger.Logger, fileManager filehandler.FileManager, stagingDir string) {
	if strings.Trim(stagingDir, " ") == "" {
		l.WriteError("working-dir must have a non-whitespace value")
	}

	if !fileManager.FolderExists(stagingDir) {
		l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}
}

func buildMetadataCsv(metadata *SongMetadata) string {
	if metadata == nil {
		return "||"
	}

	var copyright = metadata.Copyright
	if strings.EqualFold(metadata.InChurch, "Y") {
		copyright = "Church"
	}

	return fmt.Sprintf("%s|%s|%s", metadata.BookLocation, metadata.Authors, copyright)
}
