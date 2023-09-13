package cmd

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"

// 	"github.com/pjkaufman/dotfiles/go-tools/utils"
// 	"github.com/spf13/cobra"
// )

// var contentElRegex = regexp.MustCompile("(^<?xml\s+version=["'][\d.]+["']\s+encoding=["'][a-zA-Z\d-.]+["'].*?\?>)")
// var contentSrcAttributeRegex = regexp.MustCompile("(src=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")

// fixEncodingCmd represents the fixEncoding command
// var fixEncodingCmd = &cobra.Command{
// 	Use:   "fix-encoding",
// 	Short: "Cleans up the navigation hrefs by removing ids from the links",
// 	Long: `Goes and removes the ids from the navigation hrefs since they do not seem to play well with the Kindle
// 	library. It will not remove an id if 2 hrefs are referencing the same file.

// 	For example: epub-lint clean-nav-ids -f file-path
// 	will remove ids from the navigation links in the file so long as there are not 2 references to the same file.
// 	`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		validatefixEncodingFlags(filePath)
// 		fileText := utils.ReadInFileContents(filePath)

// 		var newText = RemoveIdsFromContentLinks(fileText, filePath)
// 		for file, instanceCount := range ExistingFileLinks {
// 			if instanceCount == 0 {
// 				continue
// 			}

// 			utils.WriteWarn(fmt.Sprintf(`"%s" has more than one reference to "%s". Consider restructuring the epub to fix up the TOC/nav.`, filePath, file))
// 		}

// 		if fileText == newText {
// 			return
// 		}

// 		utils.WriteFileContents(filePath, newText)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(fixEncodingCmd)

// 	fixEncodingCmd.Flags().StringVarP(&filePaths, "file-path", "f", "", "the file to remove the ids from")
// 	fixEncodingCmd.MarkFlagRequired("file-path")
// }

// func RemoveIdsFromContentLinks(text, filePath string) string {
// 	ExistingFileLinks = make(map[string]int)

// 	return contentElRegex.ReplaceAllStringFunc(text, removeHashTagsFromLinks)
// }

// func removeHashTagsFromLinks(part string) string {
// 	var groups = contentSrcAttributeRegex.FindStringSubmatch(part)
// 	if len(groups) != 5 {
// 		fmt.Printf(`possible problem with content tag: \"%s\", %v`, part, groups)

// 		return part
// 	}

// 	var file = groups[2]
// 	if count, exists := ExistingFileLinks[file]; exists {
// 		ExistingFileLinks[file] = count + 1
// 	} else {
// 		ExistingFileLinks[file] = 0
// 	}

// 	return strings.Replace(part, groups[0], groups[1]+file+groups[4], 1)
// }
