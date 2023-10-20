package strings

import "fmt"

const (
	cliLineSeparator    = "-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"
	fileSummaryTemplate = `
%[1]s
Before:
%s %s
After:
%s %s
%[1]s
`
	filesSummaryTemplate = `
%[1]s
Before:
%s
After:
%s
%[1]s
`
)

func FileSizeSummary(originalFile, newFile string, oldKbSize, newKbSize float64) string {
	return fmt.Sprintf(fileSummaryTemplate, cliLineSeparator, originalFile, KbSizeToString(oldKbSize), newFile, KbSizeToString(newKbSize))
}

func FilesSizeSummary(oldKbSizeSum, newKbSizeSum float64) string {
	return fmt.Sprintf(filesSummaryTemplate, cliLineSeparator, KbSizeToString(oldKbSizeSum), KbSizeToString(newKbSizeSum))
}
