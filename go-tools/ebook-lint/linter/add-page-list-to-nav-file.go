package linter

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type PageIdInfo struct {
	Id     string
	Number int
	File   string
}

var validPageListAbbreviations = []string{"page", "pg", "p"}
var validPageListAbbrevsRegex = regexp.MustCompile(strings.Join(validPageListAbbreviations, "|"))
var navWithEpubPageList = regexp.MustCompile(`<nav[^\n>]*epub:type=["']page-list["'][^\n>]*>`)

func AddPageListToNavFile(text string, pageIds []PageIdInfo) string {
	if len(pageIds) == 0 || navWithEpubPageList.MatchString(text) {
		return text
	}

	var navPageList = createPageListForEpub3(pageIds)

	return strings.Replace(text, "</body>", navPageList+"\n</body>", 1)
}

func createPageListForEpub3(pageIds []PageIdInfo) string {
	if len(pageIds) == 0 {
		return ""
	}

	sort.Slice(pageIds, func(i, j int) bool {
		return pageIds[i].Number < pageIds[j].Number
	})

	var listEls strings.Builder
	for _, pageId := range pageIds {
		listEls.WriteString(fmt.Sprintf("    <li><a href=\"%s#%s\">%d</a></li>\n", pageId.File, pageId.Id, pageId.Number))
	}

	return fmt.Sprintf(`<nav epub:type="page-list" hidden="">
  <ol>
%s  </ol>
</nav>`, listEls.String())
}
