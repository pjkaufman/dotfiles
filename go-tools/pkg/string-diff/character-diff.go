package stringdiff

import (
	"bytes"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/fatih/color"
)

var (
	red   = color.New(color.BgRed, color.FgBlack).SprintFunc()
	green = color.New(color.BgGreen, color.FgBlack).SprintFunc()
)

// GetPrettyDiffString gets the diff string of the 2 passed in values where removals have a red background and additions have a green background
func GetPrettyDiffString(original, new string) string {
	diffString := diff.CharacterDiff(original, new)

	var buff bytes.Buffer
	var diffsLen = len(diffString)
	var char, nextChar, nextNextChar, section string
	var inSection bool
	for i := 0; i < len(diffString); {
		char = string(diffString[i])
		if char == "(" && i+2 < diffsLen && !inSection {
			nextChar = string(diffString[i+1])
			nextNextChar = string(diffString[i+2])
			if nextChar == "+" && nextNextChar == "+" {
				inSection = true

				i += 3
				continue
			} else if nextChar == "~" && nextNextChar == "~" {
				inSection = true

				i += 3
				continue
			}
		} else if char == "~" && i+2 < diffsLen && string(diffString[i+1]) == "~" && string(diffString[i+2]) == ")" {
			inSection = false
			buff.WriteString(red(section))
			section = ""

			i += 3
			continue
		} else if char == "+" && i+2 < diffsLen && string(diffString[i+1]) == "+" && string(diffString[i+2]) == ")" {
			inSection = false
			buff.WriteString(green(section))
			section = ""

			i += 3
			continue
		}

		if inSection {
			section += char
		} else {
			buff.WriteString(char)
		}

		i++
	}

	return convertUnicodeStringsToVisualRepresentations(buff.String())
}

func convertUnicodeStringsToVisualRepresentations(val string) string {
	val = strings.ReplaceAll(val, "â\u0080¦", "…")
	val = strings.ReplaceAll(val, "â\u0080\u0093", "–")
	val = strings.ReplaceAll(val, "â\u0097\u0087", "◇")

	return val
}
