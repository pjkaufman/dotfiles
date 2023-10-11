package converter

import "strings"

type MdFileInfo struct {
	FilePath     string
	FileName     string
	FileContents string
}

func BuildHtmlBody(stylesHtml string, mdInfo []MdFileInfo) (string, error) {
	var html = strings.Builder{}
	html.WriteString(stylesHtml + "\n")

	for _, mdData := range mdInfo {
		fileContentInHtml, err := ConvertMdToHtmlSong(mdData.FilePath, mdData.FileContents)
		if err != nil {
			return "", err
		}

		html.WriteString(fileContentInHtml + "\n")
	}

	return html.String(), nil
}
