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
		err := ValidateCreateSongsFlags(stagingDir, stylesFilePath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		if !filehandler.FolderExists(stagingDir) {
			logger.WriteError(fmt.Sprintf(`working-dir: "%s" must exist`, stagingDir))
		}

		if !filehandler.FileExists(stylesFilePath) {
			logger.WriteError(fmt.Sprintf(`styles-file: "%s" must exist`, stylesFilePath))
		}

		logger.WriteInfo("Converting Markdown files to html")

		var styles = filehandler.ReadInFileContents(stylesFilePath)

		// var htmlFile = strings.Builder{}
		// htmlFile.WriteString(styles + "\n")

		files := filehandler.MustGetAllFilesWithExtInASpecificFolder(stagingDir, ".md")
		sort.Strings(files)

		var mdInfo = make([]converter.MdFileInfo, len(files))

		for i, fileName := range files {
			var filePath = filehandler.JoinPath(stagingDir, fileName)
			fileContents := filehandler.ReadInFileContents(filePath)

			mdInfo[i] = converter.MdFileInfo{
				FilePath:     filePath,
				FileName:     fileName,
				FileContents: fileContents,
			}
			// html, err := converter.ConvertMdToHtmlSong(filePath, fileContents)
			// if err != nil {
			// 	l.WriteError(err.Error())
			// }

			// htmlFile.WriteString(html + "\n")
		}

		htmlFile, err := converter.BuildHtmlBody(styles, mdInfo)
		if err != nil {
			logger.WriteError(err.Error())
		}

		writeToFileOrStdOut(htmlFile, outputFile)

		logger.WriteInfo("Finished converting Markdown files to html")
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

func ValidateCreateSongsFlags(stagingDir, stylesFilePath string) error {
	if strings.TrimSpace(stagingDir) == "" {
		return errors.New(StagingDirArgEmpty)
	}

	if strings.TrimSpace(stylesFilePath) == "" {
		return errors.New(StylesPathArgEmpty)
	}

	if !strings.HasSuffix(stylesFilePath, ".html") {
		return errors.New(StylesPathNotHtmlFile)
	}

	return nil
}
