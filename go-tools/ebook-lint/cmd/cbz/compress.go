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
	dir     string
	verbose bool
)

const (
	DirArgEmpty = "directory must have a non-whitespace value"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compresses all of the png and jpeg files in the cbz files in the specified directory",
	Example: heredoc.Doc(`To compress images in all cbzs in a folder:
	ebook-lint cbz compress -d folder
	
	To compress images in all cbzs in the current directory:
	ebook-lint cbz compress
	`),
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCompressFlags(dir)
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteInfo("Started compressing all cbzs\n")

		cbzs := filehandler.MustGetAllFilesWithExtInASpecificFolder(dir, ".cbz")

		var totalBeforeFileSize, totalAfterFileSize float64
		for _, cbz := range cbzs {
			logger.WriteInfo(fmt.Sprintf("starting cbz compression for %s...", cbz))

			compressCbz(dir, cbz)

			var originalFile = cbz + ".original"

			var newKbSize = filehandler.MustGetFileSize(filehandler.JoinPath(dir, cbz))
			var oldKbSize = filehandler.MustGetFileSize(filehandler.JoinPath(dir, originalFile))

			logger.WriteInfo(filesize.FileSizeSummary(originalFile, cbz, oldKbSize, newKbSize))

			totalBeforeFileSize += oldKbSize
			totalAfterFileSize += newKbSize
		}

		logger.WriteInfo(filesize.FilesSizeSummary(totalBeforeFileSize, totalAfterFileSize))

		logger.WriteInfo("Finished compressing all cbzs")
	},
}

func init() {
	CbzCmd.AddCommand(compressCmd)

	compressCmd.Flags().StringVarP(&dir, "directory", "d", ".", "the location to run the cbz image compression in")
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

func ValidateCompressFlags(dir string) error {
	if strings.TrimSpace(dir) == "" {
		return errors.New(DirArgEmpty)
	}

	return nil
}
