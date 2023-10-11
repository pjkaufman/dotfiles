package cmd

import (
	"errors"
	"fmt"
	"strings"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/spf13/cobra"
)

const (
	CoverPathArgEmpty  = "cover-file must have a non-whitespace value"
	CoverPathNotMdFile = "cover-file must be an md file"
)

var coverOutputFile string
var coverInputFilePath string

// createCoverCmd represents the createCover command
var createCoverCmd = &cobra.Command{
	Use:   "create-cover",
	Short: "Takes in the cover file path and creates the html cover file",
	Long: `Takes in the cover file to make the html cover file
	
	For example: song-converter create-cover -f cover-file -o output-file
	Converts the cover file from Markdown into html as the specified output file.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCreateCoverFlags(coverInputFilePath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		if !filehandler.FileExists(coverInputFilePath) {
			logger.WriteError(fmt.Sprintf(`cover-file: "%s" must exist`, coverInputFilePath))
		}

		logger.WriteInfo("Converting files to html cover")

		var coverMd = filehandler.ReadInFileContents(coverInputFilePath)
		htmlFile := converter.BuildHtmlCover(coverMd)

		writeToFileOrStdOut(htmlFile, coverOutputFile)

		logger.WriteInfo("Finished creating html cover file")
	},
}

func init() {
	rootCmd.AddCommand(createCoverCmd)

	createCoverCmd.Flags().StringVarP(&coverInputFilePath, "cover-file", "f", "", "the markdown cover file source")
	createCoverCmd.Flags().StringVarP(&coverOutputFile, "output", "o", "", "the html file to write the output to")
	createCoverCmd.MarkFlagRequired("cover-file")
}

func ValidateCreateCoverFlags(songsCoverFilePath string) error {
	if strings.TrimSpace(songsCoverFilePath) == "" {
		return errors.New(CoverPathArgEmpty)
	}

	if !strings.HasSuffix(songsCoverFilePath, ".md") {
		return errors.New(CoverPathNotMdFile)
	}

	return nil
}
