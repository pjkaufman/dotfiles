//go:build unit

package converter_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/stretchr/testify/assert"
)

type ConvertMdToHtmlSongTestCase struct {
	InputFilePath       string
	ExistingFiles       map[string]struct{}
	ExistingFolders     map[string]struct{}
	ExistingFileContent map[string]string
	PathToFolders       map[string][]string
	CmdErr              error
	ExpectedError       string
	ExpectPanic         bool
	ExpectedHtml        string
}

// errors that get handled as errors are represented as panics
var ConvertMdToHtmlSongTestCases = map[string]ConvertMdToHtmlSongTestCase{
	// "make sure that an empty file path causes a validation error": {
	// 	InputFilePath: "",
	// 	ExpectedError: cmd.FilePathArgEmpty,
	// 	ExpectPanic:   true,
	// },
	// "make sure that an non-markdown file path causes a validation error": {
	// 	InputFilePath: "file.txt",
	// 	ExpectedError: cmd.FilePathNotMarkdownFile,
	// 	ExpectPanic:   true,
	// },
	// "make sure that the file path not existing causes a validation error": {
	// 	InputFilePath: "file.md",
	// 	ExpectedError: `file-path: "file.md" must exist`,
	// 	ExpectPanic:   true,
	// },
	"a valid file should properly get turned into html": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `---
melody: 
key: Key C
authors: Chris Knauf 
in-church: N
verse: 
location: (MS16) (B20)
type: song
tags: 🎵
---

# He Is

\~ 1 \~ He is fairer than the lily of the valley, He is brighter than the morning star.  
He is purer than the snow, fresher than the breeze, Lovelier by far than all of these.

\~ 2 \~ But He calms all the storms, and conquers the raging seas.  
He is the high and lofty One who inhabits eternity,  
Creator of the universe, and He\'s clothed with majesty.  
He is and ever more shall be.  
(Repeat first stanza)
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div class="keep-together">
<h1 id="he-is">He Is</h1>
<div><div class="metadata"><div><div class="author">Chris Knauf</div></div><div><div class="key"><b>Key C</b></div></div><div><div class="location">(MS16) (B20)</div></div></div></div><br><br>
<p>~ 1 ~ He is fairer than the lily of the valley, He is brighter than the morning star.<br>
He is purer than the snow, fresher than the breeze, Lovelier by far than all of these.</p>
<p>~ 2 ~ But He calms all the storms, and conquers the raging seas.<br>
He is the high and lofty One who inhabits eternity,<br>
Creator of the universe, and He&rsquo;s clothed with majesty.<br>
He is and ever more shall be.<br>
(Repeat first stanza)</p>
</div>
<br>`,
	},
	"a valid file with another title should properly get converted into an html file": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `---
melody: 
key: Key G
authors: LaVerne & Edith Tripp
in-church: N
verse: 
location: (MS68) (B1)
copyright: unknown
type: song
tags: 🎵
---

# Above It All (There Stands Jesus)

Above it all, There stands Jesus. Above it all, He\'s still my King. \*He took my life,  
And He made me happy\*. Above it all, He\'s still the \*\*same\*\*.

\*This fleeting life is but a vapor\*  
\*\*King\*\*
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div class="keep-together">
<h1 id="above-it-all-there-stands-jesus">Above It All <span class="other-title">(There Stands Jesus)</span></h1>
<div><div class="metadata"><div><div class="author">LaVerne & Edith Tripp</div></div><div><div class="key"><b>Key G</b></div></div><div><div class="location">(MS68) (B1)</div></div></div></div><br><br>
<p>Above it all, There stands Jesus. Above it all, He&rsquo;s still my King. *He took my life,<br>
And He made me happy*. Above it all, He&rsquo;s still the **same**.</p>
<p>*This fleeting life is but a vapor*<br>
**King**</p>
</div>
<br>`,
	},
	"a valid file with just the melody in the second row should properly get converted into an html file": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `---
melody: (tune The Kingdom of God is Not Meat and Drink)
key: 
authors: I. Amundson
in-church: Y
verse: 
location: 
type: song
tags: 🎵
---

# Behold The Heavens

\~ 1 \~ Behold the heavens are open; Behold the face of the King  
We now have a new way of walking; We now have a new song to sing.

\~ 2 \~ Behold the heavens are open; Forever the veil has been rent;  
We\'re walking together in Zion; And there are no limits in Him.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div class="keep-together">
<h1 id="behold-the-heavens">Behold The Heavens</h1>
<div><div class="metadata row-padding"><div><div class="author"><b>I. Amundson</b></div></div><div><div class="key">&nbsp;&nbsp;&nbsp;&nbsp;</div></div><div><div class="location">&nbsp;&nbsp;&nbsp;&nbsp;</div></div></div><div class="metadata"><div><div class="melody-75"><b>(tune The Kingdom of God is Not Meat and Drink)</b></div></div></div></div><br><br>
<p>~ 1 ~ Behold the heavens are open; Behold the face of the King<br>
We now have a new way of walking; We now have a new song to sing.</p>
<p>~ 2 ~ Behold the heavens are open; Forever the veil has been rent;<br>
We&rsquo;re walking together in Zion; And there are no limits in Him.</p>
</div>
<br>`,
	},
	"a valid file with a verse in the second row should properly get converted into an html file": {
		InputFilePath: "file.md",
		ExistingFiles: map[string]struct{}{
			"file.md": {},
		},
		ExistingFileContent: map[string]string{
			"file.md": `---
melody: 
key: Key F
authors: 
in-church: 
verse: Ps. 57:9-11
location: (MS4) (B4)
type: song
tags: 🎵
---

# Be Thou Exalted

\~ 1 \~ Be Thou exalted, oh God (x3) above the heavens.  
Let Thy glory be above the whole earth.  
(Repeat)

\~ 2 \~ I will praise Thee oh Lord among the people.  
I will sing unto Thee among the nations.  
For Thy mercy is great unto the heavens  
And Thy truth unto the clouds.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedHtml: `<div class="keep-together">
<h1 id="be-thou-exalted">Be Thou Exalted</h1>
<div><div class="metadata row-padding"><div><div class="author">&nbsp;&nbsp;&nbsp;&nbsp;</div></div><div><div class="key"><b>Key F</b></div></div><div><div class="location">(MS4) (B4)</div></div></div><div class="metadata"><div><div class="melody">&nbsp;&nbsp;&nbsp;&nbsp;</div></div><div><div class="verse">Ps. 57:9-11</div></div></div></div><br><br>
<p>~ 1 ~ Be Thou exalted, oh God (x3) above the heavens.<br>
Let Thy glory be above the whole earth.<br>
(Repeat)</p>
<p>~ 2 ~ I will praise Thee oh Lord among the people.<br>
I will sing unto Thee among the nations.<br>
For Thy mercy is great unto the heavens<br>
And Thy truth unto the clouds.</p>
</div>
<br>`,
	},
}

func TestConvertMdToHtmlSong(t *testing.T) {
	for name, args := range ConvertMdToHtmlSongTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleConvertMdToHtmlSongPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, nil, nil, args.ExistingFileContent)
			actual := converter.ConvertMdToHtmlSong(log, fileHandler, args.InputFilePath)

			assert.Equal(t, args.ExpectedHtml, actual)
		})
	}
}

func handleConvertMdToHtmlSongPanic(t *testing.T, args ConvertMdToHtmlSongTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
