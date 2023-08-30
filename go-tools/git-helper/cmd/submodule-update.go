package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

// updateCmd represents the create command
var updateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the branch in the specified submodule if it does not already exist",
	Long: `Creates the specified branch in the provided submodule for all instances of the submodule in the provided folder so long as it is not already present.
	
	For example: git-tools submodule create -s Submodule -p ./repos/ -a ticket-abbreviation -b branch-name
	This command would create a 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		branchName = strings.Trim(branchName, " ")
		if branchName == "" {
			utils.WriteError("branch-name must have a non-whitespace value")
			return
		}

		repoPathExists := utils.FolderExists(repoFolderPath)
		if !repoPathExists {
			utils.WriteError("repo-parent-path must exist and be a directory")
			return
		}

		folders := getListOfFoldersWithSubmodule(repoFolderPath, submoduleName)
		for _, folder := range folders {
			var submoduleDir = filepath.Join(append(pathToSubmodule, submoduleName)...)
			utils.MustChangeDirectoryTo(submoduleDir)
			checkoutLatestFromMaster(submoduleDir)

			utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "checkout", branchName)

			utils.MustChangeDirectoryTo(upADirectory)

			utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to stage changes to "%s" for "%s"`, submoduleName, folder), "add", submoduleName)
			utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to commit changes for "%s"`, folder), "commit", "-m", "Updated "+submoduleName)
			utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to push changes for "%s"`, folder), "push")
		}
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
