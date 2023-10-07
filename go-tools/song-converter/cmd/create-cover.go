package cmd

import (
	"fmt"
	"strings"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/spf13/cobra"
)

var coverOutputFile string
var coverInputFilePath string

// createCoverCmd represents the createCover command
var createCoverCmd = &cobra.Command{
	Use:   "create-cover",
	Short: "Takes in the cover file path and creates the html cover file",
	Long: `Takes in the cover file to make the html cover file
	
	For example: song-converter create-cover -f cover-file -o output-file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		CreateCover(log, fileHandler, coverInputFilePath, coverOutputFile)
	},
}

func init() {
	rootCmd.AddCommand(createCoverCmd)

	createCoverCmd.Flags().StringVarP(&coverInputFilePath, "cover-file", "f", "", "the markdown cover file source")
	createCoverCmd.Flags().StringVarP(&coverOutputFile, "output", "o", "", "the html file to write the output to")
	createCoverCmd.MarkFlagRequired("cover-file")
	createCoverCmd.MarkFlagRequired("styles-file")
}

func CreateCover(l logger.Logger, fileManager filehandler.FileManager, songsCoverFilePath, outputFile string) {
	validateCreateCoverFlags(l, fileManager, songsCoverFilePath)

	l.WriteInfo("Converting files to html cover")

	var htmlFile = strings.Builder{}

	var styles = fileManager.ReadInFileContents(stylesFilePath)
	htmlFile.WriteString(styles)

	var coverHtml = converter.ConvertMdToHtmlCover(l, fileManager, songsCoverFilePath)
	htmlFile.WriteString(coverHtml)

	var htmlOutput = htmlFile.String()
	writeToFileOrStdOut(l, fileManager, htmlOutput, outputFile)

	l.WriteInfo("Finished creating html cover file")
}

func validateCreateCoverFlags(l logger.Logger, fileManager filehandler.FileManager, songsCoverFilePath string) {
	if strings.Trim(songsCoverFilePath, " ") == "" {
		l.WriteError("cover-file must have a non-whitespace value")
	}

	if !strings.HasSuffix(songsCoverFilePath, ".md") {
		l.WriteError("cover-file must be an md file")
	}

	if !fileManager.FileExists(songsCoverFilePath) {
		l.WriteError(fmt.Sprintf(`cover-file: "%s" must exist`, songsCoverFilePath))
	}
}
