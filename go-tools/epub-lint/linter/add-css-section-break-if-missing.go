package linter

import (
	"fmt"
	"strings"
)

const (
	hrCharacter = `hr.character {
overflow: visible;
border:0;
text-align:center;
}`
	hrContentAfterTemplate = `hr.character:after {
content: "%s";
display:inline-block;
position:relative;
font-size:1em;
padding:1em;
}`
)

func AddCssSectionBreakIfMissing(fileContent, contextBreak string) string {
	if strings.Contains(fileContent, hrCharacter) {
		return fileContent
	}

	return fmt.Sprintf("%s\n%s\n%s", fileContent, hrCharacter, fmt.Sprintf(hrContentAfterTemplate, contextBreak))
}
