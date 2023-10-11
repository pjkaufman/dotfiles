package converter

import (
	"fmt"
	"strings"
)

func BuildHtmlCover(stylesHtml, coverMd string) string {
	var html = strings.Builder{}
	html.WriteString(stylesHtml)

	coverHtml := mdToHTML([]byte(coverMd))
	coverHtml = fmt.Sprintf("<div style=\"text-align: center\">\n%s</div>\n", coverHtml)
	coverHtml = strings.ReplaceAll(coverHtml, "\n\n", "\n")
	html.WriteString(coverHtml)

	return html.String()
}
