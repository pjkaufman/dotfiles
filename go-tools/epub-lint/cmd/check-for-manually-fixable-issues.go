package cmd

import (
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	cssPaths        string
	runAll          bool
	runBrokenLines  bool
	runSectionBreak bool
	runPageBreak    bool
	runOxfordCommas bool
)

// fixableCmd represents the fixable command
var fixableCmd = &cobra.Command{
	Use:   "fixable",
	Short: "Runs the specified fixable actions that require manual input to determine what to do.",
	Long: `Goes through all of the content files and runs the specified fixable actions on them asking
	for user input on each value found that matches the potential fix criteria.
	Potential things that can be fixed:
	- Broken paragraph endings
	- Section breaks being hardcoded instead of an hr
	- Page breaks being hardcoded instead of an hr
	- Oxford commas being missing before or's or and's
	
	For example: epub-lint fixable -f file-paths -c css-paths -a
	Will attempt to go through all of the potentially fixable issues in the specified files.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		CheckForBrokenLines(log, fileHandler, filePaths, cssPaths, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas)
	},
}

func init() {
	rootCmd.AddCommand(fixableCmd)

	fixableCmd.Flags().StringVarP(&filePaths, "file-paths", "f", "", "the list of files to update in a comma separated list")
	fixableCmd.Flags().StringVarP(&cssPaths, "css-paths", "c", "", "the list of css files which could be used for css additions")
	fixableCmd.Flags().BoolVarP(&runAll, "run-all", "a", false, "whether to run all of the fixable suggestions")
	fixableCmd.Flags().BoolVarP(&runBrokenLines, "run-broken-lines", "b", false, "whether to run the logic for getting broken line suggestions")
	fixableCmd.Flags().BoolVarP(&runSectionBreak, "run-section-breaks", "s", false, "whether to run the logic for getting section break suggestions (must be used with css-paths)")
	fixableCmd.Flags().BoolVarP(&runPageBreak, "run-page-breaks", "p", false, "whether to run the logic for getting page break suggestions (must be used with css-paths)")
	fixableCmd.Flags().BoolVarP(&runOxfordCommas, "run-oxford-commas", "o", false, "whether to run the logic for getting oxford comma suggestions")
	fixableCmd.MarkFlagRequired("file-paths")
}

func CheckForBrokenLines(l logger.Logger, fileManager filehandler.FileManager, filePaths, cssPaths string, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas bool) {
	var files = strings.Split(filePaths, ",")
	var cssFiles = strings.Split(cssPaths, ",")
	validateBrokenLinesFlags(l, fileManager, files, cssFiles, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas)

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
			var contextBreakSuggestions = linter.GetPotentialSectionBreaks(newText, contextBreak)

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

		if runAll || runOxfordCommas {
			var oxfordCommaSuggestions = linter.GetPotentialMissingOxfordCommas(newText)
			newText, _ = promptAboutSuggestions(l, oxfordCommaSuggestions, newText)
		}

		if fileText == newText {
			continue
		}

		fileManager.WriteFileContents(filePath, newText)
	}

	handleCssChanges(l, fileManager, addCssSectionIfMissing, addCssPageIfMissing, cssFiles, contextBreak)
}

func validateBrokenLinesFlags(l logger.Logger, fileManager filehandler.FileManager, filePaths, cssPaths []string, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas bool) {
	if !runAll && !runBrokenLines && !runSectionBreak && !runPageBreak && !runOxfordCommas {
		l.WriteError("either run-all, run-broken-lines, run-section-breaks, run-page-breaks, or run-oxford-commas must be specified")
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
