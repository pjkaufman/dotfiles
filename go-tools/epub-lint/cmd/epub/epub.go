package epub

import (
	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/cmd"
	"github.com/spf13/cobra"
)

// epubCmd represents the epub command
var epubCmd = &cobra.Command{
	Use:   "epub",
	Short: "Deals with epub related commands",
	Long:  `Handles operations on epub files in particular`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	cmd.RootCmd.AddCommand(epubCmd)
}
