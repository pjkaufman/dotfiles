package cmd

import (
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undoes the previous commit while still retaining it",
	Run: func(cmd *cobra.Command, args []string) {
		commandhandler.MustRunCommand(gitProgramName, "failed to undo the last commit for the current repo", "reset", "--soft", "HEAD~1")
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
