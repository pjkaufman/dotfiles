package cmd

import (
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	"github.com/spf13/cobra"
)

// resetCmd represents the submoduleReset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the submodules in the current repo to what it is on master",
	Run: func(cmd *cobra.Command, args []string) {
		commandhandler.MustRunCommand(gitProgramName, "failed to update the submodule for the current repo", "submodule", "update", "--init", "--recursive")
	},
}

func init() {
	submoduleCmd.AddCommand(resetCmd)
}
