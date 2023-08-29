package main

import "github.com/pjkaufman/dotfiles/go-tools/git-helper/cmd"

func main() {
	cmd.Execute()
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

// def updateMode():
//     deBranch = input("DE branch to checkout: ").strip()
//     if deBranch == "":
//         print("Please enter a DE branch to use.")
//         return

//     folders = getListOfFoldersWithSubmodules()
//     for folder in folders:
//         os.chdir(folder)
//         os.chdir(SUBMODULE_PATH)
//         os.system("git checkout master && git pull")
//         # need to checkout the branch specified
//         os.system("git checkout " + deBranch)
//         os.chdir("..")
//         os.system('git add DataEntities && git commit -m "Updated DE" && git push')

//         os.chdir("../..")

//         os.chdir("..")

// def navigateToBaseDir(baseDir):
//     os.chdir(baseDir)
//     print(os.getcwd())

// def main():
//     print("\nRun Modes")
//     print(
//         "Update - Updates the submodule reference to DE to the specified branch. This should only be done once create has been run once."
//     )
//     print(
//         "Create - Creates DE update branches for the Realm repos that do not have the ticket abbreviation\n"
//     )
//     mode = input("Enter run mode: ").lower().strip()

//     navigateToBaseDir(ABSOLUTE_PATH_TO_REALM_REPOS_BASE_DIR)
//     if "update" == mode:
//         updateMode()
//     elif "create" == mode:
//         createMode()
//     else:
//         print("Please enter a valid run mode.")
//         return

// if __name__ == "__main__":
//     main()
