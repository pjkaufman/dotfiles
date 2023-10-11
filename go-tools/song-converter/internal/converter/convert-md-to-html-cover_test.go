//go:build unit

package converter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/stretchr/testify/assert"
)

type ConvertMdToHtmlCoverTestCase struct {
	InputStylesContent string
	InputCoverMd       string
	ExpectedHtml       string
}

// TODO: handle styles addition
var ConvertMdToHtmlCoverTestCases = map[string]ConvertMdToHtmlCoverTestCase{
	"a valid file should properly get turned into html with no style content": {
		InputStylesContent: "",
		InputCoverMd: `# Church Songs - E Version

<br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/>

### Key:

#### Books

R=Red Book (Songs We Love)  
MS=More Songs section of Songs We Love  
B=Blue Book

#### Authors

CA= Cyndi Aarrestad  
EHW= Ewald Wanagas  
FTP= Frank Paterson  
GBS= Gail Shepherd  
ZW= Zelma Wanagas

<br/> <br/>

*\*When searching electronic format, punctuation & spelling will cause
non returns*

*\*Punctuation alters the alphabetical order*`,
		ExpectedHtml: `<div style="text-align: center">
<h1 id="church-songs-e-version">Church Songs - E Version</h1>
<p><br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/> <br/></p>
<h3 id="key">Key:</h3>
<h4 id="books">Books</h4>
<p>R=Red Book (Songs We Love)<br>
MS=More Songs section of Songs We Love<br>
B=Blue Book</p>
<h4 id="authors">Authors</h4>
<p>CA= Cyndi Aarrestad<br>
EHW= Ewald Wanagas<br>
FTP= Frank Paterson<br>
GBS= Gail Shepherd<br>
ZW= Zelma Wanagas</p>
<p><br/> <br/></p>
<p><em>*When searching electronic format, punctuation &amp; spelling will cause
non returns</em></p>
<p><em>*Punctuation alters the alphabetical order</em></p>
</div>
`,
	},
}

func TestConvertMdToHtmlCover(t *testing.T) {
	for name, args := range ConvertMdToHtmlCoverTestCases {
		t.Run(name, func(t *testing.T) {
			actual := converter.BuildHtmlCover(args.InputStylesContent, args.InputCoverMd)

			assert.Equal(t, args.ExpectedHtml, actual)
		})
	}
}
