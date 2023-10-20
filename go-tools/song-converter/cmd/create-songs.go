package cmd

import (
	"errors"
	"sort"
	"strings"

	"github.com/MakeNowJust/heredoc"
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
	Example: heredoc.Doc(`To write the output of converting the files in the specified folder to html to a file:
	song-converter create-songs -d working-dir -s styles.html -o songs.html

	To write the output of converting the files in the specified folder to html to std out:
	song-converter create-songs -d working-dir -s styles.html
	`),
	Long: heredoc.Doc(`How it works:
	- Reads in all of the files in the specified folder
	- Sorts the files alphabetically
	- Adds the styles to the start of the html
	- Converts each file into html
	- Writes the content to the specified source
	`),
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCreateSongsFlags(stagingDir, stylesFilePath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		filehandler.FolderMustExist(stagingDir, "working-dir")
		filehandler.FileMustExist(stylesFilePath, "styles-file")

		var isWritingToFile = strings.TrimSpace(coverOutputFile) == ""
		if isWritingToFile {
			logger.WriteInfo("Converting Markdown files to html")
		}

		var styles = filehandler.ReadInFileContents(stylesFilePath)

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
		}

		htmlFile, err := converter.BuildHtmlBody(styles, mdInfo)
		if err != nil {
			logger.WriteError(err.Error())
		}

		writeToFileOrStdOut(htmlFile, bodyHtmlOutputFile)

		if isWritingToFile {
			logger.WriteInfo("Finished converting Markdown files to html")
		}
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
