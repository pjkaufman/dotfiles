package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var filePath string

// mdToHtmlCmd represents the mdToHtml command
var mdToHtmlCmd = &cobra.Command{
	Use:   "md-to-html",
	Short: "Converts the provided Markdown file's contents to html and puts it back out to std out",
	Long: `Takes the contents of a Markdown file and converts it to html based on the YAML frontmatter contents
	
	For example: song-converter md-to-html -f file-path
	converts the Markdown file to an html file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateMdToHtmlFlags(filePath)
	},
}

func init() {
	rootCmd.AddCommand(mdToHtmlCmd)

	mdToHtmlCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the Markdown file to convert to html")
	mdToHtmlCmd.MarkFlagRequired("file-path")
}

func validateMdToHtmlFlags(filePath string) {
	if strings.Trim(filePath, " ") == "" {
		utils.WriteError("file-path must have a non-whitespace value")
	}

	if !strings.HasSuffix(filePath, ".md") {
		utils.WriteError("file-path must be a markdown file")
	}

	if !utils.FileExists(filePath) {
		utils.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}
