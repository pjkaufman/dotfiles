package cmd

import (
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

// resetCmd represents the submoduleReset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Resets the submodules in the current repo to what it is on master",
	Run: func(cmd *cobra.Command, args []string) {
		utils.MustRunCommand(gitProgramName, "failed to update the submodule for the current repo", "submodule", "update", "--init", "--recursive")
	},
}

func init() {
	submoduleCmd.AddCommand(resetCmd)
}
