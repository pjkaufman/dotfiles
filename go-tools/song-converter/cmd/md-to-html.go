package cmd

import (
	"fmt"
	"strings"

	"github.com/adrg/frontmatter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var filePath string

type SongMetadata struct {
	Melody   string `yaml:"melody"`
	Key      string `yaml:"key"`
	Authors  string `yaml:"authors"`
	InChurch string `yaml:"in-church"`
	Verse    string `yaml:"verse"`
	Location string `yaml:"location"`
}

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
		MdToHtml(log, fileHandler, filePath)
	},
}

func init() {
	rootCmd.AddCommand(mdToHtmlCmd)

	mdToHtmlCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the Markdown file to convert to html")
	mdToHtmlCmd.MarkFlagRequired("file-path")
}

func MdToHtml(l logger.Logger, fileManager filehandler.FileManager, filePath string) string {
	validateMdToHtmlFlags(l, fileManager, filePath)

	contents := fileManager.ReadInFileContents(filePath)

	var metadata SongMetadata
	_, err := frontmatter.Parse(strings.NewReader(contents), &metadata)
	if err != nil {
		l.WriteError(fmt.Sprintf(`There was an error getting the frontmatter for file '%s': %s`, filePath, err))
	}

	return ""
}

func validateMdToHtmlFlags(l logger.Logger, fileManager filehandler.FileManager, filePath string) {
	if strings.Trim(filePath, " ") == "" {
		l.WriteError("file-path must have a non-whitespace value")
	}

	if !strings.HasSuffix(filePath, ".md") {
		l.WriteError("file-path must be a markdown file")
	}

	if !fileManager.FileExists(filePath) {
		l.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}
