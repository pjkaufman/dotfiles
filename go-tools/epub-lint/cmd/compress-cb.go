package cmd

// import (
// 	"fmt"

// 	"github.com/pjkaufman/dotfiles/go-tools/utils"
// 	"github.com/spf13/cobra"
// )

// const cbzExt = "cbz"
// const separator = "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"

// var imageExtsToCompress = []string{".jpg", ".jpeg", ".png"}
// var imgpCompressionArgs = []string{"-x", "800x800", "-e", "-O", "-q", "40", "-w", "-m"}

// // compressCBZCmd represents the compressCBZ command
// var compressCBZCmd = &cobra.Command{
// 	Use:   "compress-cbz",
// 	Short: "Takes in a file extension of either cbz or cbr and then compresses all of the images in these files in the current directory",
// 	Long: `Compress jpeg and png images in cbz or cbr files in the current directory

// 	For example: epub-lint compress-cb
// 	`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		utils.WriteInfo("Starting compression for each cbz in dir")

// 		var files = utils.GetFilesInCurrentDirWithExt(cbzExt)
// 		for _, file := range files {
// 			utils.WriteInfo(fmt.Sprintf("Starting cbz compressing for \"%s\"...", file))

// 			err := utils.UnzipRezipAndRunOperation(file, cbzExt, func() error {
// 				utils.MustChangeDirectoryTo(cbzExt)

// 				var images = utils.GetFilesInFolderWithExts(".", imageExtsToCompress...)
// 				for _, image := range images {
// 					utils.MustRunCommand("imgp", fmt.Sprintf("failed to update image \"%s\"", image), append(imgpCompressionArgs, fmt.Sprintf(`"%s"`, image))...)
// 				}

// 				utils.MustChangeDirectoryTo("..")

// 				return nil
// 			})
// 			if err != nil {
// 				utils.WriteError(fmt.Sprintf("failed to unzip and rezip file \"%s\": %s", file, err))
// 			}

// 			var newOriginalFile = file + ".original"
// 			utils.MustRename(file, newOriginalFile)
// 			utils.MustRename("compress.zip", file)

// 			utils.WriteInfo(separator)
// 			utils.WriteInfo(fmt.Sprintf(`%.1fM %s`, utils.MustGetMbSize(newOriginalFile), newOriginalFile))
// 			utils.WriteInfo(fmt.Sprintf(`%.1fM %s`, utils.MustGetMbSize(file), file))
// 			utils.WriteInfo(separator)
// 		}
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(compressCBZCmd)
// }
