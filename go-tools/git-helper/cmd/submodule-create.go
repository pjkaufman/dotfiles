package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var (
	submoduleName        string
	repoFolderPath       string
	ticketAbbreviation   string
	branchName           string
	pathToSubmodule      = []string{"src", "modules"}
	getCurrentBranchArgs = []string{"branch", "--show-current"}
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the branch in the specified submodule if it does not already exist",
	Long: `Creates the specified branch in the provided submodule for all instances of the submodule in the provided folder so long as it is not already present.
	
	For example: git-tools submodule create -s Submodule -p ./repos/ -a ticket-abbreviation -b branch-name
	This command would create a 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ticketAbbreviation = strings.Trim(ticketAbbreviation, " ")
		if ticketAbbreviation == "" {
			os.Stderr.WriteString("ticket-abbreviation must have a non-whitespace value\n")
		}

		branchName = strings.Trim(branchName, " ")
		if branchName == "" {
			utils.WriteError("branch-name must have a non-whitespace value")
		}

		repoPathExists := utils.FolderExists(repoFolderPath)
		if !repoPathExists {
			utils.WriteError("repo-parent-path must exist and be a directory")
		}

		// fmt.Printf(`create -s "%s" -p "%s" -a "%s" -b "%s"`+"\n", submoduleName, repoFolderPath, ticketAbbreviation, branchName)

		folders := getListOfFoldersWithSubmodule(repoFolderPath, submoduleName)
		// fmt.Println(folders)

		var currentBranch string
		for _, folder := range folders {
			utils.MustChangeDirectoryTo(folder)

			currentBranch = utils.MustGetCommandOutput(gitProgramName, fmt.Sprintf(`failed to get current branch for "%s"`, folder), getCurrentBranchArgs...)
			if strings.Contains(currentBranch, ticketAbbreviation) {
				continue
			}

			utils.WriteInfo(currentBranch + " does not contain " + ticketAbbreviation)
			createSubmoduleUpdateBranch(folder, submoduleName)
		}
	},
}

func init() {
	submoduleCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&submoduleName, "submodule", "s", "", "the name of the submodule to operate on")
	createCmd.Flags().StringVarP(&repoFolderPath, "repo-parent-path", "p", "", "the path to the parent folder of the repos to operate on")
	createCmd.Flags().StringVarP(&ticketAbbreviation, "ticket-abbreviation", "a", "", "the ticket abbreviation to use to determine whether we should update a repo and to help determine the name for submodule branch")
	createCmd.Flags().StringVarP(&branchName, "branch-name", "b", "", "the submodule branch name to checkout and use")
	createCmd.MarkFlagRequired("submodule")
	createCmd.MarkFlagRequired("repo-parent-path")
	createCmd.MarkFlagRequired("ticket-abbreviation")
	createCmd.MarkFlagRequired("branch-name")
}

func getListOfFoldersWithSubmodule(path, submoduleName string) []string {
	var folders []string
	for _, dir := range utils.GetFoldersInCurrentFolder(path) {
		var pathParts = []string{path, dir}
		var folderPath = filepath.Join(pathParts...)
		pathParts = append(pathParts, pathToSubmodule...)
		pathParts = append(pathParts, submoduleName)
		var submoduleFolderPath = filepath.Join(pathParts...)

		var exists = utils.FolderExists(submoduleFolderPath)
		if !exists {
			continue
		}

		folders = append(folders, folderPath)
	}

	return folders
}

func createSubmoduleUpdateBranch(folder, submodule string) {
	utils.WriteInfo("Creating the DE branch for " + folder)
	checkoutLatestFromMaster(folder)

	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "checkout", "-B", "vikings/"+ticketAbbreviation+"update-"+submodule)

	var submoduleDir = filepath.Join(append(pathToSubmodule, submodule)...)
	utils.MustChangeDirectoryTo(filepath.Join(append(pathToSubmodule, submodule)...))

	checkoutLatestFromMaster(submoduleDir)

	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to checkout "%s" for "%s"`, branchName, folder), "checkout", branchName)

	utils.MustChangeDirectoryTo(upADirectory)

	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to stage changes to "%s" for "%s"`, submodule, folder), "add", submodule)
	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to commit changes for "%s"`, folder), "commit", "-m", "Updated "+submodule)
	pushOutput := utils.MustGetCommandOutput(gitProgramName, fmt.Sprintf(`failed to push changes for "%s"`, folder), "push")

	fmt.Println("TODO: handle the push output - " + pushOutput)
}

func checkoutLatestFromMaster(folder string) {
	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to checkout master for "%s"`, folder), "checkout", "master")
	utils.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "pull")
}
