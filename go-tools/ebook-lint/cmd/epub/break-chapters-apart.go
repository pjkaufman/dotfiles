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
- Look for instances of "Chapter .+"
- For any hits, check with user if that is indeed a chapter indicator
- If so, indicate that that files needs to be split there
- Keep doing this until all instances of Chapter are found.
There should only be 1 per file.
- Then go ahead and copy the start of the file to the start of the body
plus the element for "Chapter .+" at the start of the file
- Create a new page for that chapter (chapter_.+.html)
- Add new page to opf
- Repeat this process until all instances are handled

For example: ebook-lint epub break-chapters -f test.epub
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
