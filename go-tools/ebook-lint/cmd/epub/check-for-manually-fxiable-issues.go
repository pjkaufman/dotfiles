package epub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/ebook-lint/linter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	runAll          bool
	runBrokenLines  bool
	runSectionBreak bool
	runPageBreak    bool
	runOxfordCommas bool
	runAlthoughBut  bool
)

const (
	OneRunBoolArgMustBeEnabled   = "either run-all, run-broken-lines, run-section-breaks, run-page-breaks, run-oxford-commas, or run-although-but must be specified"
	CssPathsEmptyWhenArgIsNeeded = "css-paths must have a value when including handling section or page breaks"
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
	- Possible instances of sentences that have although ..., but in them
	
	For example: ebook-lint epub fixable -f file-paths -c css-paths -a
	Will attempt to go through all of the potentially fixable issues in the specified files.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateManuallyFixableFlags(epubFile, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas, runAlthoughBut)
		if err != nil {
			logger.WriteError(err.Error())
		}

		filehandler.FileMustExist(epubFile, "epub-file")

		logger.WriteInfo("Started showing manually fixable issues...\n")

		var epubFolder = filehandler.GetFileFolder(epubFile)
		var dest = filehandler.JoinPath(epubFolder, "epub")
		filehandler.UnzipRunOperationAndRezip(epubFile, dest, func() {
			opfFolder, epubInfo := getEpubInfo(dest, epubFile)
			validateFilesExist(opfFolder, epubInfo.HtmlFiles)
			validateFilesExist(opfFolder, epubInfo.CssFiles)

			var addCssSectionIfMissing bool = false
			var addCssPageIfMissing bool = false
			var contextBreak string
			if runAll || runSectionBreak {
				contextBreak = logger.GetInputString("What is the section break for the epub?:")

				if strings.TrimSpace(contextBreak) == "" {
					logger.WriteError("Please provide a non-whitespace section break")
				}
			}

			var cssFiles = make([]string, len(epubInfo.CssFiles))
			var i = 0
			for cssFile := range epubInfo.CssFiles {
				cssFiles[i] = cssFile
				i++
			}

			if (runAll || runSectionBreak || runPageBreak) && len(cssFiles) == 0 {
				logger.WriteError(CssPathsEmptyWhenArgIsNeeded)
			}

			for file := range epubInfo.HtmlFiles {
				var filePath = getFilePath(opfFolder, file)
				fileText := filehandler.ReadInFileContents(filePath)

				var newText = fileText
				if runAll || runBrokenLines {
					var brokenLineFixSuggestions = linter.GetPotentiallyBrokenLines(newText)
					newText, _ = promptAboutSuggestions("Potential Broken Lines", brokenLineFixSuggestions, newText)
				}

				if runAll || runSectionBreak {
					var contextBreakSuggestions = linter.GetPotentialSectionBreaks(newText, contextBreak)

					var contextBreakUpdated bool
					newText, contextBreakUpdated = promptAboutSuggestions("Potential Section Breaks", contextBreakSuggestions, newText)
					addCssSectionIfMissing = addCssSectionIfMissing || contextBreakUpdated
				}

				if runAll || runPageBreak {
					var pageBreakSuggestions = linter.GetPotentialPageBreaks(newText)

					var pageBreakUpdated bool
					newText, pageBreakUpdated = promptAboutSuggestions("Potential Page Breaks", pageBreakSuggestions, newText)
					addCssPageIfMissing = addCssPageIfMissing || pageBreakUpdated
				}

				if runAll || runOxfordCommas {
					var oxfordCommaSuggestions = linter.GetPotentialMissingOxfordCommas(newText)
					newText, _ = promptAboutSuggestions("Potential Missing Oxford Commas", oxfordCommaSuggestions, newText)
				}

				if runAll || runAlthoughBut {
					var althoughButSuggestions = linter.GetPotentialAlthoughButInstances(newText)
					newText, _ = promptAboutSuggestions("Potential Although But Instances", althoughButSuggestions, newText)
				}

				if fileText == newText {
					continue
				}

				filehandler.WriteFileContents(filePath, newText)
			}

			handleCssChanges(addCssSectionIfMissing, addCssPageIfMissing, opfFolder, cssFiles, contextBreak)
		})

		logger.WriteInfo("\nFinished showing manually fixable issues...")
	},
}

func init() {
	EpubCmd.AddCommand(fixableCmd)

	fixableCmd.Flags().BoolVarP(&runAll, "run-all", "a", false, "whether to run all of the fixable suggestions")
	fixableCmd.Flags().BoolVarP(&runBrokenLines, "run-broken-lines", "b", false, "whether to run the logic for getting broken line suggestions")
	fixableCmd.Flags().BoolVarP(&runSectionBreak, "run-section-breaks", "s", false, "whether to run the logic for getting section break suggestions (must be used with css-paths)")
	fixableCmd.Flags().BoolVarP(&runPageBreak, "run-page-breaks", "p", false, "whether to run the logic for getting page break suggestions (must be used with css-paths)")
	fixableCmd.Flags().BoolVarP(&runOxfordCommas, "run-oxford-commas", "o", false, "whether to run the logic for getting oxford comma suggestions")
	fixableCmd.Flags().BoolVarP(&runAlthoughBut, "run-although-but", "n", false, "whether to run the logic for getting although but suggestions")
	fixableCmd.Flags().StringVarP(&epubFile, "epub-file", "f", "", "the epub file to find manually fixable issues in")
	fixableCmd.MarkFlagRequired("epub-file")
}

func ValidateManuallyFixableFlags(epubPath string, runAll, runBrokenLines, runSectionBreak, runPageBreak, runOxfordCommas, runAlthoughBut bool) error {
	err := validateCommonEpubFlags(epubPath)
	if err != nil {
		return err
	}

	if !runAll && !runBrokenLines && !runSectionBreak && !runPageBreak && !runOxfordCommas && !runAlthoughBut {
		return errors.New(OneRunBoolArgMustBeEnabled)
	}

	return nil
}

func promptAboutSuggestions(suggestionsTitle string, suggestions map[string]string, fileText string) (string, bool) {
	var valueReplaced = false
	var newText = fileText

	if len(suggestions) == 0 {
		return newText, valueReplaced
	}

	logger.WriteInfo(cliLineSeparator)
	logger.WriteInfo(suggestionsTitle + ":")
	logger.WriteInfo(cliLineSeparator + "\n")

	for original, suggestion := range suggestions {
		resp := logger.GetInputString(fmt.Sprintf("Would you like to update \"%s\" to \"%s\"? (Y/N): ", strings.TrimLeft(original, "\n"), strings.TrimLeft(suggestion, "\n")))
		if strings.EqualFold(resp, "Y") {
			newText = strings.Replace(newText, original, suggestion, 1)
			valueReplaced = true
		}

		logger.WriteInfo("")
	}

	return newText, valueReplaced
}

func handleCssChanges(addCssSectionIfMissing, addCssPageIfMissing bool, opfFolder string, cssFiles []string, contextBreak string) {
	if !addCssSectionIfMissing && !addCssPageIfMissing {
		return
	}

	var cssSelectionPrompt = "Please enter the number of the css file to append the css to:\n"
	for i, file := range cssFiles {
		cssSelectionPrompt += fmt.Sprintf("%d. %s\n", i, file)
	}

	var selectedCssFileIndex = logger.GetInputInt(cssSelectionPrompt)
	if selectedCssFileIndex < 0 || selectedCssFileIndex >= len(cssFiles) {
		logger.WriteError(fmt.Sprintf("Please select a valid css file value instead of \"%d\".", selectedCssFileIndex))
	}

	var cssFile = cssFiles[selectedCssFileIndex]
	var cssFilePath = filehandler.JoinPath(opfFolder, cssFile)
	var css = filehandler.ReadInFileContents(cssFilePath)
	var newCssText = css

	if addCssSectionIfMissing {
		newCssText = linter.AddCssSectionBreakIfMissing(newCssText, contextBreak)
	}

	if addCssPageIfMissing {
		newCssText = linter.AddCssPageBreakIfMissing(newCssText)
	}

	if newCssText != css {
		filehandler.WriteFileContents(cssFilePath, newCssText)
	}
}
