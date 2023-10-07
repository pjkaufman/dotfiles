package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var cssPaths string
var runAll bool
var runBrokenLines bool
var runSectionBreak bool
var runPageBreak bool

// brokenLinesCmd represents the brokenLines command
var brokenLinesCmd = &cobra.Command{
	Use:   "broken-lines",
	Short: "Replaces a list of common strings and the extra strings specified in the extra file for the provide file(s)",
	Long: `Goes and replaces a common set of strings in a file as well as any extra instances that are specified.
	It will print out the successful extra replacements with the number of replacements made followed by warnings
	for any extra strings that it tried to find and replace values for, but did not find any instances to replace.
	
	For example: epub-lint replace-strings -f file-paths -e extra-replace-file-path
	will replace the common strings and the extra strings parsed out of the extra replace file 
	from the provided file(s)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		CheckForBrokenLines(log, fileHandler, filePaths, cssPaths, runAll, runBrokenLines, runSectionBreak, runPageBreak)
	},
}

func init() {
	rootCmd.AddCommand(brokenLinesCmd)

	brokenLinesCmd.Flags().StringVarP(&filePaths, "file-paths", "f", "", "the list of files to update in a comma separated list")
	brokenLinesCmd.Flags().StringVarP(&cssPaths, "css-paths", "c", "", "the list of css files which could be used for css additions")
	brokenLinesCmd.Flags().BoolVarP(&runAll, "run-all", "a", false, "whether to run all of the fixable suggestions")
	brokenLinesCmd.Flags().BoolVarP(&runBrokenLines, "run-broken-lines", "b", false, "whether to run the logic for getting broken line suggestions")
	brokenLinesCmd.Flags().BoolVarP(&runSectionBreak, "run-section-breaks", "s", false, "whether to run the logic for getting section break suggestions (must be used with css-paths)")
	brokenLinesCmd.Flags().BoolVarP(&runPageBreak, "run-page-breaks", "p", false, "whether to run the logic for getting page break suggestions (must be used with css-paths)")
	brokenLinesCmd.MarkFlagRequired("file-paths")
}

func CheckForBrokenLines(l logger.Logger, fileManager filehandler.FileManager, filePaths, cssPaths string, runAll, runBrokenLines, runSectionBreak, runPageBreak bool) {
	var files = strings.Split(filePaths, ",")
	var cssFiles = strings.Split(cssPaths, ",")
	validateBrokenLinesFlags(l, fileManager, files, cssFiles, runAll, runBrokenLines, runSectionBreak, runPageBreak)

	var addCssSectionIfMissing bool = false
	var addCssPageIfMissing bool = false
	var contextBreak string
	if runAll || runSectionBreak {
		contextBreak = l.GetInputString("What is the section break for the epub?:")

		if strings.Trim(contextBreak, " ") == "" {
			l.WriteError("Please provide a non-whitespace section break")
		}
	}

	for _, filePath := range files {
		fileText := fileManager.ReadInFileContents(filePath)

		var newText = fileText
		if runAll || runBrokenLines {
			var brokenLineFixSuggestions = linter.GetPotentiallyBrokenLines(newText)
			newText, _ = promptAboutSuggestions(l, brokenLineFixSuggestions, newText)
		}

		if runAll || runSectionBreak {
			var contextBreakSuggestions = linter.GetPotentialContextBreaks(newText, contextBreak)

			var contextBreakUpdated bool
			newText, contextBreakUpdated = promptAboutSuggestions(l, contextBreakSuggestions, newText)
			addCssSectionIfMissing = addCssSectionIfMissing || contextBreakUpdated
		}

		if runAll || runPageBreak {
			var pageBreakSuggestions = linter.GetPotentialPageBreaks(newText)

			var pageBreakUpdated bool
			newText, pageBreakUpdated = promptAboutSuggestions(l, pageBreakSuggestions, newText)
			addCssPageIfMissing = addCssPageIfMissing || pageBreakUpdated
		}

		// TODO: add oxford comma checks here

		if fileText == newText {
			continue
		}

		fileManager.WriteFileContents(filePath, newText)
	}

	handleCssChanges(l, fileManager, addCssSectionIfMissing, addCssPageIfMissing, cssFiles, contextBreak)
}

func validateBrokenLinesFlags(l logger.Logger, fileManager filehandler.FileManager, filePaths, cssPaths []string, runAll, runBrokenLines, runSectionBreak, runPageBreak bool) {
	if !runAll && !runBrokenLines && !runSectionBreak && !runPageBreak {
		l.WriteError("either run-all, run-broken-lines, run-section-breaks, or run-page-breaks must be specified")
	}

	for _, filePath := range filePaths {
		filePathExists := fileManager.FileExists(filePath)

		if !filePathExists {
			l.WriteError(fmt.Sprintf(`file-paths: "%s" must exist`, filePath))
		}
	}

	for _, cssPath := range cssPaths {
		filePathExists := fileManager.FileExists(cssPath)

		if !filePathExists {
			l.WriteError(fmt.Sprintf(`css-paths: "%s" must exist`, cssPath))
		}
	}

	if (runAll || runSectionBreak || runPageBreak) && (len(cssPaths) == 0 || len(cssPaths) == 1 && strings.Trim(cssPaths[0], " ") == "") {
		l.WriteError(`css-paths: must have a value when including handling section or page breaks`)
	}
}

func promptAboutSuggestions(l logger.Logger, suggestions map[string]string, fileText string) (string, bool) {
	var valueReplaced = false
	for original, suggestion := range suggestions {
		resp := l.GetInputString(fmt.Sprintf("Would you like to update \"%s\" to \"%s\"? (Y/N): ", strings.TrimLeft(original, "\n"), strings.TrimLeft(suggestion, "\n")))
		if strings.EqualFold(resp, "Y") {
			fileText = strings.Replace(fileText, original, suggestion, 1)
			valueReplaced = true
		}
	}

	return fileText, valueReplaced
}

func handleCssChanges(l logger.Logger, fileManager filehandler.FileManager, addCssSectionIfMissing, addCssPageIfMissing bool, cssFiles []string, contextBreak string) {
	if !addCssSectionIfMissing && !addCssPageIfMissing {
		return
	}

	var cssSelectionPrompt = "Please enter the number of the css file to append the css to:\n"
	for i, file := range cssFiles {
		cssSelectionPrompt += fmt.Sprintf("%d. %s\n", i, file)
	}

	var selectedCssFileIndex = l.GetInputInt(cssSelectionPrompt)
	if selectedCssFileIndex < 0 || selectedCssFileIndex >= len(cssFiles) {
		l.WriteError(fmt.Sprintf("Please select a valid css file value instead of \"%d\".", selectedCssFileIndex))
	}

	var cssFile = cssFiles[selectedCssFileIndex]
	var css = fileManager.ReadInFileContents(cssFile)
	var newCssText = css

	if addCssSectionIfMissing {
		newCssText = linter.AddCssSectionBreakIfMissing(newCssText, contextBreak)
	}

	if addCssPageIfMissing {
		newCssText = linter.AddCssPageBreakIfMissing(newCssText)
	}

	if newCssText != css {
		fileManager.WriteFileContents(cssFile, newCssText)
	}
}
