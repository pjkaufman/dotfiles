package linter

import "regexp"

var letterEndingParagraphRegex = regexp.MustCompile(`(\n[^\n]*([a-zA-z]|,"?))( ?)(</p>\n)(<p[^>]*>)([^\n]*</p>)`)

/*
*
case 1:
<p class="calibre1"><a id="p169"></a>If there are objects with a simple structure and the same properties, then they can be recognized as a single "set" allowing decomposition of the </p>
<p class="calibre1">"set" rather than each object separately. </p>
case 2:

<p class="calibre1">The information provided by Edward Clark had a brief project plan. </p>
<p class="calibre1">Tatsuya ran his eyes through the original document, and Miyuki read the translated text. </p>
<p class="calibre1">Minami placed a cup of freshly brewed tea in front of them. As if on a signal, Tatsuya and Miyuki simultaneously looked up from the electronic paper with the details of the project. </p>
<p class="calibre1">"…The so-called Dione, it does not seem to be Saturn's companion </p>
<p class="calibre1">'Dione', but the goddess of Greek myths." </p>
<p class="calibre1">"Right. Wife of Zeus, who gave birth to Aphrodite. From that version of the myth, where Aphrodite is born from sea foam." </p>
*/
// TODO: handle more than 2 paragraphs back to back needing to be merged
func GetPotentiallyBrokenLines(fileContent string) map[string]string {
	var subMatches = letterEndingParagraphRegex.FindAllStringSubmatch(fileContent, -1)
	var originalToSuggested = make(map[string]string, len(subMatches))
	if len(subMatches) == 0 {
		return originalToSuggested
	}

	for _, groups := range subMatches {
		var space = groups[3]
		if space == "" {
			space = " "
		}

		originalToSuggested[groups[0]] = groups[1] + space + groups[6]
	}

	return originalToSuggested
}
