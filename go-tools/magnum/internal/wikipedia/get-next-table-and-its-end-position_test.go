//go:build unit

package wikipedia_test

import (
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/wikipedia"
	"github.com/stretchr/testify/assert"
)

type GetNextTableAndItsEndPositionTestCase struct {
	InputHtml         string
	ExpectedTableHtml string
	ExpectedStopIndex int
}

const (
	theWrongWayToUseHealingMagicLightNovelTable = `<table class="wikitable" width="100%" style="">


	<tbody><tr style="border-bottom: 3px solid rgb(204, 204, 255); --darkreader-inline-border-bottom: #000075;" data-darkreader-inline-border-bottom="">
	<th width="4%"><abbr title="Number">No.</abbr>
	</th>
	<th width="24%">Original release date
	</th>
	<th width="24%">Original ISBN
	</th>
	<th width="24%">English release date
	</th>
	<th width="24%">English ISBN
	</th></tr><tr style="text-align: center;"><th scope="row" id="vol1" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">1</th><td> March 25, 2016<sup id="cite_ref-6" class="reference"><a href="#cite_note-6">[6]</a></sup></td><td><style data-mw-deduplicate="TemplateStyles:r1215172403">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free.id-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-free a{background-size:contain}.mw-parser-output .id-lock-limited.id-lock-limited a,.mw-parser-output .id-lock-registration.id-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-limited a,body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-registration a{background-size:contain}.mw-parser-output .id-lock-subscription.id-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .id-lock-subscription a{background-size:contain}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}body:not(.skin-timeless):not(.skin-minerva) .mw-parser-output .cs1-ws-icon a{background-size:contain}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:#d33}.mw-parser-output .cs1-visible-error{color:#d33}.mw-parser-output .cs1-maint{display:none;color:#2C882D;margin-left:0.3em}.mw-parser-output .cs1-format{font-size:95%}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}html.skin-theme-clientpref-night .mw-parser-output .cs1-maint{color:#18911F}html.skin-theme-clientpref-night .mw-parser-output .cs1-visible-error,html.skin-theme-clientpref-night .mw-parser-output .cs1-hidden-error{color:#f8a397}@media(prefers-color-scheme:dark){html.skin-theme-clientpref-os .mw-parser-output .cs1-visible-error,html.skin-theme-clientpref-os .mw-parser-output .cs1-hidden-error{color:#f8a397}html.skin-theme-clientpref-os .mw-parser-output .cs1-maint{color:#18911F}}</style><style class="darkreader darkreader--sync" media="screen"></style><a href="/wiki/Special:BookSources/978-4-04-068185-6" title="Special:BookSources/978-4-04-068185-6">978-4-04-068185-6</a></td><td>August 23, 2022<sup id="cite_ref-7" class="reference"><a href="#cite_note-7">[7]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-1-64273-200-9" title="Special:BookSources/978-1-64273-200-9">978-1-64273-200-9</a></td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol2" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">2</th><td> June 24, 2016<sup id="cite_ref-8" class="reference"><a href="#cite_note-8">[8]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-068427-7" title="Special:BookSources/978-4-04-068427-7">978-4-04-068427-7</a></td><td>May 15, 2023<sup id="cite_ref-9" class="reference"><a href="#cite_note-9">[9]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-1-64273-232-0" title="Special:BookSources/978-1-64273-232-0">978-1-64273-232-0</a></td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol3" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">3</th><td> September 23, 2016<sup id="cite_ref-10" class="reference"><a href="#cite_note-10">[10]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-068636-3" title="Special:BookSources/978-4-04-068636-3">978-4-04-068636-3</a></td><td>August 22, 2023</td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-1-64273-286-3" title="Special:BookSources/978-1-64273-286-3">978-1-64273-286-3</a></td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol4" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">4</th><td> January 25, 2017<sup id="cite_ref-11" class="reference"><a href="#cite_note-11">[11]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-069054-4" title="Special:BookSources/978-4-04-069054-4">978-4-04-069054-4</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol5" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">5</th><td> April 25, 2017<sup id="cite_ref-12" class="reference"><a href="#cite_note-12">[12]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-069191-6" title="Special:BookSources/978-4-04-069191-6">978-4-04-069191-6</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol6" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">6</th><td> September 25, 2017<sup id="cite_ref-13" class="reference"><a href="#cite_note-13">[13]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-069498-6" title="Special:BookSources/978-4-04-069498-6">978-4-04-069498-6</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol7" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">7</th><td> February 24, 2018<sup id="cite_ref-14" class="reference"><a href="#cite_note-14">[14]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-069727-7" title="Special:BookSources/978-4-04-069727-7">978-4-04-069727-7</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol8" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">8</th><td> July 25, 2018<sup id="cite_ref-15" class="reference"><a href="#cite_note-15">[15]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-065023-4" title="Special:BookSources/978-4-04-065023-4">978-4-04-065023-4</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol9" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">9</th><td> November 24, 2018<sup id="cite_ref-16" class="reference"><a href="#cite_note-16">[16]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-065306-8" title="Special:BookSources/978-4-04-065306-8">978-4-04-065306-8</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol10" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">10</th><td> April 25, 2019<sup id="cite_ref-17" class="reference"><a href="#cite_note-17">[17]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-065682-3" title="Special:BookSources/978-4-04-065682-3">978-4-04-065682-3</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol11" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">11</th><td> October 25, 2019<sup id="cite_ref-18" class="reference"><a href="#cite_note-18">[18]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-064060-0" title="Special:BookSources/978-4-04-064060-0">978-4-04-064060-0</a></td><td>—</td><td>—</td></tr>
	<tr style="text-align: center;"><th scope="row" id="vol12" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">12</th><td> March 25, 2020<sup id="cite_ref-19" class="reference"><a href="#cite_note-19">[19]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-064538-4" title="Special:BookSources/978-4-04-064538-4">978-4-04-064538-4</a></td><td>—</td><td>—</td></tr>
	</tbody></table>`
	theWrongWayToUseHealingMagicLightNovelSection = `<h2><span class="mw-headline" id="Media">Media</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=The_Wrong_Way_to_Use_Healing_Magic&amp;action=edit&amp;section=3" title="Edit section: Media"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></h2>
	<h3><span class="mw-headline" id="Light_novels">Light novels</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=The_Wrong_Way_to_Use_Healing_Magic&amp;action=edit&amp;section=4" title="Edit section: Light novels"><span>edit</span></a><span class="mw-editsection-bracket">]</span></span></h3>
	<p>The series written by Kurokata began serialization online in March 2014 on the user-generated novel publishing website <a href="/wiki/Sh%C5%8Dsetsuka_ni_Nar%C5%8D" title="Shōsetsuka ni Narō">Shōsetsuka ni Narō</a>. It was later acquired by <a href="/wiki/Media_Factory" title="Media Factory">Media Factory</a>, who have published twelve volumes with illustrations by KeG between March 25, 2016 and March 25, 2020 under their MF Books imprint. The light novel is licensed in North America by One Peace Books.<sup id="cite_ref-One-Peace_5-0" class="reference"><a href="#cite_note-One-Peace-5">[5]</a></sup>
</p>
` + theWrongWayToUseHealingMagicLightNovelTable + `
<p>A sequel light novel series by the same author and illustrator, titled <i>The Wrong Way to Use Healing Magic Returns</i>, began publication on December 25, 2023.
</p>
<table class="wikitable" width="100%" style="">


<tbody><tr style="border-bottom: 3px solid rgb(204, 204, 255); --darkreader-inline-border-bottom: #000075;" data-darkreader-inline-border-bottom="">
<th width="4%"><abbr title="Number">No.</abbr>
</th>
<th width="48%">Japanese release date
</th>
<th width="48%">Japanese ISBN
</th></tr><tr style="text-align: center;"><th scope="row" id="vol1" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">1</th><td> December 25, 2023<sup id="cite_ref-20" class="reference"><a href="#cite_note-20">[20]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-683145-3" title="Special:BookSources/978-4-04-683145-3">978-4-04-683145-3</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol2" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">2</th><td> March 25, 2024<sup id="cite_ref-21" class="reference"><a href="#cite_note-21">[21]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1215172403"><a href="/wiki/Special:BookSources/978-4-04-683481-2" title="Special:BookSources/978-4-04-683481-2">978-4-04-683481-2</a></td></tr>
</tbody></table>`
)

// TODO: come back and figure out the discrepancy of 40 between actual and expected since the table is correctly being grabbed
var GetNextTableAndItsEndPositionTestCases = map[string]GetNextTableAndItsEndPositionTestCase{
	"a simple table row should get the correct amount of rows returned": {
		InputHtml:         theWrongWayToUseHealingMagicLightNovelSection,
		ExpectedTableHtml: theWrongWayToUseHealingMagicLightNovelTable,
		ExpectedStopIndex: 11724,
	},
}

func TestGetNextTableAndItsEndPosition(t *testing.T) {
	for name, args := range GetNextTableAndItsEndPositionTestCases {
		t.Run(name, func(t *testing.T) {
			actualTableHtml, actualStopPosition := wikipedia.GetNextTableAndItsEndPosition(args.InputHtml)

			assert.Equal(t, args.ExpectedTableHtml, actualTableHtml, "actual html was not the expected value")
			assert.Equal(t, args.ExpectedStopIndex, actualStopPosition, "actual stop position of the table was not the expected value")
		})
	}
}
