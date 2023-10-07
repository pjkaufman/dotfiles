//go:build unit

package converter_test

import (
	"testing"

	filehandler "github.com/pjkaufman/dotfiles/go-tools/pkg/file-handler"
	"github.com/pjkaufman/dotfiles/go-tools/pkg/logger"
	"github.com/pjkaufman/dotfiles/go-tools/song-converter/internal/converter"
	"github.com/stretchr/testify/assert"
)

type ConvertMdToCsvTestCase struct {
	InputFilePath       string
	ExistingFiles       map[string]struct{}
	ExistingFileContent map[string]string
	ExpectedError       string
	ExpectPanic         bool
	ExpectedCsv         string
}

// errors that get handled as errors are represented as panics
var ConvertMdToCsvTestCases = map[string]ConvertMdToCsvTestCase{
	"make sure that the file path not existing causes an error": {
		InputFilePath: "file.md",
		ExpectedError: `could not read in file contents for "file.md": path not found`,
		ExpectPanic:   true,
	},
	"a valid file should properly get turned into a csv row": {
		InputFilePath: "He is Lord.md",
		ExistingFiles: map[string]struct{}{
			"He is Lord.md": {},
		},
		ExistingFileContent: map[string]string{
			"He is Lord.md": `---
melody: 
key: Key F
authors: 
in-church: 
verse: 
location: R53
type: song
tags: 🎵
---

# He Is Lord

He is Lord, He is Lord. He is risen from the dead and He is Lord.  
Ev\'ry knee shall bow ev\'ry tongue confess, That Jesus Christ is Lord.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "He is Lord|Red Book page 53||\n",
	},
	"make sure that no location is handled properly": {
		InputFilePath: "Bless This House.md",
		ExistingFiles: map[string]struct{}{
			"Bless This House.md": {},
		},
		ExistingFileContent: map[string]string{
			"Bless This House.md": `---
melody: 
key: 
authors: May H Brahe, Helen Taylor
in-church: N
verse: 
location: 
type: song
tags: 🎵
---

# Bless This House

\~ 1 \~ Bless this house, O Lord, we pray. Make it safe by night and day.  
Bless these walls so firm and stout, Keeping want and troubles out.

\~ 2 \~ Bless the roof and chimney top. Let thy love flow all about.  
Bless this house that it may prove Ever open to joy and truth.

\~ 3 \~ Bless us all that we may be Fit, O Lord, to dwell with thee.  
Bless us so that one day we, May dwell, dear Lord, with thee.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "Bless This House||May H Brahe, Helen Taylor|\n",
	},
	"make sure that blue book locations are handled properly": {
		InputFilePath: "Bigger Than All My Problems.md",
		ExistingFiles: map[string]struct{}{
			"Bigger Than All My Problems.md": {},
		},
		ExistingFileContent: map[string]string{
			"Bigger Than All My Problems.md": `---
melody: 
key: Key C
authors: Bill & Gloria Gaither 
in-church: N
verse: 
location: (B6)
type: song
tags: 🎵
---

# Bigger Than All My Problems

\~ 1 \~ Bigger than all the shadows that fall across my path  
God is bigger than any mountain that I can or cannot see;  
Bigger than my confusion, bigger than anything,  
God is bigger than any mountain that I can or cannot see.

CHORUS:  
Bigger than all my problems, bigger than all my fears;  
God is bigger than any mountain that I can or cannot see;  
Bigger than all my questions, bigger than anything,  
God is bigger than any mountain that I can or cannot see.

\~ 2 \~ Bigger than all the giants of fear and unbelief,  
God is bigger than any mountain that I can or cannot see;  
Bigger than all my hang-ups, bigger than anything,  
God is bigger than any mountain that I can or cannot see.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "Bigger Than All My Problems|(Blue Book page 6)|Bill & Gloria Gaither|\n",
	},
	"make sure that more songs we love locations are handled properly": {
		InputFilePath: "Flow Thou River.md",
		ExistingFiles: map[string]struct{}{
			"Flow Thou River.md": {},
		},
		ExistingFileContent: map[string]string{
			"Flow Thou River.md": `---
melody: 
key: Key C or E Flat
authors: M. Heartwell 
in-church: Y
verse: 
location: R32 (MS13) (B15)
type: song
tags: 🎵
---

# Flow Thou River

\~ 1 \~ Flow thou River, flow thou River, Forth from the Throne of God.  
Flow thou River, flow thou River, Forth from the Throne of God.

\~ 2 \~ \* That life may spring forth,  
That life may spring forth in a dry and thirsty land. That life may spring forth,  
That life may spring forth in a dry and thirsty land.

\~ 3 \~ \* Healing waters, healing waters, Flow from the Throne of God.  
Healing waters, healing waters, Flow from the Throne of God.

\~ 4 \~ \*\* Strength for today, strength for tomorrow, Flows from the Throne of God.  
Strength for today, strength for tomorrow, Flows from the Throne of God.

\*Differs from published songs by M. Heartwell  
\*\*Added by G.B.S.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "Flow Thou River|Red Book page 32 (More Songs We Love page 13) (Blue Book page 15)|M. Heartwell|Church\n",
	},
	"make sure that copyright for authors in the church are set to 'Church'": {
		InputFilePath: "Fill My Soul With Thy Spirit.md",
		ExistingFiles: map[string]struct{}{
			"Fill My Soul With Thy Spirit.md": {},
		},
		ExistingFileContent: map[string]string{
			"Fill My Soul With Thy Spirit.md": `---
melody: 
key: Key E Flat or F or C
authors: A. Ellis
in-church: Y
verse: 
location: R199 (B15)
type: song
tags: 🎵
---

# Fill My Soul With Thy Spirit

Fill my soul with Thy Spirit, Fill my heart with Thy love;  
Let my soul be rekindled with fire from above.  
Let me drink from that fountain; Flowing boundless and free,  
Fill my soul with Thy Spirit, With love fill thou me.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "Fill My Soul With Thy Spirit|Red Book page 199 (Blue Book page 15)|A. Ellis|Church\n",
	},
	"make sure that copyright YAML property gets used when not in church and it is present2": {
		InputFilePath: "A Glorious Church.md",
		ExistingFiles: map[string]struct{}{
			"A Glorious Church.md": {},
		},
		ExistingFileContent: map[string]string{
			"A Glorious Church.md": `---
melody: 
key: 
authors: Ralph E. Hudson
in-church: N
verse: 
location: (MS68)
copyright: Public Domain
type: song
tags: 🎵
---

# A Glorious Church

\~ 1 \~ Do you hear them coming, Brother, Thronging up the steeps of light,  
Clad in glorious shining garments Blood-washed garments pure and white?

CHORUS:  
\'Tis a glorious Church without spot or wrinkle, Washed in the blood of the Lamb.  
\'Tis a glorious Church, without spot or wrinkle, Washed in the blood of the Lamb.

\~ 2 \~ Do you hear the stirring anthems Filling all the earth and sky?  
\'Tis a grand victorious army. Lift its banner up on high!

\~ 3 \~ Never fear the clouds of sorrow; Never fear the storms of sin.  
Even now our joys begin.

\~ 4 \~ Wave the banner, shout His praises, For our victory is nigh!  
We shall join our conquering Savior. We shall reign with Him on high.
`,
		},
		ExpectedError: "",
		ExpectPanic:   false,
		ExpectedCsv:   "A Glorious Church|(More Songs We Love page 68)|Ralph E. Hudson|Public Domain\n",
	},
}

func TestConvertMdToCsv(t *testing.T) {
	for name, args := range ConvertMdToCsvTestCases {
		t.Run(name, func(t *testing.T) {
			defer handleConvertMdToCsvPanic(t, args)

			var log = logger.NewMockLoggerHandler()
			var fileHandler = filehandler.NewMockFileHandler(log, args.ExistingFiles, nil, nil, args.ExistingFileContent)
			actual := converter.ConvertMdToCsv(log, fileHandler, args.InputFilePath, args.InputFilePath)
			assert.Equal(t, args.ExpectedCsv, actual)
		})
	}
}

func handleConvertMdToCsvPanic(t *testing.T, args ConvertMdToCsvTestCase) {
	if r := recover(); r != nil {
		assert.True(t, args.ExpectPanic, "an error was not expected")
		assert.Equal(t, args.ExpectedError, r, "the error message did not match the expected error message")
	}
}
