package cmd

import (
	"os"

	"github.com/pjkaufman/dotfiles/go-tools/ebook-lint/cmd/epub"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ebook-lint",
	Short: "A set of functions that are helpful for linting epubs",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(epub.EpubCmd)
}
