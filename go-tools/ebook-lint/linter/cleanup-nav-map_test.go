//go:build unit

package linter_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/ebook-lint/linter"
	"github.com/stretchr/testify/assert"
)

const (
	sampleNavMapFile1 = `<?xml version="1.0" encoding="UTF-8"?><ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
<head>
<meta content="9780316398572" name="dtb:uid"/>
<meta content="1" name="dtb:depth"/>
<meta content="0" name="dtb:totalPageCount"/>
<meta content="0" name="dtb:maxPageNumber"/>
</head>
<docTitle><text>The Asterisk War, Vol. 1: Encounter with a Fiery Princess</text></docTitle>
<docAuthor><text>Yuu Miyazaki and okiura</text></docAuthor>
<navMap>
<navPoint id="cover" playOrder="1"><navLabel><text>Cover</text></navLabel><content src="cover.xhtml"/></navPoint>
<navPoint id="welcome" playOrder="2"><navLabel><text>Welcome</text></navLabel><content src="welcome.xhtml"/></navPoint>
<navPoint playOrder="3" id="photo-insert"><navLabel><text>Insert</text></navLabel><content src="photo-insert.xhtml"/></navPoint>
<navPoint id="titlepage" playOrder="4"><navLabel><text>Title Page</text></navLabel><content src="titlepage.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8376" playOrder="5"><navLabel><text>Map</text></navLabel><content src="preface001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8378" playOrder="6"><navLabel><text>Chapter 1: Glühen Rose</text></navLabel><content src="chapter001.xhtml#Ref_8204"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8379" playOrder="7"><navLabel><text>Chapter 2: Asterisk, The City of Academic Warfare</text></navLabel><content src="chapter004.xhtml#Ref_8207"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8380" playOrder="8"><navLabel><text>Chapter 3: Her Noble Eyes</text></navLabel><content src="chapter009.xhtml#Ref_8210"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8381" playOrder="9"><navLabel><text>Chapter 4: Reminiscence and Reunion</text></navLabel><content src="chapter012.xhtml#Ref_8213"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8382" playOrder="10"><navLabel><text>Chapter 5: The Ser Veresta</text></navLabel><content src="chapter016.xhtml#Ref_8216"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8383" playOrder="11"><navLabel><text>Chapter 6: A Holiday for Two</text></navLabel><content src="chapter019.xhtml#Ref_8219"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8384" playOrder="12"><navLabel><text>Chapter 7: Unchained</text></navLabel><content src="chapter022.xhtml#Ref_8222"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8385" playOrder="13"><navLabel><text>Epilogue</text></navLabel><content src="chapter027.xhtml#Ref_8228"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8386" playOrder="14"><navLabel><text>Afterword</text></navLabel><content src="chapter028.xhtml#Ref_8225"/></navPoint>
<navPoint id="id_1_25" playOrder="15"><navLabel><text>Manga Preview</text></navLabel><content src="chapter047.xhtml"/></navPoint>
<navPoint id="newsletterSignup1" playOrder="16"><navLabel><text>Yen Newsletter</text></navLabel><content src="newsletterSignup.xhtml"/></navPoint>
<navPoint id="toc" playOrder="17"><navLabel><text>Table of Contents</text></navLabel><content src="toc.xhtml"/></navPoint>
<navPoint id="copyright" playOrder="18"><navLabel><text>Copyright</text></navLabel><content src="copyright.xhtml"/></navPoint>
</navMap>
</ncx>`
	expectedSampleNavMapFile1 = `<?xml version="1.0" encoding="UTF-8"?><ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
<head>
<meta content="9780316398572" name="dtb:uid"/>
<meta content="1" name="dtb:depth"/>
<meta content="0" name="dtb:totalPageCount"/>
<meta content="0" name="dtb:maxPageNumber"/>
</head>
<docTitle><text>The Asterisk War, Vol. 1: Encounter with a Fiery Princess</text></docTitle>
<docAuthor><text>Yuu Miyazaki and okiura</text></docAuthor>
<navMap>
<navPoint id="cover" playOrder="1"><navLabel><text>Cover</text></navLabel><content src="cover.xhtml"/></navPoint>
<navPoint id="welcome" playOrder="2"><navLabel><text>Welcome</text></navLabel><content src="welcome.xhtml"/></navPoint>
<navPoint playOrder="3" id="photo-insert"><navLabel><text>Insert</text></navLabel><content src="photo-insert.xhtml"/></navPoint>
<navPoint id="titlepage" playOrder="4"><navLabel><text>Title Page</text></navLabel><content src="titlepage.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8376" playOrder="5"><navLabel><text>Map</text></navLabel><content src="preface001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8378" playOrder="6"><navLabel><text>Chapter 1: Glühen Rose</text></navLabel><content src="chapter001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8379" playOrder="7"><navLabel><text>Chapter 2: Asterisk, The City of Academic Warfare</text></navLabel><content src="chapter004.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8380" playOrder="8"><navLabel><text>Chapter 3: Her Noble Eyes</text></navLabel><content src="chapter009.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8381" playOrder="9"><navLabel><text>Chapter 4: Reminiscence and Reunion</text></navLabel><content src="chapter012.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8382" playOrder="10"><navLabel><text>Chapter 5: The Ser Veresta</text></navLabel><content src="chapter016.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8383" playOrder="11"><navLabel><text>Chapter 6: A Holiday for Two</text></navLabel><content src="chapter019.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8384" playOrder="12"><navLabel><text>Chapter 7: Unchained</text></navLabel><content src="chapter022.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8385" playOrder="13"><navLabel><text>Epilogue</text></navLabel><content src="chapter027.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8386" playOrder="14"><navLabel><text>Afterword</text></navLabel><content src="chapter028.xhtml"/></navPoint>
<navPoint id="id_1_25" playOrder="15"><navLabel><text>Manga Preview</text></navLabel><content src="chapter047.xhtml"/></navPoint>
<navPoint id="newsletterSignup1" playOrder="16"><navLabel><text>Yen Newsletter</text></navLabel><content src="newsletterSignup.xhtml"/></navPoint>
<navPoint id="toc" playOrder="17"><navLabel><text>Table of Contents</text></navLabel><content src="toc.xhtml"/></navPoint>
<navPoint id="copyright" playOrder="18"><navLabel><text>Copyright</text></navLabel><content src="copyright.xhtml"/></navPoint>
</navMap>
</ncx>`
	sampleNavMapFile2 = `
<navMap>
  <navPoint id="navPoint-1" playOrder="1">
    <navLabel>
      <text>Prologue</text>
    </navLabel>
    <content src="Text/Body.xhtml"/>
  </navPoint>
  <navPoint id="navPoint-2" playOrder="2">
    <navLabel>
      <text>Chapter 1  MaslenitsaThe Sun Festival</text>
    </navLabel>
    <content src="Text/Body.xhtml#Chapter_1_E28093_MaslenitsaThe_Sun_Festival"/>
  </navPoint>
  <navPoint id="navPoint-3" playOrder="3">
    <navLabel>
      <text>Chapter 2  Homecoming</text>
    </navLabel>
    <content src="Text/Body.xhtml#Chapter_2_E28093_Homecoming"/>
  </navPoint>
  <navPoint id="navPoint-4" playOrder="4">
    <navLabel>
      <text>Chapter 3  Invaders</text>
    </navLabel>
    <content src="Text/Body.xhtml#Chapter_3_E28093_Invaders"/>
  </navPoint>
  <navPoint id="navPoint-5" playOrder="5">
    <navLabel>
      <text>Chapter 4  The ShervidIllusory Princess of the Hollow Shadow</text>
    </navLabel>
    <content src="Text/Body.xhtml#Chapter_4_E28093_The_ShervidIllusory_Princess_of_the_Hollow_Shadow"/>
  </navPoint>
  <navPoint id="navPoint-6" playOrder="6">
    <navLabel>
      <text>Translator's Notes and References</text>
    </navLabel>
    <content src="Text/Body.xhtml#Translator27s_Notes_and_References"/>
  </navPoint>
</navMap>`
	sampleNavMapFile3 = `<?xml version="1.0" encoding="UTF-8"?><ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
<head>
<meta content="9780316398572" name="dtb:uid"/>
<meta content="1" name="dtb:depth"/>
<meta content="0" name="dtb:totalPageCount"/>
<meta content="0" name="dtb:maxPageNumber"/>
</head>
<docTitle><text>The Asterisk War, Vol. 1: Encounter with a Fiery Princess</text></docTitle>
<docAuthor><text>Yuu Miyazaki and okiura</text></docAuthor>
<navMap>
<navPoint id="cover" playOrder="1"><navLabel><text>Cover</text></navLabel><content src="cover.xhtml"/></navPoint>
<navPoint id="welcome" playOrder="#"><navLabel><text>Welcome</text></navLabel><content src="welcome.xhtml"/></navPoint>
<navPoint playOrder="3" id="photo-insert"><navLabel><text>Insert</text></navLabel><content src="photo-insert.xhtml"/></navPoint>
<navPoint id="titlepage" playOrder="fdaf"><navLabel><text>Title Page</text></navLabel><content src="titlepage.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8376" playOrder="5"><navLabel><text>Map</text></navLabel><content src="preface001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8378" playOrder="23"><navLabel><text>Chapter 1: Glühen Rose</text></navLabel><content src="chapter001.xhtml#Ref_8204"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8379" playOrder=""><navLabel><text>Chapter 2: Asterisk, The City of Academic Warfare</text></navLabel><content src="chapter004.xhtml#Ref_8207"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8380" playOrder="12"><navLabel><text>Chapter 3: Her Noble Eyes</text></navLabel><content src="chapter009.xhtml#Ref_8210"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8381" playOrder="1"><navLabel><text>Chapter 4: Reminiscence and Reunion</text></navLabel><content src="chapter012.xhtml#Ref_8213"/></navPoint>
</navMap>
</ncx>`
	expectedSampleNavMapFile3 = `<?xml version="1.0" encoding="UTF-8"?><ncx xmlns="http://www.daisy.org/z3986/2005/ncx/" version="2005-1">
<head>
<meta content="9780316398572" name="dtb:uid"/>
<meta content="1" name="dtb:depth"/>
<meta content="0" name="dtb:totalPageCount"/>
<meta content="0" name="dtb:maxPageNumber"/>
</head>
<docTitle><text>The Asterisk War, Vol. 1: Encounter with a Fiery Princess</text></docTitle>
<docAuthor><text>Yuu Miyazaki and okiura</text></docAuthor>
<navMap>
<navPoint id="cover" playOrder="1"><navLabel><text>Cover</text></navLabel><content src="cover.xhtml"/></navPoint>
<navPoint id="welcome" playOrder="2"><navLabel><text>Welcome</text></navLabel><content src="welcome.xhtml"/></navPoint>
<navPoint playOrder="3" id="photo-insert"><navLabel><text>Insert</text></navLabel><content src="photo-insert.xhtml"/></navPoint>
<navPoint id="titlepage" playOrder="4"><navLabel><text>Title Page</text></navLabel><content src="titlepage.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8376" playOrder="5"><navLabel><text>Map</text></navLabel><content src="preface001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8378" playOrder="6"><navLabel><text>Chapter 1: Glühen Rose</text></navLabel><content src="chapter001.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8379" playOrder="7"><navLabel><text>Chapter 2: Asterisk, The City of Academic Warfare</text></navLabel><content src="chapter004.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8380" playOrder="8"><navLabel><text>Chapter 3: Her Noble Eyes</text></navLabel><content src="chapter009.xhtml"/></navPoint>
<navPoint id="File_AsteriskWar_TOC_8381" playOrder="9"><navLabel><text>Chapter 4: Reminiscence and Reunion</text></navLabel><content src="chapter012.xhtml"/></navPoint>
</navMap>
</ncx>`
)

type CleanupNavMapTestCase struct {
	InputText    string
	ExpectedText string
	expectedErr  error
}

var CleanupNavMapTestCases = map[string]CleanupNavMapTestCase{
	"make sure that content has the ids removed": {
		InputText:    sampleNavMapFile1,
		ExpectedText: expectedSampleNavMapFile1,
		expectedErr:  nil,
	},
	"make sure an error is thrown when there are multiple references to the same file in the nav map": {
		InputText:    sampleNavMapFile2,
		ExpectedText: sampleNavMapFile2,
		expectedErr:  fmt.Errorf(`there is more than one reference to "%s". Consider restructuring the epub to fix up the nav map`, "Text/Body.xhtml"),
	},
	"make sure no nav map results in an error being thrown": {
		InputText:    "",
		ExpectedText: "",
		expectedErr:  linter.ErrNoNavMap,
	},
	"make sure no nav map ending element results in an error being thrown": {
		InputText:    "<navMap>",
		ExpectedText: "<navMap>",
		expectedErr:  linter.ErrNoEndOfNavMap,
	},
	"make sure that a nav map with a nav point without a src for the content element throws an error": {
		InputText: `<navMap>
  <navPoint id="navPoint-6" playOrder="6">
    <navLabel>
      <text>Translator's Notes and References</text>
    </navLabel>
    <content/>
  </navPoint>
</navMap>`,
		ExpectedText: `<navMap>
  <navPoint id="navPoint-6" playOrder="6">
    <navLabel>
      <text>Translator's Notes and References</text>
    </navLabel>
    <content/>
  </navPoint>
</navMap>`,
		expectedErr: fmt.Errorf("possible problem with content tag src attribute: []"),
	},
	"make sure that the play order is properly set when it is either not a number or out of order": {
		InputText:    sampleNavMapFile3,
		ExpectedText: expectedSampleNavMapFile3,
		expectedErr:  nil,
	},
}

func TestCleanupNavMap(t *testing.T) {
	for name, args := range CleanupNavMapTestCases {
		t.Run(name, func(t *testing.T) {
			actual, err := linter.CleanupNavMap(args.InputText)

			assert.Equal(t, args.ExpectedText, actual)
			assert.Equal(t, args.expectedErr, err)
		})
	}
}
