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
	FilePathArgEmpty        = "file-path must have a non-whitespace value"
	FilePathNotMarkdownFile = "file-path must be a markdown file"
)

var stagingDir string
var stylesFilePath string
var bodyHtmlOutputFile string

// mdToHtmlCmd represents the mdToHtml command
var mdToHtmlCmd = &cobra.Command{
	Use:   "md-to-html",
	Short: "Converts the provided Markdown file's contents to html and puts it back out to std out",
	Long: `Takes the contents of a Markdown file and converts it to html based on the YAML frontmatter contents
	
	For example: song-converter md-to-html -f file-path
	converts the Markdown file to an html file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		MdToHtml(log, fileHandler, stagingDir, stylesFilePath, bodyHtmlOutputFile)
	},
}

func init() {
	rootCmd.AddCommand(mdToHtmlCmd)

	mdToHtmlCmd.Flags().StringVarP(&stagingDir, "working-dir", "d", "", "the directory where the Markdown files are located")
	mdToHtmlCmd.Flags().StringVarP(&stylesFilePath, "styles-file", "s", "", "the html styles file to start the cover with")
	mdToHtmlCmd.Flags().StringVarP(&bodyHtmlOutputFile, "output", "o", "", "the html file to write the output to")
	mdToHtmlCmd.MarkFlagRequired("styles-file")
	mdToHtmlCmd.MarkFlagRequired("working-dir")
	//-s styles-file
}

func MdToHtml(l logger.Logger, fileManager filehandler.FileManager, stagingDir, stylesFilePath, outputFile string) {
	validateMdToHtmlFlags(l, fileManager, stagingDir, stylesFilePath)

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

func validateMdToHtmlFlags(l logger.Logger, fileManager filehandler.FileManager, stagingDir, stylesFilePath string) {
	if strings.Trim(stagingDir, " ") == "" {
		l.WriteError("working-dir must have a non-whitespace value")
	}

	if !fileManager.FolderExists(stagingDir) {
		l.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
	}

	if strings.Trim(stylesFilePath, " ") == "" {
		l.WriteError("styles-file must have a non-whitespace value")
	}

	if !strings.HasSuffix(stylesFilePath, ".html") {
		l.WriteError("styles-file must be an html file")
	}

	if !fileManager.FileExists(stylesFilePath) {
		l.WriteError(fmt.Sprintf(`styles-file: "%s" must exist`, stylesFilePath))
	}
}
