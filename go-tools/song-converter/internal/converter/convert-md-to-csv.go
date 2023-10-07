package converter

import (
	"fmt"
	"strings"

	"github.com/adrg/frontmatter"
	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func ConvertMdToCsv(l logger.Logger, fileManager filehandler.FileManager, fileName, filePath string, csvContents strings.Builder) {
	contents := fileManager.ReadInFileContents(filePath)
	var metadata SongMetadata
	_, err := frontmatter.Parse(strings.NewReader(contents), &metadata)
	if err != nil {
		l.WriteError(fmt.Sprintf(`There was an error getting the frontmatter for file '%s': %s`, filePath, err))
	}

	csvContents.WriteString(strings.Replace(fileName, ".md", "", 1) + "|" + buildMetadataCsv(&metadata) + "\n")
}

func buildMetadataCsv(metadata *SongMetadata) string {
	if metadata == nil {
		return "||"
	}

	var copyright = metadata.Copyright
	if strings.EqualFold(metadata.InChurch, "Y") {
		copyright = "Church"
	}

	return fmt.Sprintf("%s|%s|%s", updateBookLocationInfo(metadata.BookLocation), metadata.Authors, copyright)
}

func updateBookLocationInfo(bookLocation string) string {
	if bookLocation == "" {
		return ""
	}

	var newBookLocation = strings.ReplaceAll(bookLocation, "B", "Blue Book page ")
	newBookLocation = strings.ReplaceAll(newBookLocation, "R", "Red Book page ")
	newBookLocation = strings.ReplaceAll(newBookLocation, "MS", "More Songs We Love page ")

	return newBookLocation
}
