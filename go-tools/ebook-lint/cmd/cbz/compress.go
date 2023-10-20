package cbz

import (
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc"
	filesize "github.com/pjkaufman/dotfiles/go-tools/ebook-lint/file-size"
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	lintDir string
	verbose bool
)

const (
	LintDirArgEmpty = "directory must have a non-whitespace value"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compresses all of the png and jpeg files in the cbz files in the specified directory.",
	Example: heredoc.Doc(`To compress images in all cbzs in a folder:
	ebook-lint cbz compress -d folder
	
	To compress images in all cbzs in the current directory:
	ebook-lint cbz compress -d folder
	`),
	Long: "Gets all of the .cbz files in the specified directory and compresses pngs and jpegs.",
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCompressFlags(lintDir)
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteInfo("Starting compression and linting for each epub\n")

		cbzs := filehandler.MustGetAllFilesWithExtInASpecificFolder(lintDir, ".cbz")

		var totalBeforeFileSize, totalAfterFileSize float64
		for _, cbz := range cbzs {
			logger.WriteInfo(fmt.Sprintf("starting cbz compression for %s...", cbz))

			compressCbz(lintDir, cbz)

			var originalFile = cbz + ".original"

			var newKbSize = filehandler.MustGetFileSize(filehandler.JoinPath(lintDir, cbz))
			var oldKbSize = filehandler.MustGetFileSize(filehandler.JoinPath(lintDir, originalFile))

			logger.WriteInfo(filesize.FileSizeSummary(originalFile, cbz, oldKbSize, newKbSize))

			totalBeforeFileSize += oldKbSize
			totalAfterFileSize += newKbSize
		}

		logger.WriteInfo(filesize.FilesSizeSummary(totalBeforeFileSize, totalAfterFileSize))

		logger.WriteInfo("Finished compression and linting")
	},
}

func init() {
	CbzCmd.AddCommand(compressCmd)

	compressCmd.Flags().StringVarP(&lintDir, "directory", "d", ".", "the location to run the cbz image compression in")
	compressCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "whether or not to show extra information about the image compression")
}

func compressCbz(lintDir, cbz string) {
	var src = filehandler.JoinPath(lintDir, cbz)
	var dest = filehandler.JoinPath(lintDir, "cbz")

	filehandler.UnzipRunOperationAndRezip(src, dest, func() {
		var imageFiles = filehandler.MustGetAllFilesWithExtsInASpecificFolderAndSubFolders(dest, commandhandler.CompressableImageExts...)

		for i, imageFile := range imageFiles {
			if verbose {
				logger.WriteInfo(fmt.Sprintf(`%d of %d: compressing "%s"`, i, len(imageFiles), imageFile))
			}

			commandhandler.CompressImage(imageFile)
		}
	})
}

func ValidateCompressFlags(lintDir string) error {
	if strings.TrimSpace(lintDir) == "" {
		return errors.New(LintDirArgEmpty)
	}

	return nil
}
