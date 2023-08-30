package main

import "github.com/pjkaufman/dotfiles/go-tools/git-helper/cmd"

func main() {
	cmd.Execute()
}

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
