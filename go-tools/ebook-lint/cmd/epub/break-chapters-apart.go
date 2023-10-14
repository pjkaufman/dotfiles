package epub

import (
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// breakChaptersApartCmd represents the breakChaptersApart command
var breakChaptersApartCmd = &cobra.Command{
	Use:   "break-chapters",
	Short: "Works on making sure that all chapters are located in their own files.",
	Long: `Works on making sure that all chapters are located in their own files.
How it works (not implemented yet):
- Look for instances of "["
- For any hits, check with user if the contents of [] should be considered
translator's notes
- If so, check the last 5 lines for "*" as the reference point
- If not found, then assume it is the end of the last line
- Lastly check for instances of the word note and repeat the
prior process
- Once all values have been found, make sure they are in order based
on the opf
- Create the TN page
- Add TN page to opf
- Replace references
- Finish everything up

For example: ebook-lint epub create-notes -f test.epub
`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateBreakChaptersApartFlags(epubFile)
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteWarn("TODO: implement logic described above")
	},
}

func init() {
	EpubCmd.AddCommand(breakChaptersApartCmd)

	breakChaptersApartCmd.Flags().StringVarP(&epubFile, "epub-file", "f", "", "the epub file to replace strings in in")
	breakChaptersApartCmd.MarkFlagRequired("epub-file")
}

func ValidateBreakChaptersApartFlags(epubPath string) error {
	err := validateCommonEpubFlags(epubPath)
	if err != nil {
		return err
	}

	return nil
}
