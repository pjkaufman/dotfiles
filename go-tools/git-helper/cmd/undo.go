package cmd

import (
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undoes the previous commit while still retaining it",
	Run: func(cmd *cobra.Command, args []string) {
		UndoCommit(commandhandler.NewCommandHandler(logger.NewLoggerHandler()))
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}

func UndoCommit(cm commandhandler.CommandManager) {
	cm.MustRunCommand(gitProgramName, "failed to undo the last commit for the current repo", "reset", "--soft", "HEAD~1")
}
