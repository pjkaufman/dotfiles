package linter

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/pjkaufman/dotfiles/go-tools/utils"
)

type PageIdInfo struct {
	Id     string
	Number int
	File   string
}

var validPageListAbbreviations = []string{"page", "pg", "p"}
var validPageListIdsRegex = regexp.MustCompile(fmt.Sprintf(`id[ \t]*=["']((%s)\d+)["']`, strings.Join(validPageListAbbreviations, "|")))
var validPageListAbbrevsRegex = regexp.MustCompile(strings.Join(validPageListAbbreviations, "|"))
var navWithEpubPageList = regexp.MustCompile(`<nav[^\n>]*epub:type=["']page-list["'][^\n>]*>`)

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

func parseIdToIdInfo(id, file string) PageIdInfo {
	var startOfId = strings.Index(id, "=")
	var val = id[startOfId+2 : len(id)-1]
	var num = validPageListAbbrevsRegex.ReplaceAllString(val, "")
	intVar, err := strconv.Atoi(num)
	if err != nil {
		utils.WriteWarn(fmt.Sprintf(`Possible bad id "%s" tried to parse "%s": %v`, id, num, err))
	}

	return PageIdInfo{
		Id:     val,
		Number: intVar,
		File:   file,
	}
}

func AddPageListToNcxFile(text string, pageIds []PageIdInfo) string {
	if len(pageIds) == 0 || strings.Contains(text, "<pageList") {
		return text
	}

	var pageList = createPageListForNcx(pageIds, strings.Count(text, "playOrder=")+1)

	return strings.Replace(text, "</ncx>", pageList+"\n</ncx>", 1)
}

func createPageListForNcx(pageIds []PageIdInfo, startingPlayOrder int) string {
	if len(pageIds) == 0 {
		return ""
	}

	var playOrder = startingPlayOrder
	sort.Slice(pageIds, func(i, j int) bool {
		return pageIds[i].Number < pageIds[j].Number
	})

	var pageListEls strings.Builder
	for _, pageId := range pageIds {
		pageListEls.WriteString(fmt.Sprintf(`  <pageTarget id="page%[1]d" type="normal" value="%[1]d" playOrder="%d">`+"\n", pageId.Number, playOrder))
		pageListEls.WriteString(fmt.Sprintf("    <navLabel><text>%d</text></navLabel>\n", pageId.Number))
		pageListEls.WriteString(fmt.Sprintf(`    <content src="%s#%s"/>`+"\n", pageId.File, pageId.Id))
		pageListEls.WriteString("</pageTarget>\n")
		playOrder++
	}

	return fmt.Sprintf(`<pageList>
%s</pageList>`, pageListEls.String())
}

func AddPageListToNavFile(text string, pageIds []PageIdInfo) string {
	if len(pageIds) == 0 || navWithEpubPageList.MatchString(text) {
		return text
	}

	var navPageList = CreatePageListForEpub3(pageIds)

	return strings.Replace(text, "</body>", navPageList+"\n</body>", 1)
}

func CreatePageListForEpub3(pageIds []PageIdInfo) string {
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
