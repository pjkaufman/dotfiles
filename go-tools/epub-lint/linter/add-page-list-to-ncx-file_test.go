//go:build unit

package linter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

// TODO: handle play order existing (starting at more than 1) and the addition of the ncx values

const (
	emptyNcxFile             = "</ncx>"
	ncxFileWithEmptyPageList = `<body>
	  <pageList>
		</pageList>
		</body>`

//		bodyWithNavPageList = `<body>
//		<nav epub:type="page-list" hidden="">
//	  <ol>
//	    <li><a href="filename.htm#pg1">1</a></li>
//	    <li><a href="filename.htm#pg2">2</a></li>
//	    <li><a href="ch2.html#page3">3</a></li>
//	    <li><a href="ch3.html#p4">4</a></li>
//	  </ol>
//
// </nav>
// </body>`
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
	"make sure that a list of epub page ids does not modify the file if it already has a nav file list": {
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
	// "make sure that a list of epub page ids gets added to the file if it does not already have a nav file list": {
	// 	InputText: emptyBody,
	// 	InputPageIds: []linter.PageIdInfo{
	// 		{
	// 			Id:     "pg1",
	// 			Number: 1,
	// 			File:   "filename.htm",
	// 		},
	// 		{
	// 			Id:     "pg2",
	// 			Number: 2,
	// 			File:   "filename.htm",
	// 		},
	// 		{
	// 			Id:     "page3",
	// 			Number: 3,
	// 			File:   "ch2.html",
	// 		},
	// 		{
	// 			Id:     "p4",
	// 			Number: 4,
	// 			File:   "ch3.html",
	// 		},
	// 	},
	// 	ExpectedText: bodyWithNavPageList,
	// },
	// "make sure that a list of epub page ids properly gets added when the page numbers are out of order": {
	// 	InputText: emptyBody,
	// 	InputPageIds: []linter.PageIdInfo{
	// 		{
	// 			Id:     "p4",
	// 			Number: 4,
	// 			File:   "ch3.html",
	// 		},
	// 		{
	// 			Id:     "pg1",
	// 			Number: 1,
	// 			File:   "filename.htm",
	// 		},
	// 		{
	// 			Id:     "page3",
	// 			Number: 3,
	// 			File:   "ch2.html",
	// 		},
	// 		{
	// 			Id:     "pg2",
	// 			Number: 2,
	// 			File:   "filename.htm",
	// 		},
	// 	},
	// 	ExpectedText: bodyWithNavPageList,
	// },
}

func TestAddPageListToNcxFile(t *testing.T) {
	for name, args := range AddPageListToNcxFileTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.AddPageListToNcxFile(args.InputText, args.InputPageIds)

			assert.Equal(t, args.ExpectedText, actual)
		})
	}
}
