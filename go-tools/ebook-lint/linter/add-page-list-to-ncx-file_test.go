//go:build unit

package linter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/ebook-lint/linter"
	"github.com/stretchr/testify/assert"
)

const (
	emptyNcxFile             = "</ncx>"
	ncxFileWithEmptyPageList = `<body>
	  <pageList>
		</pageList>
		</body>
		</ncx>`
	ncxFileWithPageListAndNavMap = `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN"
	"http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
	
	<ncx version="2005-1" xml:lang="en" xmlns="http://www.daisy.org/z3986/2005/ncx/">
	
		<head>
	<!-- The following four metadata items are required for all NCX documents,
	including those that conform to the relaxed constraints of OPS 2.0 -->
	
			<meta name="dtb:uid" content="123456789X"/> <!-- same as in .opf -->
			<meta name="dtb:depth" content="1"/> <!-- 1 or higher -->
			<meta name="dtb:totalPageCount" content="0"/> <!-- must be 0 -->
			<meta name="dtb:maxPageNumber" content="0"/> <!-- must be 0 -->
		</head>
	
		<docTitle>
			<text>Pride and Prejudice</text>
		</docTitle>
	
		<docAuthor>
			<text>Austen, Jane</text>
		</docAuthor>
	
		<navMap>
			<navPoint class="chapter" id="chapter1" playOrder="1">
				<navLabel><text>Chapter 1</text></navLabel>
				<content src="chapter1.xhtml"/>
			</navPoint>
		</navMap>
	
	</ncx>`
	ncxFileWithPageListAndNavMapAndPageList = `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN"
	"http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
	
	<ncx version="2005-1" xml:lang="en" xmlns="http://www.daisy.org/z3986/2005/ncx/">
	
		<head>
	<!-- The following four metadata items are required for all NCX documents,
	including those that conform to the relaxed constraints of OPS 2.0 -->
	
			<meta name="dtb:uid" content="123456789X"/> <!-- same as in .opf -->
			<meta name="dtb:depth" content="1"/> <!-- 1 or higher -->
			<meta name="dtb:totalPageCount" content="0"/> <!-- must be 0 -->
			<meta name="dtb:maxPageNumber" content="0"/> <!-- must be 0 -->
		</head>
	
		<docTitle>
			<text>Pride and Prejudice</text>
		</docTitle>
	
		<docAuthor>
			<text>Austen, Jane</text>
		</docAuthor>
	
		<navMap>
			<navPoint class="chapter" id="chapter1" playOrder="1">
				<navLabel><text>Chapter 1</text></navLabel>
				<content src="chapter1.xhtml"/>
			</navPoint>
		</navMap>
	
	<pageList>
  <pageTarget id="page1" type="normal" value="1" playOrder="2">
    <navLabel><text>1</text></navLabel>
    <content src="chapter1.xhtml#pg1"/>
  </pageTarget>
  <pageTarget id="page2" type="normal" value="2" playOrder="3">
    <navLabel><text>2</text></navLabel>
    <content src="chapter1.xhtml#pg2"/>
  </pageTarget>
  <pageTarget id="page3" type="normal" value="3" playOrder="4">
    <navLabel><text>3</text></navLabel>
    <content src="chapter2.html#page3"/>
  </pageTarget>
  <pageTarget id="page4" type="normal" value="4" playOrder="5">
    <navLabel><text>4</text></navLabel>
    <content src="chapter3.html#p4"/>
  </pageTarget>
</pageList>
</ncx>`
	ncxFileWithPageListAndNoNavMap = `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN"
	"http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
	
	<ncx version="2005-1" xml:lang="en" xmlns="http://www.daisy.org/z3986/2005/ncx/">
	
		<head>
	<!-- The following four metadata items are required for all NCX documents,
	including those that conform to the relaxed constraints of OPS 2.0 -->
	
			<meta name="dtb:uid" content="123456789X"/> <!-- same as in .opf -->
			<meta name="dtb:depth" content="1"/> <!-- 1 or higher -->
			<meta name="dtb:totalPageCount" content="0"/> <!-- must be 0 -->
			<meta name="dtb:maxPageNumber" content="0"/> <!-- must be 0 -->
		</head>
	
		<docTitle>
			<text>Pride and Prejudice</text>
		</docTitle>
	
		<docAuthor>
			<text>Austen, Jane</text>
		</docAuthor>

	</ncx>`
	ncxFileWithPageListAndNoNavMapAndPageList = `<?xml version="1.0" encoding="UTF-8"?>
	<!DOCTYPE ncx PUBLIC "-//NISO//DTD ncx 2005-1//EN"
	"http://www.daisy.org/z3986/2005/ncx-2005-1.dtd">
	
	<ncx version="2005-1" xml:lang="en" xmlns="http://www.daisy.org/z3986/2005/ncx/">
	
		<head>
	<!-- The following four metadata items are required for all NCX documents,
	including those that conform to the relaxed constraints of OPS 2.0 -->
	
			<meta name="dtb:uid" content="123456789X"/> <!-- same as in .opf -->
			<meta name="dtb:depth" content="1"/> <!-- 1 or higher -->
			<meta name="dtb:totalPageCount" content="0"/> <!-- must be 0 -->
			<meta name="dtb:maxPageNumber" content="0"/> <!-- must be 0 -->
		</head>
	
		<docTitle>
			<text>Pride and Prejudice</text>
		</docTitle>
	
		<docAuthor>
			<text>Austen, Jane</text>
		</docAuthor>

	<pageList>
  <pageTarget id="page1" type="normal" value="1" playOrder="1">
    <navLabel><text>1</text></navLabel>
    <content src="chapter1.xhtml#pg1"/>
  </pageTarget>
  <pageTarget id="page2" type="normal" value="2" playOrder="2">
    <navLabel><text>2</text></navLabel>
    <content src="chapter1.xhtml#pg2"/>
  </pageTarget>
  <pageTarget id="page3" type="normal" value="3" playOrder="3">
    <navLabel><text>3</text></navLabel>
    <content src="chapter2.html#page3"/>
  </pageTarget>
  <pageTarget id="page4" type="normal" value="4" playOrder="4">
    <navLabel><text>4</text></navLabel>
    <content src="chapter3.html#p4"/>
  </pageTarget>
</pageList>
</ncx>`
)

type AddPageListToNcxFileTestCase struct {
	InputText    string
	InputPageIds []linter.PageIdInfo
	ExpectedText string
}

var AddPageListToNcxFileTestCases = map[string]AddPageListToNcxFileTestCase{
	"make sure that a nil list of epub page ids does not modify the file": {
		InputText:    emptyNcxFile,
		InputPageIds: nil,
		ExpectedText: emptyNcxFile,
	},
	"make sure that an empty list of epub page ids does not modify the file": {
		InputText:    emptyNcxFile,
		InputPageIds: []linter.PageIdInfo{},
		ExpectedText: emptyNcxFile,
	},
	"make sure that a list of epub page ids does not modify the file if it already has a page list": {
		InputText: ncxFileWithEmptyPageList,
		InputPageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				Number: 1,
				File:   "filename.htm",
			},
			{
				Id:     "pg2",
				Number: 2,
				File:   "filename.htm",
			},
		},
		ExpectedText: ncxFileWithEmptyPageList,
	},
	"make sure that a list of epub page ids gets added to the file if it does not already have a page list": {
		InputText: ncxFileWithPageListAndNavMap,
		InputPageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				Number: 1,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "pg2",
				Number: 2,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "page3",
				Number: 3,
				File:   "chapter2.html",
			},
			{
				Id:     "p4",
				Number: 4,
				File:   "chapter3.html",
			},
		},
		ExpectedText: ncxFileWithPageListAndNavMapAndPageList,
	},
	"make sure that a list of epub page ids properly gets added when the page numbers are out of order": {
		InputText: ncxFileWithPageListAndNavMap,
		InputPageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				Number: 1,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "pg2",
				Number: 2,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "page3",
				Number: 3,
				File:   "chapter2.html",
			},
			{
				Id:     "p4",
				Number: 4,
				File:   "chapter3.html",
			},
		},
		ExpectedText: ncxFileWithPageListAndNavMapAndPageList,
	},
	"make sure that a list of epub page ids properly gets added when there is no nav map": {
		InputText: ncxFileWithPageListAndNoNavMap,
		InputPageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				Number: 1,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "pg2",
				Number: 2,
				File:   "chapter1.xhtml",
			},
			{
				Id:     "page3",
				Number: 3,
				File:   "chapter2.html",
			},
			{
				Id:     "p4",
				Number: 4,
				File:   "chapter3.html",
			},
		},
		ExpectedText: ncxFileWithPageListAndNoNavMapAndPageList,
	},
}

func TestAddPageListToNcxFile(t *testing.T) {
	for name, args := range AddPageListToNcxFileTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.AddPageListToNcxFile(args.InputText, args.InputPageIds)

			assert.Equal(t, args.ExpectedText, actual)
		})
	}
}
