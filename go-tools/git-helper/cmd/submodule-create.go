package cmd

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
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

var prLinkRegex = regexp.MustCompile(`https:[^\n]*`)

const (
	TicketArgEmpty         = "ticket-abbreviation must have a non-whitespace value"
	BranchNameArgEmpty     = "branch-name must have a non-whitespace value"
	RepoParentPathArgEmpty = "repo-parent-path must have a non-whitespace value"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates the branch in the specified submodule if it does not already exist",
	Long: `Creates the specified branch in the provided submodule for all instances of the submodule in the provided folder so long as it is not already present.
	
	For example: git-tools submodule create -s Submodule -p ./repos/ -a ticket-abbreviation -b branch-name
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateSubmoduleCreate(ticketAbbreviation, branchName, repoFolderPath)
		if err != nil {
			logger.WriteError(err.Error())
		}

		if !filehandler.FolderExists(repoFolderPath) {
			logger.WriteError(`repo-parent-path: "%s" must exist`)
		}

		// fmt.Printf(`create -s "%s" -p "%s" -a "%s" -b "%s"`+"\n", submoduleName, repoFolderPath, ticketAbbreviation, branchName)

		folders := getListOfFoldersWithSubmodule(repoFolderPath, submoduleName)
		// fmt.Println(folders)

		var currentBranch string
		for _, folder := range folders {
			commandhandler.MustChangeDirectoryTo(folder)

			currentBranch = commandhandler.MustGetCommandOutput(gitProgramName, fmt.Sprintf(`failed to get current branch for "%s"`, folder), getCurrentBranchArgs...)
			if strings.Contains(currentBranch, ticketAbbreviation) {
				continue
			}

			logger.WriteInfo(currentBranch + " does not contain " + ticketAbbreviation)
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
	for _, dir := range filehandler.GetFoldersInCurrentFolder(path) {
		var pathParts = []string{path, dir}
		var folderPath = filepath.Join(pathParts...)
		pathParts = append(pathParts, pathToSubmodule...)
		pathParts = append(pathParts, submoduleName)
		var submoduleFolderPath = filepath.Join(pathParts...)

		var exists = filehandler.FolderExists(submoduleFolderPath)
		if !exists {
			continue
		}

		folders = append(folders, folderPath)
	}

	return folders
}

func createSubmoduleUpdateBranch(folder, submodule string) {
	logger.WriteInfo("Creating the DE branch for " + folder)
	checkoutLatestFromMaster(folder)

	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "checkout", "-B", "vikings/"+ticketAbbreviation+"update-"+submodule)

	var submoduleDir = filepath.Join(append(pathToSubmodule, submodule)...)
	commandhandler.MustChangeDirectoryTo(filepath.Join(append(pathToSubmodule, submodule)...))

	checkoutLatestFromMaster(submoduleDir)

	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to checkout "%s" for "%s"`, branchName, folder), "checkout", branchName)

	commandhandler.MustChangeDirectoryTo(upADirectory)

	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to stage changes to "%s" for "%s"`, submodule, folder), "add", submodule)
	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to commit changes for "%s"`, folder), "commit", "-m", "Updated "+submodule)
	pushOutput := commandhandler.MustGetCommandOutput(gitProgramName, fmt.Sprintf(`failed to push changes for "%s"`, folder), "push")

	fmt.Println("TODO: handle the push output - " + pushOutput)
}

func checkoutLatestFromMaster(folder string) {
	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to checkout master for "%s"`, folder), "checkout", "master")
	commandhandler.MustRunCommand(gitProgramName, fmt.Sprintf(`failed to pull latest changes for "%s"`, folder), "pull")
}

func ValidateSubmoduleCreate(ticketAbbreviation, branchName, repoFolderPath string) error {
	if strings.TrimSpace(ticketAbbreviation) == "" {
		return errors.New(TicketArgEmpty)
	}

	if strings.TrimSpace(branchName) == "" {
		return errors.New(BranchNameArgEmpty)
	}

	if strings.TrimSpace(repoFolderPath) == "" {
		return errors.New(RepoParentPathArgEmpty)
	}

	return nil
}

func GetPullRequestLink(pushOutput string) string {
	var matches = prLinkRegex.FindAllString(pushOutput, 1)
	if len(matches) == 0 {
		return ""
	}

	return matches[0]
}
