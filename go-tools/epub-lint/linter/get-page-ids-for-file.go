package linter

import (
	"fmt"
	"regexp"
	"strings"
)

var validPageListIdsRegex = regexp.MustCompile(fmt.Sprintf(`id[ \t]*=["']((%s)\d+)["']`, strings.Join(validPageListAbbreviations, "|")))

func GetPageIdsForFile(text, file string, pageIds []PageIdInfo) []PageIdInfo {
	var validPageIds = validPageListIdsRegex.FindAllStringIndex(text, -1)
	if len(validPageIds) == 0 {
		return pageIds
	}

	for _, locs := range validPageIds {
		pageIds = append(pageIds, parseIdToIdInfo(text[locs[0]:locs[1]], file))
	}

	return pageIds
}
