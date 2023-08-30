package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
	"github.com/spf13/cobra"
)

type ReplaceString struct {
	Search   *regexp.Regexp
	Replace  string
	Rational string
}

type ReplaceStringFunc struct {
	Search   *regexp.Regexp
	Replace  func(string) string
	Rational string
}

var regexEscapedPeriod = regexp.QuoteMeta(".")
var commonReplaceStrings = []ReplaceString{
	{
		Search:   regexp.MustCompile(fmt.Sprintf("(%[1]s ?){2}%[1]s", regexEscapedPeriod)),
		Replace:  "…",
		Rational: "Proper ellipses should be used where possible as it keeps things clean and consistent",
	},
	{
		Search:   regexp.MustCompile("(^|[^!])--([^>]|$)"),
		Replace:  "$1—$2",
		Rational: "Em dashes should be used where possible as it keeps things clean and consistent",
	},
	{
		Search:   regexp.MustCompile("(^|[^~])~([^~]|$)"),
		Replace:  "$1!$2",
		Rational: "Tildes should be replaced with an exclamation mark when they are by themselves as they seem interchangeable, though it could be another form of punctuation for drawing out the sound of the last letter used",
	},
	{
		Search:   regexp.MustCompile("(B|b)y the by"),
		Replace:  "${1}y the way",
		Rational: "'By the by' seems to be an improper translation of 'By the way', so we should auto-correct it to its proper English idiom",
	},
	{
		Search:   regexp.MustCompile("(S|s)neaked"),
		Replace:  "${1}nuck",
		Rational: "Use snuck instead of sneaked as it is the more commonly used version of the word nowadays",
	},
}

var commonReplaceStringFuncs = []ReplaceStringFunc{
	{
		Search: regexp.MustCompile("(A|a) bolt (o)(ut of the blue)"),
		Replace: func(part string) string {
			var firstLetter = "O"
			if part[0] == 'a' {
				firstLetter = "o"
			}

			return firstLetter + "ut of the blue"
		},
		Rational: "'a bolt out of the blue' seems to be an improper translation of 'out of the blue', so we should auto-correct it to its proper English idiom",
	},
	{
		Search: regexp.MustCompile("(L|l)ittle( wonder)"),
		Replace: func(part string) string {
			var firstLetter = "N"
			if part[0] == 'l' {
				firstLetter = "n"
			}

			return firstLetter + "o wonder"
		},
		Rational: "'little wonder' seems to be an improper translation of 'no wonder', so we should auto-correct it to its proper English idiom",
	},
}

var filePath string
var extraReplacesFilePath string

// replaceStringsCmd represents the replaceStrings command
var replaceStringsCmd = &cobra.Command{
	Use:   "replaceStrings",
	Short: "Replaces a common set of strings in a file",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint replaceStrings -f file-path -e extra-replace-file-path
	`,
	Run: func(cmd *cobra.Command, args []string) {
		filePathExists := utils.FileExists(filePath)
		if !filePathExists {
			utils.WriteError("file-path must exist")
		}

		fileText := utils.ReadInFileContents(filePath)

		var newText = CommonStringReplace(filePath)
		var numHits map[string]int

		if strings.Trim(extraReplacesFilePath, " ") != "" {

			extraReplacesFilePathExists := utils.FileExists(extraReplacesFilePath)
			if !extraReplacesFilePathExists {
				utils.WriteError("extra-replace-file-path must exist")
			}

			extraReplacesText := utils.ReadInFileContents(filePath)

			newText, numHits = ExtraStringReplace(newText, ParseTextReplacements(extraReplacesText))
			if len(numHits) == 0 {
				utils.WriteWarn("No values were listed as needing replacing")
			}

			for searchText, hits := range numHits {
				if hits == 0 {
					utils.WriteWarn(fmt.Sprintf("Did not find any replacements for `%s`", searchText))
				} else {
					utils.WriteInfo(fmt.Sprintf("`%s` was replaced %d time(s)", searchText, hits))
				}
			}
		}

		if fileText == newText {
			return
		}

		utils.WriteFileContents(filePath, newText)
	},
}

func init() {
	rootCmd.AddCommand(replaceStringsCmd)

	replaceStringsCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the path to the file to replace strings in")
	replaceStringsCmd.Flags().StringVarP(&extraReplacesFilePath, "extra-replace-text", "e", "", "the path to the file with extra strings to replace")
	replaceStringsCmd.MarkFlagRequired("file-path")
}

func CommonStringReplace(text string) string {
	var newText = text

	for _, replaceString := range commonReplaceStrings {
		newText = replaceString.Search.ReplaceAllString(newText, replaceString.Replace)
	}

	for _, replaceStringFunc := range commonReplaceStringFuncs {
		newText = replaceStringFunc.Search.ReplaceAllStringFunc(newText, replaceStringFunc.Replace)
	}

	return newText
}

func ExtraStringReplace(text string, extraFindAndReplaces map[string]string) (string, map[string]int) {
	var newText = text
	var numHits = make(map[string]int, len(extraFindAndReplaces))
	for search, replace := range extraFindAndReplaces {
		numHits[search] = strings.Count(newText, search)
		newText = strings.ReplaceAll(newText, search, replace)
	}

	return text, numHits
}

func ParseTextReplacements(text string) map[string]string {
	replaceValueToReplacement := make(map[string]string)

	var lines = strings.Split(text, "\n")
	var numLines = len(lines)
	if numLines <= 2 {
		return replaceValueToReplacement
	}

	// start after the markdown table header and divider lines
	var i = 2
	for i < numLines {
		var line = lines[i]
		i++
		var lineParts = strings.Split(line, "|")
		var numParts = len(lineParts)
		if numParts == 1 {
			continue
		} else if numParts != 4 {
			utils.WriteError(fmt.Sprintf("Could not parse \"%s\" because it does not have the proper amount of \"|\"s in it", line))
			continue
		}

		replaceValueToReplacement[strings.Trim(lineParts[1], " ")] = strings.Trim(lineParts[2], " ")
	}

	return replaceValueToReplacement
}
