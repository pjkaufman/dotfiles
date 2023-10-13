package epub

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
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
	LintDirArgEmpty = "directory must have a non-whitespace value"
	LangArgEmpty    = "lang must have a non-whitespace value"
)

// compressAndLintCmd represents the compressAndLint command
var compressAndLintCmd = &cobra.Command{
	Use:   "compress-and-lint",
	Short: "Compresses and lints all of the epub files in the specified directory even compressing images using imgp if that option is specified.",
	Long: `Gets all of the .epub files in the specified directory.
Then it lints each epub separately making sure to compress the images if specified.
Some of the things that the linting includes:
- Replacing a list of common strings
- Removing links from the nav and/or the ncx file
- Adds absolute page numbers if present in file's ids
- Adds language encoding specified if it is not present already (default is "en")
- Sets encoding on content files to utf-8 to prevent errors in some readers
	
For example: epub-lint epub compress-and-lint -d folder -i
will get all .epub files located in folder making sure to lint them and compress
any images listed in the opf file. It will print a summary of the before and after
sizes of the epubs for each file and then again at the very end of the run.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		err := ValidateCompressAndLintFlags(lintDir, lang)
		if err != nil {
			logger.WriteError(err.Error())
		}

		logger.WriteInfo("Starting compression and linting for each epub\n")

		epubs := filehandler.MustGetAllFilesWithExtInASpecificFolder(lintDir, ".epub")

		var totalBeforeFileSize, totalAfterFileSize float64
		for _, epub := range epubs {
			logger.WriteInfo(fmt.Sprintf("starting epub compressing for %s...", epub))

			LintEpub(lintDir, epub, runCompressImages)

			var originalFile = epub + ".original"
			var newKbSize = filehandler.MustGetFileSize(epub)
			var oldKbSize = filehandler.MustGetFileSize(originalFile)

			logger.WriteInfo("\n" + cliLineSeparator)
			logger.WriteInfo("Before:")
			logger.WriteInfo(fmt.Sprintf("%s %s", originalFile, kbSizeToString(oldKbSize)))
			logger.WriteInfo("After:")
			logger.WriteInfo(fmt.Sprintf("%s %s", epub, kbSizeToString(newKbSize)))
			logger.WriteInfo(cliLineSeparator + "\n")

			totalBeforeFileSize += oldKbSize
			totalAfterFileSize += newKbSize
		}

		logger.WriteInfo("\n" + cliLineSeparator)
		logger.WriteInfo("Before:")
		logger.WriteInfo(kbSizeToString(totalBeforeFileSize))
		logger.WriteInfo("After:")
		logger.WriteInfo(kbSizeToString(totalAfterFileSize))
		logger.WriteInfo(cliLineSeparator + "\n")

		logger.WriteInfo("Finished compression and linting")
	},
}

func init() {
	EpubCmd.AddCommand(compressAndLintCmd)

	compressAndLintCmd.Flags().StringVarP(&lintDir, "directory", "d", ".", "the location to run the epub lint logic")
	compressAndLintCmd.Flags().StringVarP(&lang, "lang", "l", "en", "the language to add to the xhtml, htm, or html files if the lang is not already specified")
	compressAndLintCmd.Flags().BoolVarP(&runCompressImages, "compress-images", "i", false, "whether or not to also compress images which requires imgp to be installed")
}

func LintEpub(lintDir, epub string, runCompressImages bool) {
	var src = filehandler.JoinPath(lintDir, epub)
	var dest = filehandler.JoinPath(lintDir, "epub")

	filehandler.UnzipRunOperationAndRezip(src, dest, func() {
		opfFolder, epubInfo := getEpubInfo(dest, epub)

		validateFilesExist(opfFolder, epubInfo.HtmlFiles)
		validateFilesExist(opfFolder, epubInfo.ImagesFiles)
		validateFilesExist(opfFolder, epubInfo.OtherFiles)

		// fix up all xhtml files first
		for file := range epubInfo.HtmlFiles {
			var filePath = getFilePath(opfFolder, file)
			fileText := filehandler.ReadInFileContents(filePath)
			var newText = linter.EnsureEncodingIsPresent(fileText)
			newText = linter.CommonStringReplace(newText)

			// TODO: remove images links that do not exist in the manifest
			newText = linter.EnsureLanguageIsSet(newText, lang)
			epubInfo.PageIds = linter.GetPageIdsForFile(newText, file, epubInfo.PageIds)

			if fileText == newText {
				continue
			}

			filehandler.WriteFileContents(filePath, newText)
		}

		updateNavFile(opfFolder, epubInfo.NavFile, epubInfo.PageIds)
		updateNcxFile(opfFolder, epubInfo.NcxFile, epubInfo.PageIds)
		//TODO: get all files in the repo and prompt the user whether they want to delete them

		if runCompressImages {
			compressImages(lintDir, opfFolder, epubInfo.ImagesFiles)
		}

		// TODO: cleanup TOC file's links
		/** Sample ToC
				<?xml version="1.0" encoding="utf-8"?>
		<!DOCTYPE html>

		<html xmlns:epub="http://www.idpf.org/2007/ops" xmlns="http://www.w3.org/1999/xhtml" lang="en" xml:lang="en">
		<head>
		<meta http-equiv="default-style" content="text/html; charset=UTF-8"/>
		<title>Death March to the Parallel World Rhapsody, Vol. 19</title>
		<link rel="stylesheet" href="../Styles/stylesheet.css" type="text/css"/>



		</head>
		<body>
		<h1 class="toc-title"><a id="page-iii"></a><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.1.1">Contents</span></h1>
		<p class="toc-front" id="cover"><a href="../Text/cover.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.2.1">Cover</span></a></p>
		<p class="toc-front" id="insert001"><a href="../Text/insert001.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.3.1">Insert</span></a></p>
		<p class="toc-front" id="titlepage"><a href="../Text/titlepage.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.4.1">Title Page</span></a></p>
		<p class="toc-front" id="toc-copyright"><a href="../Text/copyright.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.5.1">Copyright</span></a></p>
		<p class="toc-chapter1" id="toc-chapter001"><a href="../Text/chapter001.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.6.1">Vice-Minister of Tourism</span></a></p>
		<p class="toc-chapter" id="toc-chapter006"><a href="../Text/chapter006.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.7.1">Return to Labyrinth City</span></a></p>
		<p class="toc-chapter" id="toc-chapter010"><a href="../Text/chapter010.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.8.1">The Barren Territory</span></a></p>
		<p class="toc-chapter" id="toc-chapter013"><a href="../Text/chapter013.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.9.1">An Unexpected Reunion</span></a></p>
		<p class="toc-chapter" id="toc-chapter016"><a href="../Text/chapter016.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.10.1">Journey to Yowork Kingdom</span></a></p>
		<p class="toc-chapter" id="toc-chapter017"><a href="../Text/chapter017.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.11.1">Interlude: A Skirmish</span></a></p>
		<p class="toc-chapter" id="toc-chapter018"><a href="../Text/chapter018.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.12.1">Sacrificial Labyrinth</span></a></p>
		<p class="toc-chapter" id="toc-chapter019"><a href="../Text/chapter019.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.13.1">Requiem</span></a></p>
		<p class="toc-chapter" id="toc-chapter022"><a href="../Text/chapter022.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.14.1">Uprising</span></a></p>
		<p class="toc-chapter" id="toc-chapter023"><a href="../Text/chapter023.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.15.1">A Bitter Battlefield</span></a></p>
		<p class="toc-chapter" id="toc-chapter024"><a href="../Text/chapter024.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.16.1">Epilogue</span></a></p>
		<p class="toc-chapter" id="toc-chapter027"><a href="../Text/chapter027.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.17.1">EX: Karina and Zena's Big Adventure</span></a></p>
		<p class="toc-chapter1" id="toc-appendix001"><a href="../Text/appendix001.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.18.1">Afterword</span></a></p>
		<p class="toc-chapter" id="newsletter1"><a href="../Text/newsletterSignup.xhtml"><span xmlns="http://www.w3.org/1999/xhtml" class="koboSpan" id="kobo.19.1">Yen Newsletter</span></a></p>
		</body>
		</html>
		*/
	})

	// TODO: print out the size of all of the before and after
}

func ValidateCompressAndLintFlags(lintDir, lang string) error {
	if strings.TrimSpace(lintDir) == "" {
		return errors.New(LintDirArgEmpty)
	}

	if strings.TrimSpace(lang) == "" {
		return errors.New(LangArgEmpty)
	}

	return nil
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

var kilobytesInAMegabyte float64 = 1024
var kilobytesInAGigabyte float64 = 1000000

func kbSizeToString(size float64) string {
	if size > kilobytesInAGigabyte {
		return fmt.Sprintf("%.2f GB", size/kilobytesInAGigabyte)
	} else if size > kilobytesInAMegabyte {
		return fmt.Sprintf("%.2f MB", size/kilobytesInAMegabyte)
	}

	return fmt.Sprintf("%.2f KB", size)
}
