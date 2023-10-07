package linter

import (
	"fmt"
	"strings"
)

const hrBlankSpace = `hr.blankSpace {
border:0;
height:2em;
}`

func AddCssPageBreakIfMissing(fileContent string) string {
	if strings.Contains(fileContent, hrBlankSpace) {
		return fileContent
	}

	return fmt.Sprintf("%s\n%s", fileContent, hrBlankSpace)
}
