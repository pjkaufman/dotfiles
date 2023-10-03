package cmd

import (
	"fmt"
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
	closeMetadata           = "</div><br/><br/>"
)

var filePath string

type SongMetadata struct {
	Melody         string `yaml:"melody"`
	SongKey        string `yaml:"key"`
	Authors        string `yaml:"authors"`
	InChurch       string `yaml:"in-church"`
	VerseReference string `yaml:"verse"`
	BookLocation   string `yaml:"location"`
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

	return fmt.Sprintf("%s\n%s", metadataHtml, html)
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

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// if printAst {
	// 	fmt.Print("--- AST tree:\n")
	// 	ast.Print(os.Stdout, doc)
	// 	fmt.Print("\n")
	// }

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
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
		addRowEntry = func(value, class, nonEmptyValy string) {
			metadataHtml.WriteString(fmt.Sprintf("<div><div class=\"%s\">", class))

			if strings.Trim(value, "") != "" {
				metadataHtml.WriteString(value)
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
		metadataHtml.WriteString("\n")
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
		addBoldRowEntry(metadata.VerseReference, "verse")
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

/**
fileName=$(basename "$f" .md)
    yaml=$(parse_yaml "./stagingGround/$f")
    pandoc "./stagingGround/$f" -o "./html/build/$fileName.html"

    if [ -n "$yaml" ]; then
      metadata=$(build_metadata_div $yaml)
      sed -i "/<\/h1>/a ${metadata@Q}" "./html/build/$fileName.html"
      sed -i "s/\$'</</" "./html/build/$fileName.html"
      sed -i "s/>'/>/" "./html/build/$fileName.html"
    fi

    sed -i -r 's/^(<h1.*)\((.*)\)<(.*)/\1<span class="other-title">\(\2\)<\/span><\3/' "./html/build/$fileName.html"
    sed -i "/<hr \/>/{N;d;}" "./html/build/$fileName.html"
    sed -i "/''/d" "./html/build/$fileName.html"
    echo -e "<div class=\"keep-together\">\n$(cat "./html/build/$fileName.html")" > "./html/build/$fileName.html"
    echo -e "</div>\n<br/>" >> "./html/build/$fileName.html"
    echo -e "$(cat "./html/build/$fileName.html")" >> ./html/churchSongs.html
*/
