package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

var filePath string

var contentElRegex = regexp.MustCompile("(<content.* src=[\"'])([^\"']*)([\"'][^>\n]*>)")
var contentSrcAttributeRegex = regexp.MustCompile("(src=[\"'])([^\"'#]*)(#[^\"']+)?([\"'])")
var ExistingFileLinks map[string]int

// cleanNavIdsCmd represents the cleanNavIds command
var cleanNavIdsCmd = &cobra.Command{
	Use:   "clean-nav-ids",
	Short: "Cleans up the navigation hrefs by removing ids from the links",
	Long: `Goes and removes the ids from the navigation hrefs since they do not seem to play well with the Kindle
	library. It will not remove an id if 2 hrefs are referencing the same file.
	
	For example: epub-lint clean-nav-ids -f file-path
	will remove ids from the navigation links in the file so long as there are not 2 references to the same file.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		validateCleanNavIdsFlags(filePath)
		fileText := utils.ReadInFileContents(filePath)

		var newText = RemoveIdsFromContentLinks(fileText, filePath)
		for file, instanceCount := range ExistingFileLinks {
			if instanceCount == 0 {
				continue
			}

			utils.WriteWarn(fmt.Sprintf(`"%s" has more than one reference to "%s". Consider restructuring the epub to fix up the TOC/nav.`, filePath, file))
		}

		if fileText == newText {
			return
		}

		utils.WriteFileContents(filePath, newText)
	},
}

func init() {
	rootCmd.AddCommand(cleanNavIdsCmd)

	cleanNavIdsCmd.Flags().StringVarP(&filePaths, "file-path", "f", "", "the file to remove the ids from")
	cleanNavIdsCmd.MarkFlagRequired("file-path")
}

func RemoveIdsFromContentLinks(text, filePath string) string {
	ExistingFileLinks = make(map[string]int)

	return contentElRegex.ReplaceAllStringFunc(text, removeHashTagsFromLinks)
}

func removeHashTagsFromLinks(part string) string {
	var groups = contentSrcAttributeRegex.FindStringSubmatch(part)
	if len(groups) != 5 {
		fmt.Printf(`possible problem with content tag: \"%s\", %v`, part, groups)

		return part
	}

	var file = groups[2]
	if count, exists := ExistingFileLinks[file]; exists {
		ExistingFileLinks[file] = count + 1
	} else {
		ExistingFileLinks[file] = 0
	}

	return strings.Replace(part, groups[0], groups[1]+file+groups[4], 1)
}

// def remove_id_if_exists_and_check_if_file_already_in_list(match):
//     global ignore_file
//     global existing_files

//     match1 = match.group(1)
//     match2 = match.group(2)
//     match3 = match.group(3)

//     if ignore_file is True:
//         return "{0}{1}{2}".format(match1, match2, match3)

//     hashtag_index = match2.find("#")
//     if hashtag_index != -1:
//         match2 = match2[hashtag_index : len(match2)]

//     if match2 in existing_files:
//         ignore_file = True

//     return "{0}{1}{2}".format(match1, match2, match3)

// # <content src="titlepage.xhtml"/>
// def remove_ids_from_hrefs(text):
//     global ignore_file
//     ignore_file = False

//     global existing_files
//     existing_files = {}

//     updated_text = re.sub(
//         r"(<content.* src=[\"'])([^\"']*)([\"'][^>\n]*>)",
//         remove_id_if_exists_and_check_if_file_already_in_list,
//         text,
//     )

//     if ignore_file is True:
//         print(
//             "The file has more than one reference to the same file. Consider restructuring the epub to fix up the TOC."
//         )

//         return text

//     return updated_text

func validateCleanNavIdsFlags(filePath string) {
	filePathExists := utils.FileExists(filePath)

	if !filePathExists {
		utils.WriteError("file-path must exist")
	}
}
