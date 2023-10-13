package epub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	archivehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/archive-handler"
	commandhandler "github.com/pjkaufman/dotfiles/go-tools/pkg/command-handler"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

var (
	lintDir           string
	lang              string
	runCompressImages bool
)

const (
	FilePathArgEmpty  = "file-path must have a non-whitespace value"
	FilePathArgNonOpf = "file-path must be an opf file"
	LangArgEmpty      = "lang must have a non-whitespace value"
)

const imgComperssionProgramName = "imgp"

var compressionParams = []string{"-x", "800x800", "-e", "-O", "-q", "40", "-m", "-d", "-w"}
var compressableImageExts = []string{"png", "jpg", "jpeg"}

// compressAndLintCmd represents the compressAndLint command
var compressAndLintCmd = &cobra.Command{
	Use:   "compress-and-lint",
	Short: "Takes the opf file of an epub and uses that to lint the files within it",
	Long: `Goes and replaces a common set of strings a file as well as any extra instances that are specified
	
	For example: epub-lint compress-and-lint
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// err := ValidateCompressAndLintFlags(lang)
		// if err != nil {
		// 	logger.WriteError(err.Error())
		// }

		epubs := filehandler.MustGetAllFilesWithExtInASpecificFolder(lintDir, ".epub")
		logger.WriteInfo(strings.Join(epubs, "\n"))
		var epub = epubs[0]
		var src = filehandler.JoinPath(lintDir, epub)
		var dest = filehandler.JoinPath(lintDir, "epub")

		err := archivehandler.UnzipRezipAndRunOperation(src, dest, func() error {
			logger.WriteInfo("Success?")

			opfFiles := filehandler.MustGetAllFilesWithExtInASpecificFolderAndSubFolders(dest, ".opf")
			if len(opfFiles) < 1 {
				logger.WriteError("did not find opf file...")
			}

			var filePath = opfFiles[0]
			opfText := filehandler.ReadInFileContents(filePath)

			epubInfo, err := linter.ParseOpfFile(opfText)
			if err != nil {
				logger.WriteError(fmt.Sprintf("Failed to parse \"%s\": %s", filePath, err))
			}

			var opfFolderString = filehandler.GetFileFolder(filePath)

			validateFilesExist(opfFolderString, epubInfo.HtmlFiles)
			validateFilesExist(opfFolderString, epubInfo.ImagesFiles)
			validateFilesExist(opfFolderString, epubInfo.OtherFiles)

			// fix up all xhtml files first
			for file := range epubInfo.HtmlFiles {
				var filePath = getFilePath(opfFolderString, file)
				fileText := filehandler.ReadInFileContents(filePath)
				var newText = linter.EnsureEncodingIsPresent(fileText)
				newText = linter.CommonStringReplace(newText)

				// TODO: remove images links that do not exist in the manifest
				// TODO: remove files that exist, but are not in the manifest
				newText = linter.EnsureLanguageIsSet(newText, lang)
				epubInfo.PageIds = linter.GetPageIdsForFile(newText, file, epubInfo.PageIds)

				if fileText == newText {
					continue
				}

				filehandler.WriteFileContents(filePath, newText)
			}

			updateNavFile(opfFolderString, epubInfo.NavFile, epubInfo.PageIds)
			updateNcxFile(opfFolderString, epubInfo.NcxFile, epubInfo.PageIds)
			//TODO: get all files in the repo and prompt the user whether they want to delete them

			if runCompressImages {
				compressImages(lintDir, opfFolderString, epubInfo.ImagesFiles)
			}

			// TODO: cleanup TOC file's links
			return nil
		})
		if err != nil {
			logger.WriteError(err.Error())
		}
	},
}

func init() {
	epubCmd.AddCommand(compressAndLintCmd)

	compressAndLintCmd.Flags().StringVarP(&lintDir, "directory", "d", ".", "the location to run the epub lint logic")
	compressAndLintCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	compressAndLintCmd.Flags().BoolVarP(&runCompressImages, "compress-images", "i", false, "whether or not to also compress images which requires imgp to be installed")
}

func ValidateCompressAndLintFlags(lang string) error {
	if strings.TrimSpace(lang) == "" {
		return errors.New(LangArgEmpty)
	}

	return nil
}

func compressImages(destFolder, opfFolder string, images map[string]struct{}) {
	for imagePath := range images {
		if !isCompressableImage(imagePath) {
			continue
		}

		var params = fmt.Sprintf("%s %s %s", imgComperssionProgramName, strings.Join(compressionParams, " "), filehandler.JoinPath(opfFolder, imagePath))
		fmt.Println(commandhandler.MustGetCommandOutput("bash", fmt.Sprintf(`failed to compress "%s"`, imagePath), []string{"-c", params}...))

		// TODO: see if I can figure out why the following does not work
		// var params = append(compressionParams, "\""+filehandler.JoinPath(opfFolder, imagePath)+"\"")
		// fmt.Println(commandhandler.MustGetCommandOutput(imgComperssionProgramName, fmt.Sprintf(`failed to compress "%s"`, imagePath), params...))
	}
}

func isCompressableImage(imagePath string) bool {
	for _, ext := range compressableImageExts {
		if strings.HasSuffix(strings.ToLower(imagePath), ext) {
			return true
		}
	}

	return false
}

func validateFilesExist(opfFolder string, files map[string]struct{}) {
	for file := range files {
		var filePath = getFilePath(opfFolder, file)

		if !filehandler.FileExists(filePath) {
			logger.WriteError(fmt.Sprintf(`file from manifest not found: "%s" must exist`, filePath))
		}
	}
}

func updateNcxFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := filehandler.ReadInFileContents(filePath)

	newText, err := linter.CleanupNavMap(fileText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNcxFile(newText, pageIds)

	if fileText == newText {
		return
	}

	filehandler.WriteFileContents(filePath, newText)
}

func updateNavFile(opfFolder, file string, pageIds []linter.PageIdInfo) {
	if file == "" {
		return
	}

	var filePath = getFilePath(opfFolder, file)
	fileText := filehandler.ReadInFileContents(filePath)

	newText, err := linter.RemoveIdsFromNav(fileText)
	if err != nil {
		logger.WriteError(fmt.Sprintf("%s: %v", filePath, err))
	}

	newText = linter.AddPageListToNavFile(newText, pageIds)

	if fileText == newText {
		return
	}

	filehandler.WriteFileContents(filePath, newText)
}

func getFilePath(opfFolder, file string) string {
	return filehandler.JoinPath(opfFolder, file)
}
