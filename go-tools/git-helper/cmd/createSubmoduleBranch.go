package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var submoduleName string
var repoFolderPath string
var ticketAbbreviation string
var branchName string
var pathToSubmodule = []string{"src", "modules"}

// createSubmoduleBranchCmd represents the createSubmoduleBranch command
var createSubmoduleBranchCmd = &cobra.Command{
	Use:   "createSubmoduleBranch",
	Short: "Creates the branch in the specified submodule if it does not already exist",
	Long: `Creates the specified branch in the provided submodule for all instances of the submodule in the provided folder so long as it is not already present.
	
	For example: git-tools create -s Submodule -p ./repos/ -a ticket-abbreviation -b branch-name
	This command would create a 
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createSubmoduleBranch called")
		ticketAbbreviation = strings.Trim(ticketAbbreviation, " ")
		if ticketAbbreviation == "" {
			os.Stderr.WriteString("ticket-abbreviation must have a non-whitespace value\n")
			return
		}

		branchName = strings.Trim(branchName, " ")
		if branchName == "" {
			utils.WriteError("branch-name must have a non-whitespace value")
			return
		}

		var repoPathExists bool
		var err error
		repoPathExists, err = checkIfFolderExists(repoFolderPath)
		if err != nil {
			utils.WriteError("could not verify that repo-parent-path exists and is a directory: \"" + err.Error() + "\"")
			return
		}

		if !repoPathExists {
			utils.WriteError("repo-parent-path must exist and be a directory")
			return
		}

		fmt.Printf(`createSubmoduleBranch -s "%s" -p "%s" -a "%s" -b "%s"`+"\n", submoduleName, repoFolderPath, ticketAbbreviation, branchName)

		folders, _ := getListOfFoldersWithSubmodule(repoFolderPath, submoduleName)
		fmt.Println(folders)
	},
}

func init() {
	rootCmd.AddCommand(createSubmoduleBranchCmd)

	createSubmoduleBranchCmd.Flags().StringVarP(&submoduleName, "submodule", "s", "", "the name of the submodule to operate on")
	createSubmoduleBranchCmd.Flags().StringVarP(&repoFolderPath, "repo-parent-path", "p", "", "the path to the parent folder of the repos to operate on")
	createSubmoduleBranchCmd.Flags().StringVarP(&ticketAbbreviation, "ticket-abbreviation", "a", "", "the ticket abbreviation to use to determine whether we should update a repo and to help determine the name for submodule branch")
	createSubmoduleBranchCmd.Flags().StringVarP(&branchName, "branch-name", "b", "", "the submodule branch name to checkout and use")
	createSubmoduleBranchCmd.MarkFlagRequired("submodule")
	createSubmoduleBranchCmd.MarkFlagRequired("repo-parent-path")
	createSubmoduleBranchCmd.MarkFlagRequired("ticket-abbreviation")
	createSubmoduleBranchCmd.MarkFlagRequired("branch-name")
}

func checkIfFolderExists(path string) (bool, error) {
	folderInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	if !folderInfo.IsDir() {
		return false, nil
	}

	return true, nil
}

func getListOfFoldersWithSubmodule(path, submoduleName string) ([]string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var folders []string
	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		// var submodulePath = strings.ReplaceAll(pathToSubmodule, "/", string(os.PathSeparator))
		var pathParts = []string{path, dir.Name()}
		pathParts = append(pathParts, pathToSubmodule...)
		pathParts = append(pathParts, submoduleName)
		var submoduleFolderPath = filepath.Join(pathParts...)

		var exists bool
		exists, err = checkIfFolderExists(submoduleFolderPath)
		if err != nil {
			return nil, err
		}

		if !exists {
			continue
		}

		folders = append(folders, dir.Name())
	}
	// def getListOfFoldersWithSubmodules():
	//     filenames = os.listdir(".")

	//     foldersList = []
	//     for filename in filenames:
	//         filePath = os.path.join(os.path.abspath("."), filename)
	//         if os.path.isdir(filePath) and os.path.exists(
	//             os.path.join(filePath, SUBMODULE_PATH)
	//         ):
	//             foldersList.append(filename)

	//     return foldersList
	// err := os.Chdir(repoFolderPath)
	// if err != nil {
	// 	return nil, err
	// }

	return folders, nil
}

// def createMode():
//     ticketAbbreviation = input("Enter ticket abbreviation (i.e. VIK-224): ").strip()

//     # verify that the ticket abbreviation exists so we can properly create branches
//     # and verify when we do and do not need to create new branches
//     if ticketAbbreviation == "":
//         print("Please enter a ticket abbreviation.")
//         return

//     deBranch = input("DE branch to checkout: ").strip()
//     if deBranch == "":
//         print("Please enter a DE branch to use.")
//         return

//     folders = getListOfFoldersWithSubmodules()
//     for folder in folders:
//         os.chdir(folder)

//         getBranchNameCmd = ["git", "branch", "--show-current"]
//         branchName = str(
//             subprocess.Popen(getBranchNameCmd, stdout=subprocess.PIPE).communicate()[0]
//         )

//         print(branchName)

//         if ticketAbbreviation not in branchName:
//             print("Creating the DE branch for " + folder)
//             os.system("git checkout master && git pull")
//             # need to create a branch for the DE update
//             os.system("git checkout -B vikings/" + ticketAbbreviation + "-update-de")

//             os.chdir(SUBMODULE_PATH)
//             os.system("git checkout master && git pull")

//             # need to checkout the branch specified
//             os.system("git checkout " + deBranch)
//             os.chdir("..")
//             os.system('git add DataEntities && git commit -m "Updated DE" && git push')

//             os.chdir("../..")

//         os.chdir("..")
