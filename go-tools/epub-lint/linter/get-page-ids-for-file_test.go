//go:build unit

package linter_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

const (
	fileName1 = "filename.html"
	fileName2 = "ch2.xhtml"
)

type GetPageIdsForFileTestCase struct {
	InputText           string
	InputFileName       string
	InputPageIds        []linter.PageIdInfo
	ExpectedFilePageIds []linter.PageIdInfo
}

var GetPageIdsForFileTestCases = map[string]GetPageIdsForFileTestCase{
	"make sure that a file with no page ids returns an empty array": {
		InputText:           "<p>No id here</p>",
		InputFileName:       fileName1,
		InputPageIds:        []linter.PageIdInfo{},
		ExpectedFilePageIds: []linter.PageIdInfo{},
	},
	"make sure that a file with a page id using the word 'page' can properly be pulled back": {
		InputText: `<p id="page1234"></p>
		<p id="page1235"></p>
		<p id="page1236"></p>
		<p id="page1237"></p>
		<p id="page1238"></p>
		<p id="page1239"></p>`,
		InputFileName: fileName1,
		InputPageIds:  []linter.PageIdInfo{},
		ExpectedFilePageIds: []linter.PageIdInfo{
			{
				Id:     "page1234",
				File:   fileName1,
				Number: 1234,
			},
			{
				Id:     "page1235",
				File:   fileName1,
				Number: 1235,
			},
			{
				Id:     "page1236",
				File:   fileName1,
				Number: 1236,
			},
			{
				Id:     "page1237",
				File:   fileName1,
				Number: 1237,
			},
			{
				Id:     "page1238",
				File:   fileName1,
				Number: 1238,
			},
			{
				Id:     "page1239",
				File:   fileName1,
				Number: 1239,
			},
		},
	},
	"make sure that a file with a page id using the word 'pg' can properly be pulled back": {
		InputText: `<p id="pg1"></p>
		<p id="pg2"></p>
		<p id="pg3"></p>
		<p id="pg4"></p>`,
		InputFileName: fileName1,
		InputPageIds:  []linter.PageIdInfo{},
		ExpectedFilePageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				File:   fileName1,
				Number: 1,
			},
			{
				Id:     "pg2",
				File:   fileName1,
				Number: 2,
			},
			{
				Id:     "pg3",
				File:   fileName1,
				Number: 3,
			},
			{
				Id:     "pg4",
				File:   fileName1,
				Number: 4,
			},
		},
	},
	"make sure that a file with a page id using the word 'p' can properly be pulled back": {
		InputText: `<p id="p1"></p>
		<p id="p2"></p>
		<p id="p3"></p>
		<p id="p4"></p>`,
		InputFileName: fileName1,
		InputPageIds:  []linter.PageIdInfo{},
		ExpectedFilePageIds: []linter.PageIdInfo{
			{
				Id:     "p1",
				File:   fileName1,
				Number: 1,
			},
			{
				Id:     "p2",
				File:   fileName1,
				Number: 2,
			},
			{
				Id:     "p3",
				File:   fileName1,
				Number: 3,
			},
			{
				Id:     "p4",
				File:   fileName1,
				Number: 4,
			},
		},
	},
	"make sure that a file with a page id using mixed prefixes can properly be pulled back": {
		InputText: `<p id="pg1"></p>
		<p id="page2"></p>
		<p id="p3"></p>
		<p id="pg4"></p>`,
		InputFileName: fileName1,
		InputPageIds:  []linter.PageIdInfo{},
		ExpectedFilePageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				File:   fileName1,
				Number: 1,
			},
			{
				Id:     "page2",
				File:   fileName1,
				Number: 2,
			},
			{
				Id:     "p3",
				File:   fileName1,
				Number: 3,
			},
			{
				Id:     "pg4",
				File:   fileName1,
				Number: 4,
			},
		},
	},
	"make sure that a file with page ids returns an array with values added onto it from the current text": {
		InputText: `<p id="page1234"></p>
		<p id="page1235"></p>
		<p id="page1236"></p>
		<p id="page1237"></p>
		<p id="page1238"></p>
		<p id="page1239"></p>`,
		InputFileName: fileName2,
		InputPageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				File:   fileName1,
				Number: 1,
			},
			{
				Id:     "page2",
				File:   fileName1,
				Number: 2,
			},
		},
		ExpectedFilePageIds: []linter.PageIdInfo{
			{
				Id:     "pg1",
				File:   fileName1,
				Number: 1,
			},
			{
				Id:     "page2",
				File:   fileName1,
				Number: 2,
			},
			{
				Id:     "page1234",
				File:   fileName2,
				Number: 1234,
			},
			{
				Id:     "page1235",
				File:   fileName2,
				Number: 1235,
			},
			{
				Id:     "page1236",
				File:   fileName2,
				Number: 1236,
			},
			{
				Id:     "page1237",
				File:   fileName2,
				Number: 1237,
			},
			{
				Id:     "page1238",
				File:   fileName2,
				Number: 1238,
			},
			{
				Id:     "page1239",
				File:   fileName2,
				Number: 1239,
			},
		},
	},
}

func TestGetPageIdsForFile(t *testing.T) {
	for name, args := range GetPageIdsForFileTestCases {
		t.Run(name, func(t *testing.T) {
			actual := linter.GetPageIdsForFile(args.InputText, args.InputFileName, args.InputPageIds)

			assert.Equal(t, args.ExpectedFilePageIds, actual, "page ids should be equal")
		})
	}
}
