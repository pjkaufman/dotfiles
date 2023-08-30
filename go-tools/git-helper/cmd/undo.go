/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

// undoCmd represents the undo command
var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undoes the previous commit while still retaining it",
	Run: func(cmd *cobra.Command, args []string) {
		utils.MustRunCommand(gitProgramName, "failed to update the submodule for the current repo", " reset", "--soft", "HEAD~1")
	},
}

func init() {
	rootCmd.AddCommand(undoCmd)
}
