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

const (
	StylesPathArgEmpty    = "styles-file must have a non-whitespace value"
	StylesPathNotHtmlFile = "styles-file must be a html file"
	StagingDirArgEmpty    = "working-dir must have a non-whitespace value"
)

var stagingDir string
var stylesFilePath string
var bodyHtmlOutputFile string

// CreateSongsCmd represents the CreateSongs command
var CreateSongsCmd = &cobra.Command{
	Use:   "create-songs",
	Short: "Converts all Markdown files in the specified folder into html in alphabetical order and starts that file with the styles provided",
	Long: `Takes in all of the Markdown files in the specified folder and converts them all to html in alphabetical order.
	The styles file provided will be the start of the generated html.
	
	For example: song-converter create-songs -d working-dir -s styles.html -o songs.html
	Converts the Markdown files in working-dir into html in alphabetical order with styles.html's contents starting the generated html.
	The contents will be written to songs.html.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		CreateSongs(log, fileHandler, stagingDir, stylesFilePath, bodyHtmlOutputFile)
	},
}

func init() {
	rootCmd.AddCommand(CreateSongsCmd)

	CreateSongsCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	CreateSongsCmd.Flags().StringVarP(&stylesFilePath, "styles-file", "s", "", "the html styles file to start the cover with")
	CreateSongsCmd.Flags().StringVarP(&bodyHtmlOutputFile, "output", "o", "", "the html file to write the output to")
	CreateSongsCmd.MarkFlagRequired("styles-file")
	CreateSongsCmd.MarkFlagRequired("working-dir")
}

func CreateSongs(l logger.Logger, fileManager filehandler.FileManager, stagingDir, stylesFilePath, outputFile string) {
	validateCreateSongsFlags(l, fileManager, stagingDir, stylesFilePath)

	l.WriteInfo("Converting Markdown files to html")

	var styles = fileManager.ReadInFileContents(stylesFilePath)

	var htmlFile = strings.Builder{}
	htmlFile.WriteString(styles + "\n")

	files := fileManager.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")
	sort.Strings(files)

	for _, fileName := range files {
		var filePath = fileManager.JoinPath(stagingDir, fileName)

		htmlFile.WriteString(converter.ConvertMdToHtmlSong(l, fileManager, filePath) + "\n")
	}

	writeToFileOrStdOut(l, fileManager, htmlFile.String(), outputFile)

	l.WriteInfo("Finished converting Markdown files to html")
}

func validateCreateSongsFlags(l logger.Logger, fileManager filehandler.FileManager, stagingDir, stylesFilePath string) {
	if strings.Trim(stagingDir, " ") == "" {
		l.WriteError(StagingDirArgEmpty)
	}

	if !fileManager.FolderExists(stagingDir) {
		l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}

	if strings.Trim(stylesFilePath, " ") == "" {
		l.WriteError(StylesPathArgEmpty)
	}

	if !strings.HasSuffix(stylesFilePath, ".html") {
		l.WriteError(StylesPathNotHtmlFile)
	}

	if !fileManager.FileExists(stylesFilePath) {
		l.WriteError(fmt.Sprintf(`styles-file: "%s" must exist`, stylesFilePath))
	}
}
