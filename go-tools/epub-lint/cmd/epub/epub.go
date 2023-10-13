package epub

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

// EpubCmd represents the epub command
var EpubCmd = &cobra.Command{
	Use:   "epub",
	Short: "Deals with epub related commands",
	Long:  `Handles operations on epub files in particular`,
}

func init() {
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
