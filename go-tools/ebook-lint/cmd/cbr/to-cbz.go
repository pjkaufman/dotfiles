package cbr

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var dir string

const (
	DirArgEmpty = "directory must have a non-whitespace value"
)

// cbrToCbzCmd represents the cbzToCbr command
var cbrToCbzCmd = &cobra.Command{
	Use:   "to-cbz",
	Short: "Compresses all of the png and jpeg files in the cbz files in the specified directory.",
	Example: heredoc.Doc(`To compress images in all cbzs in a folder:
	ebook-lint cbz cbr-to-cbz -d folder
	
	To compress images in all cbzs in the current directory:
	ebook-lint cbz cbr-to-cbz -d folder
	`),
	Long: "Gets all of the .cbz files in the specified directory and cbzToCbres pngs and jpegs.",
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCbzToCbrFlags(dir)
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteInfo("Starting converting cbr files to cbz files\n")

		cbrs := filehandler.MustGetAllFilesWithExtInASpecificFolder(dir, ".cbr")
		for _, cbr := range cbrs {
			logger.WriteInfo(fmt.Sprintf("starting to convert %s to a cbz file...", cbr))

			filehandler.ConvertRarToCbz(cbr)
		}

		logger.WriteInfo("\nFinished converting cbr files to cbz files")
	},
}

func init() {
	CbrCmd.AddCommand(cbrToCbzCmd)

	cbrToCbzCmd.Flags().StringVarP(&dir, "directory", "d", ".", "the location to run the cbz image cbzToCbrion in")
}

func ValidateCbzToCbrFlags(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return errors.New(DirArgEmpty)
	}

	return nil
}
