package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

// updateCmd represents the create command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the submodule branch to the specified branch name",
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var cmdHanlder = commandhandler.NewCommandHandler(log)
		var fileHandler = filehandler.NewFileHandler(log)
		UpdateSubmoduleBranches(log, cmdHanlder, fileHandler, ticketAbbreviation, branchName, repoFolderPath)
	},
}

func init() {
	submoduleCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&submoduleName, "submodule", "s", "", "the name of the submodule to operate on")
	updateCmd.Flags().StringVarP(&repoFolderPath, "repo-parent-path", "p", "", "the path to the parent folder of the repos to operate on")
	updateCmd.Flags().StringVarP(&branchName, "branch-name", "b", "", "the submodule branch name to checkout and use")
	updateCmd.MarkFlagRequired("submodule")
	updateCmd.MarkFlagRequired("repo-parent-path")
	updateCmd.MarkFlagRequired("branch-name")
}

func UpdateSubmoduleBranches(l logger.Logger, cmdManager commandhandler.CommandManager, fileManager filehandler.FileManager, ticketAbbreviation, branchName, repoFolderPath string) {
	validateSubmoduleUpdate(l, fileManager, ticketAbbreviation, branchName)

	folders := getListOfFoldersWithSubmodule(fileManager, repoFolderPath, submoduleName)
	for _, folder := range folders {
		var submoduleDir = filepath.Join(append(pathToSubmodule, submoduleName)...)
		cmdManager.MustChangeDirectoryTo(submoduleDir)
		checkoutLatestFromMaster(cmdManager, submoduleDir)

		cmdManager.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "checkout", branchName)

		cmdManager.MustChangeDirectoryTo(upADirectory)

		cmdManager.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to stage changes to "%s" for "%s"`, submoduleName, folder), "add", submoduleName)
		cmdManager.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to commit changes for "%s"`, folder), "commit", "-m", "Updated "+submoduleName)
		cmdManager.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to push changes for "%s"`, folder), "push")
	}
}

func validateSubmoduleUpdate(l logger.Logger, fileManager filehandler.FileManager, ticketAbbreviation, branchName string) {
	if strings.Trim(branchName, " ") == "" {
		l.WriteError("branch-name must have a non-whitespace value")
		return
	}

	if !fileManager.FolderExists(repoFolderPath) {
		l.WriteError("repo-parent-path must exist and be a directory")
		return
	}
}
