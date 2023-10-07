package converter

import (
	"fmt"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
)

func ConvertMdToHtmlCover(l logger.Logger, fileManager filehandler.FileManager, filePath string) string {
	contents := fileManager.ReadInFileContents(filePath)
	html := mdToHTML([]byte(contents))

	return fmt.Sprintf("<div style=\"text-align: center\">\n%s</div>\n", html)
}
