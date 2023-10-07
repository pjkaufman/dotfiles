//go:build unit

package linter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

type GetPotentiallyBrokenLinesTestCase struct {
	InputText           string
	ExpectedSuggestions map[string]string
}

var GetPotentiallyBrokenLinesTestCases = map[string]GetPotentiallyBrokenLinesTestCase{
	"make sure that a file with no potentially broken paragraphs gives no suggestions": {
		InputText: `<p>Here is some content.</p>
<p>Here is some more content</p>`,
		ExpectedSuggestions: map[string]string{},
	},
	"make sure that a file with paragraphs that end in a letter get picked up as potentially needing a change": {
		InputText: `<p>Here is some content.</p>
		<p class="calibre1"><a id="p169"></a>If there are objects with a simple structure and the same properties, then they can be recognized as a single "set" allowing decomposition of the </p>
		<p class="calibre1">"set" rather than each object separately. </p>`,
		ExpectedSuggestions: map[string]string{
			`
		<p class="calibre1"><a id="p169"></a>If there are objects with a simple structure and the same properties, then they can be recognized as a single "set" allowing decomposition of the </p>
		<p class="calibre1">"set" rather than each object separately. </p>`: `
		<p class="calibre1"><a id="p169"></a>If there are objects with a simple structure and the same properties, then they can be recognized as a single "set" allowing decomposition of the "set" rather than each object separately. </p>`,
		},
	},
	"make sure that a file with paragraphs that end in a comma get picked up as potentially needing a change": {
		InputText: `<p class="calibre1">The information provided by Edward Clark had a brief project plan. </p>
		<p class="calibre1">Tatsuya ran his eyes through the original document, and Miyuki read the translated text. </p>
		<p class="calibre1">Minami placed a cup of freshly brewed tea in front of them. As if on a signal, Tatsuya and Miyuki simultaneously looked up from the electronic paper with the details of the project. </p>
		<p class="calibre1">"…The so-called Dione, it does not seem to be Saturn's companion </p>
		<p class="calibre1">'Dione', but the goddess of Greek myths." </p>
		<p class="calibre1">"Right. Wife of Zeus, who gave birth to Aphrodite. From that version of the myth, where Aphrodite is born from sea foam." </p>`,
		ExpectedSuggestions: map[string]string{
			`
		<p class="calibre1">"…The so-called Dione, it does not seem to be Saturn's companion </p>
		<p class="calibre1">'Dione', but the goddess of Greek myths." </p>`: `
		<p class="calibre1">"…The so-called Dione, it does not seem to be Saturn's companion 'Dione', but the goddess of Greek myths." </p>`,
		},
	},
	"make sure that a file with paragraphs that end in a number get picked up as potentially needing a change": {
		InputText: `<p class="calibre1">The Deputy Director showed interest in these words and encouraged him to continue. The Director of the Intelligence Department was absent at this meeting. The Deputy Director who was present was also a person not disclosed to the public. </p>
		<p class="calibre1">"I think you all already know that the Yotsuba family, to which Shiba Tatsuya belongs, is in a cooperative relationship with the 1-0-1 </p>
		<p class="calibre1">Battalion." </p>
		<p class="calibre1">After Onda's words, those sitting at the table nodded. </p>`,
		ExpectedSuggestions: map[string]string{
			`
		<p class="calibre1">"I think you all already know that the Yotsuba family, to which Shiba Tatsuya belongs, is in a cooperative relationship with the 1-0-1 </p>
		<p class="calibre1">Battalion." </p>`: `
		<p class="calibre1">"I think you all already know that the Yotsuba family, to which Shiba Tatsuya belongs, is in a cooperative relationship with the 1-0-1 Battalion." </p>`,
		},
	},
	"make sure that a file with multiple paragraphs back to back thatpotentially are broken, get condensed down into 1 suggestions": {
		InputText: `<p>Some content here.</p>
		<p>Here is a list, </p>
		<p>a set of todos,</p>
		<p>and its own sentence. </p>`,
		ExpectedSuggestions: map[string]string{
			`
		<p>Here is a list, </p>
		<p>a set of todos,</p>
		<p>and its own sentence. </p>`: `
		<p>Here is a list, a set of todos, and its own sentence. </p>`,
		},
	},
}

func TestGetPotentiallyBrokenLines(t *testing.T) {
	for name, args := range GetPotentiallyBrokenLinesTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.GetPotentiallyBrokenLines(args.InputText)

			assert.Equal(t, args.ExpectedSuggestions, actual)
		})
	}
}
