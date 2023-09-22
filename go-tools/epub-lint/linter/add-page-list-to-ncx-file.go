package linter

import (
	"fmt"
	"sort"
	"strings"
)

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
