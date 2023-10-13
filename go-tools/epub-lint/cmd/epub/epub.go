package epub

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

const (
	EpubPathArgEmpty   = "epub-file must have a non-whitespace value"
	EpubPathArgNonEpub = "epub-file must be an Epub file"
)

var epubFile string

// EpubCmd represents the epub command
var EpubCmd = &cobra.Command{
	Use:   "epub",
	Short: "Deals with epub related commands",
	Long:  `Handles operations on epub files in particular`,
}

func init() {
	EpubCmd.PersistentFlags().StringVarP(&epubFile, "epub-file", "f", "", "the epub file to work on")
	EpubCmd.MarkPersistentFlagRequired("epub-file")
}

func validateCommonEpubFlags(epubPath string) error {
	if strings.TrimSpace(epubPath) == "" {
		return errors.New(EpubPathArgEmpty)
	}

	if !strings.HasSuffix(epubPath, ".epub") {
		return errors.New(EpubPathArgNonEpub)
	}

	return nil
}
