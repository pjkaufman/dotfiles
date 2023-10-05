package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	FilePathArgEmpty        = "file-path must have a non-whitespace value"
	FilePathNotMarkdownFile = "file-path must be a markdown file"
	emptyColumnContent      = "&nbsp;&nbsp;&nbsp;&nbsp;"
	closeMetadata           = "</div><br><br>"
)

var filePath string
var otherTitleRegex = regexp.MustCompile(`^(<h1.*)\((.*)\)<(.*)`)

type SongMetadata struct {
	Melody         string `yaml:"melody"`
	SongKey        string `yaml:"key"`
	Authors        string `yaml:"authors"`
	InChurch       string `yaml:"in-church"`
	VerseReference string `yaml:"verse"`
	BookLocation   string `yaml:"location"`
	Copyright      string `yaml:"copyright"`
}

// mdToHtmlCmd represents the mdToHtml command
var mdToHtmlCmd = &cobra.Command{
	Use:   "md-to-html",
	Short: "Converts the provided Markdown file's contents to html and puts it back out to std out",
	Long: `Takes the contents of a Markdown file and converts it to html based on the YAML frontmatter contents
	
	For example: song-converter md-to-html -f file-path
	converts the Markdown file to an html file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var log = logger.NewLoggerHandler()
		var fileHandler = filehandler.NewFileHandler(log)
		MdToHtml(log, fileHandler, filePath)
	},
}

func init() {
	rootCmd.AddCommand(mdToHtmlCmd)

	mdToHtmlCmd.Flags().StringVarP(&filePath, "file-path", "f", "", "the Markdown file to convert to html")
	mdToHtmlCmd.MarkFlagRequired("file-path")
}

func MdToHtml(l logger.Logger, fileManager filehandler.FileManager, filePath string) string {
	validateMdToHtmlFlags(l, fileManager, filePath)

	contents := fileManager.ReadInFileContents(filePath)

	var metadata SongMetadata
	mdContent, err := frontmatter.Parse(strings.NewReader(contents), &metadata)
	if err != nil {
		l.WriteError(fmt.Sprintf(`There was an error getting the frontmatter for file '%s': %s`, filePath, err))
	}

	var metadataHtml = buildMetadataDiv(&metadata)
	html := mdToHTML([]byte(mdContent))
	html = otherTitleRegex.ReplaceAllString(html, `${1}<span class="other-title">(${2})</span><${3}`)

	// just in case we encounter this scenario where non-breaking space is encoded as its unicode value
	html = strings.ReplaceAll(html, "\u00a0\u00a0\n", "<br>\n")
	html = strings.ReplaceAll(html, "\\&", "&")
	html = strings.Replace(html, "</h1>\n", "</h1>\n"+metadataHtml, 1)
	html = strings.ReplaceAll(html, "\n\n", "\n")

	return fmt.Sprintf("<div class=\"keep-together\">\n%s</div>\n<br>", html)
}

func validateMdToHtmlFlags(l logger.Logger, fileManager filehandler.FileManager, filePath string) {
	if strings.Trim(filePath, " ") == "" {
		l.WriteError(FilePathArgEmpty)
	}

	if !strings.HasSuffix(filePath, ".md") {
		l.WriteError(FilePathNotMarkdownFile)
	}

	if !fileManager.FileExists(filePath) {
		l.WriteError(fmt.Sprintf(`file-path: "%s" must exist`, filePath))
	}
}

func mdToHTML(md []byte) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}

func buildMetadataDiv(metadata *SongMetadata) string {
	if metadata == nil {
		return ""
	}

	var metadataCount = 0
	var row1 = 0
	var row2 = 0

	metadataCount, row2 = updateCountsIfMetatdataExists(metadata.Melody, metadataCount, row2)
	metadataCount, row2 = updateCountsIfMetatdataExists(metadata.VerseReference, metadataCount, row2)
	metadataCount, row1 = updateCountsIfMetatdataExists(metadata.Authors, metadataCount, row1)
	metadataCount, row1 = updateCountsIfMetatdataExists(metadata.SongKey, metadataCount, row1)
	metadataCount, row1 = updateCountsIfMetatdataExists(metadata.BookLocation, metadataCount, row1)

	if metadataCount == 0 {
		return ""
	}

	var metadataHtml = strings.Builder{}
	metadataHtml.WriteString("<div>")

	var (
		addRowEntry = func(value, class, nonEmptyValue string) {
			metadataHtml.WriteString(fmt.Sprintf("<div><div class=\"%s\">", class))

			if strings.Trim(value, "") != "" {
				metadataHtml.WriteString(nonEmptyValue)
			} else {
				metadataHtml.WriteString(emptyColumnContent)
			}

			metadataHtml.WriteString("</div></div>")
		}
		addBoldRowEntry = func(value, class string) {
			addRowEntry(value, class, fmt.Sprintf("<b>%s</b>", value))
		}
		addRegularRowEntry = func(value, class string) {
			addRowEntry(value, class, value)
		}
	)

	if row1 != 0 {
		if row2 != 0 {
			metadataHtml.WriteString("<div class=\"metadata row-padding\">")
		} else {
			metadataHtml.WriteString("<div class=\"metadata\">")
		}

		if strings.EqualFold(metadata.InChurch, "Y") {
			addBoldRowEntry(metadata.Authors, "author")
		} else {
			addRegularRowEntry(metadata.Authors, "author")
		}

		addBoldRowEntry(metadata.SongKey, "key")
		addRegularRowEntry(metadata.BookLocation, "location")

		metadataHtml.WriteString("</div>")
	}

	if row2 == 0 {
		metadataHtml.WriteString(closeMetadata)

		return metadataHtml.String()
	}

	metadataHtml.WriteString("<div class=\"metadata\">")
	if row2 == 1 && metadata.Melody != "" {
		addBoldRowEntry(metadata.Melody, "melody-75")
	} else {
		addBoldRowEntry(metadata.Melody, "melody")
		addRegularRowEntry(metadata.VerseReference, "verse")
	}

	metadataHtml.WriteString("</div>")
	metadataHtml.WriteString(closeMetadata)

	return metadataHtml.String()
}

func updateCountsIfMetatdataExists(value string, metadataElements, rowElements int) (int, int) {
	if (strings.Trim(value, "")) == "" {
		return metadataElements, rowElements
	}

	return metadataElements + 1, rowElements + 1
}
