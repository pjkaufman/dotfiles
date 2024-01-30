//go:build unit

package wikipedia_test

import (
	"testing"
	"time"

	"github.com/pjkaufman/dotfiles/go-tools/magnum/internal/wikipedia"
	"github.com/stretchr/testify/assert"
)

type ParseWikipediaTableToVolumeInfoTestCase struct {
	InputTableHtml     string
	InputNamePrefix    string
	ExpectedVolumeInfo []wikipedia.VolumeInfo
}

const (
	mushokuTensieTable = `<table class="wikitable" style="text-align: center;">

<tbody><tr>
<th rowspan="2" scope="col" width="3%">Volume no.
</th>
<th rowspan="2" scope="col" width="20%">Content
</th>
<th rowspan="2" scope="col" width="10%">Japanese release date
</th>
<th rowspan="2" scope="col" width="10%">Japanese ISBN
</th>
<th colspan="2" scope="col" width="15%">English release date
</th>
<th rowspan="2" scope="col" width="10%">English ISBN
</th></tr>
<tr>
<th scope="col">Digital
</th>
<th scope="col">Physical
</th></tr>
<tr>
<th scope="row">1
</th>
<td style="text-align: left;">Web Novel 1
</td>
<td><span data-sort-value="000000002014-01-24-0000" style="white-space:nowrap">January 24, 2014</span><sup id="cite_ref-LN_1_8-0" class="reference"><a href="#cite_note-LN_1-8">[6]</a></sup>
</td>
<td><style data-mw-deduplicate="TemplateStyles:r1133582631">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free a,.mw-parser-output .citation .cs1-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-limited a,.mw-parser-output .id-lock-registration a,.mw-parser-output .citation .cs1-lock-limited a,.mw-parser-output .citation .cs1-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-subscription a,.mw-parser-output .citation .cs1-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:#d33}.mw-parser-output .cs1-visible-error{color:#d33}.mw-parser-output .cs1-maint{display:none;color:#3a3;margin-left:0.3em}.mw-parser-output .cs1-format{font-size:95%}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}</style><a href="/wiki/Special:BookSources/978-4-04-066220-6" title="Special:BookSources/978-4-04-066220-6">978-4-04-066220-6</a>
</td>
<td><span data-sort-value="000000002019-04-04-0000" style="white-space:nowrap">April 4, 2019</span>
</td>
<td><span data-sort-value="000000002019-05-21-0000" style="white-space:nowrap">May 21, 2019</span><sup id="cite_ref-9" class="reference"><a href="#cite_note-9">[7]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64275-138-3" title="Special:BookSources/978-1-64275-138-3">978-1-64275-138-3</a>
</td></tr>
<tr>
<th scope="row">2
</th>
<td style="text-align: left;">Web Novel 2
</td>
<td><span data-sort-value="000000002014-03-25-0000" style="white-space:nowrap">March 25, 2014</span><sup id="cite_ref-LN_2_10-0" class="reference"><a href="#cite_note-LN_2-10">[8]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-066393-7" title="Special:BookSources/978-4-04-066393-7">978-4-04-066393-7</a>
</td>
<td><span data-sort-value="000000002019-05-23-0000" style="white-space:nowrap">May 23, 2019</span>
</td>
<td><span data-sort-value="000000002019-07-30-0000" style="white-space:nowrap">July 30, 2019</span><sup id="cite_ref-11" class="reference"><a href="#cite_note-11">[9]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64275-140-6" title="Special:BookSources/978-1-64275-140-6">978-1-64275-140-6</a>
</td></tr>
<tr>
<tr>
<th scope="row">26
</th>
<td style="text-align: left;">End of Web Novel 23 and 24
</td>
<td><span data-sort-value="000000002022-11-25-0000" style="white-space:nowrap">November 25, 2022</span><sup id="cite_ref-58" class="reference"><a href="#cite_note-58">[56]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-681933-8" title="Special:BookSources/978-4-04-681933-8">978-4-04-681933-8</a>
</td>
<td>
</td>
<td><span data-sort-value="000000002024-03-12-0000" style="white-space:nowrap">March 12, 2024</span><sup id="cite_ref-59" class="reference"><a href="#cite_note-59">[57]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/979-8-88-843435-2" title="Special:BookSources/979-8-88-843435-2">979-8-88-843435-2</a>
</td></tr>
</tbody></table>`
	skeletonKnightTable = `<table class="wikitable" style="text-align: center; width:100%">

<tbody><tr>
<th rowspan="2" scope="col">No.
</th>
<th rowspan="2" scope="col">Original release date
</th>
<th rowspan="2" scope="col">Original ISBN
</th>
<th colspan="2" scope="col">English release date
</th>
<th rowspan="2" scope="col">English ISBN
</th></tr>
<tr>
<th scope="col">Digital
</th>
<th scope="col">Physical
</th></tr>
<tr>
<th scope="row">1
</th>
<td><span data-sort-value="000000002015-06-25-0000" style="white-space:nowrap">June 25, 2015</span><sup id="cite_ref-LN_1_6-0" class="reference"><a href="#cite_note-LN_1-6">[6]</a></sup>
</td>
<td><style data-mw-deduplicate="TemplateStyles:r1133582631">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free a,.mw-parser-output .citation .cs1-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-limited a,.mw-parser-output .id-lock-registration a,.mw-parser-output .citation .cs1-lock-limited a,.mw-parser-output .citation .cs1-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-subscription a,.mw-parser-output .citation .cs1-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:#d33}.mw-parser-output .cs1-visible-error{color:#d33}.mw-parser-output .cs1-maint{display:none;color:#3a3;margin-left:0.3em}.mw-parser-output .cs1-format{font-size:95%}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}</style><a href="/wiki/Special:BookSources/978-4-86-554054-3" title="Special:BookSources/978-4-86-554054-3">978-4-86-554054-3</a>
</td>
<td><span data-sort-value="000000002019-03-28-0000" style="white-space:nowrap">March 28, 2019</span>
</td>
<td><span data-sort-value="000000002019-06-11-0000" style="white-space:nowrap">June 11, 2019</span><sup id="cite_ref-7" class="reference"><a href="#cite_note-7">[7]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-275064-5" title="Special:BookSources/978-1-64-275064-5">978-1-64-275064-5</a>
</td></tr>
<tr>
<th scope="row">2
</th>
<td><span data-sort-value="000000002015-10-25-0000" style="white-space:nowrap">October 25, 2015</span><sup id="cite_ref-LN_2_8-0" class="reference"><a href="#cite_note-LN_2-8">[8]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554075-8" title="Special:BookSources/978-4-86-554075-8">978-4-86-554075-8</a>
</td>
<td><span data-sort-value="000000002019-06-20-0000" style="white-space:nowrap">June 20, 2019</span>
</td>
<td><span data-sort-value="000000002019-09-24-0000" style="white-space:nowrap">September 24, 2019</span><sup id="cite_ref-9" class="reference"><a href="#cite_note-9">[9]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-275129-1" title="Special:BookSources/978-1-64-275129-1">978-1-64-275129-1</a>
</td></tr>
<tr>
<th scope="row">3
</th>
<td><span data-sort-value="000000002016-03-25-0000" style="white-space:nowrap">March 25, 2016</span><sup id="cite_ref-LN_3_10-0" class="reference"><a href="#cite_note-LN_3-10">[10]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554109-0" title="Special:BookSources/978-4-86-554109-0">978-4-86-554109-0</a>
</td>
<td><span data-sort-value="000000002019-08-08-0000" style="white-space:nowrap">August 8, 2019</span>
</td>
<td><span data-sort-value="000000002019-10-15-0000" style="white-space:nowrap">October 15, 2019</span><sup id="cite_ref-11" class="reference"><a href="#cite_note-11">[11]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-275706-4" title="Special:BookSources/978-1-64-275706-4">978-1-64-275706-4</a>
</td></tr>
<tr>
<th scope="row">4
</th>
<td><span data-sort-value="000000002016-07-25-0000" style="white-space:nowrap">July 25, 2016</span><sup id="cite_ref-LN_4_12-0" class="reference"><a href="#cite_note-LN_4-12">[12]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554146-5" title="Special:BookSources/978-4-86-554146-5">978-4-86-554146-5</a>
</td>
<td><span data-sort-value="000000002019-11-14-0000" style="white-space:nowrap">November 14, 2019</span>
</td>
<td><span data-sort-value="000000002020-02-25-0000" style="white-space:nowrap">February 25, 2020</span><sup id="cite_ref-13" class="reference"><a href="#cite_note-13">[13]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-505195-4" title="Special:BookSources/978-1-64-505195-4">978-1-64-505195-4</a>
</td></tr>
<tr>
<th scope="row">5
</th>
<td><span data-sort-value="000000002016-11-25-0000" style="white-space:nowrap">November 25, 2016</span><sup id="cite_ref-LN_5_14-0" class="reference"><a href="#cite_note-LN_5-14">[14]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554169-4" title="Special:BookSources/978-4-86-554169-4">978-4-86-554169-4</a>
</td>
<td><span data-sort-value="000000002020-02-27-0000" style="white-space:nowrap">February 27, 2020</span>
</td>
<td><span data-sort-value="000000002020-08-18-0000" style="white-space:nowrap">August 18, 2020</span><sup id="cite_ref-15" class="reference"><a href="#cite_note-15">[15]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-505464-1" title="Special:BookSources/978-1-64-505464-1">978-1-64-505464-1</a>
</td></tr>
<tr>
<th scope="row">6
</th>
<td><span data-sort-value="000000002017-04-25-0000" style="white-space:nowrap">April 25, 2017</span><sup id="cite_ref-LN_6_16-0" class="reference"><a href="#cite_note-LN_6-16">[16]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554209-7" title="Special:BookSources/978-4-86-554209-7">978-4-86-554209-7</a>
</td>
<td><span data-sort-value="000000002020-06-04-0000" style="white-space:nowrap">June 4, 2020</span>
</td>
<td><span data-sort-value="000000002020-09-08-0000" style="white-space:nowrap">September 8, 2020</span><sup id="cite_ref-17" class="reference"><a href="#cite_note-17">[17]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-505725-3" title="Special:BookSources/978-1-64-505725-3">978-1-64-505725-3</a>
</td></tr>
<tr>
<th scope="row">7
</th>
<td><span data-sort-value="000000002017-08-25-0000" style="white-space:nowrap">August 25, 2017</span><sup id="cite_ref-LN_7_18-0" class="reference"><a href="#cite_note-LN_7-18">[18]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554249-3" title="Special:BookSources/978-4-86-554249-3">978-4-86-554249-3</a>
</td>
<td><span data-sort-value="000000002020-08-20-0000" style="white-space:nowrap">August 20, 2020</span>
</td>
<td><span data-sort-value="000000002020-12-01-0000" style="white-space:nowrap">December 1, 2020</span><sup id="cite_ref-19" class="reference"><a href="#cite_note-19">[19]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-505795-6" title="Special:BookSources/978-1-64-505795-6">978-1-64-505795-6</a>
</td></tr>
<tr>
<th scope="row">8
</th>
<td><span data-sort-value="000000002018-03-25-0000" style="white-space:nowrap">March 25, 2018</span><sup id="cite_ref-LN_8_20-0" class="reference"><a href="#cite_note-LN_8-20">[20]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554318-6" title="Special:BookSources/978-4-86-554318-6">978-4-86-554318-6</a>
</td>
<td><span data-sort-value="000000002020-11-05-0000" style="white-space:nowrap">November 5, 2020</span>
</td>
<td><span data-sort-value="000000002021-02-23-0000" style="white-space:nowrap">February 23, 2021</span><sup id="cite_ref-21" class="reference"><a href="#cite_note-21">[21]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-505977-6" title="Special:BookSources/978-1-64-505977-6">978-1-64-505977-6</a>
</td></tr>
<tr>
<th scope="row">9
</th>
<td><span data-sort-value="000000002019-03-25-0000" style="white-space:nowrap">March 25, 2019</span><sup id="cite_ref-LN_9_22-0" class="reference"><a href="#cite_note-LN_9-22">[22]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554432-9" title="Special:BookSources/978-4-86-554432-9">978-4-86-554432-9</a>
</td>
<td><span data-sort-value="000000002021-05-27-0000" style="white-space:nowrap">May 27, 2021</span>
</td>
<td><span data-sort-value="000000002021-08-03-0000" style="white-space:nowrap">August 3, 2021</span><sup id="cite_ref-23" class="reference"><a href="#cite_note-23">[23]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-827204-2" title="Special:BookSources/978-1-64-827204-2">978-1-64-827204-2</a>
</td></tr>
<tr>
<th scope="row">10
</th>
<td><span data-sort-value="000000002022-03-25-0000" style="white-space:nowrap">March 25, 2022</span><sup id="cite_ref-LN_10_24-0" class="reference"><a href="#cite_note-LN_10-24">[24]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-86-554549-4" title="Special:BookSources/978-4-86-554549-4">978-4-86-554549-4</a>
</td>
<td><span data-sort-value="000000002023-02-09-0000" style="white-space:nowrap">February 9, 2023</span>
</td>
<td><span data-sort-value="000000002023-04-11-0000" style="white-space:nowrap">April 11, 2023</span><sup id="cite_ref-25" class="reference"><a href="#cite_note-25">[25]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-827264-6" title="Special:BookSources/978-1-64-827264-6">978-1-64-827264-6</a>
</td></tr>
<tr>
<th scope="row">11
</th>
<td>TBA
</td>
<td>TBA
</td>
<td>TBA
</td>
<td><span data-sort-value="000000002025-04-22-0000" style="white-space:nowrap">April 22, 2025</span><sup id="cite_ref-26" class="reference"><a href="#cite_note-26">[26]</a></sup>
</td>
<td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64-827337-7" title="Special:BookSources/978-1-64-827337-7">978-1-64-827337-7</a>
</td></tr>
</tbody></table>`
	asteriskWarTable = `<table class="wikitable" width="100%" style="">


<tbody><tr style="border-bottom: 3px solid #CCF">
<th width="4%"><abbr title="Number">No.</abbr>
</th>
<th width="24%">Original release date
</th>
<th width="24%">Original ISBN
</th>
<th width="24%">English release date
</th>
<th width="24%">English ISBN
</th></tr>


<tr style="text-align: center;"><th scope="row" id="vol1" style="text-align: center; font-weight: normal; background-color: transparent;">1</th><td> September 25, 2012<sup id="cite_ref-First_14-1" class="reference"><a href="#cite_note-First-14">[7]</a></sup></td><td><style data-mw-deduplicate="TemplateStyles:r1133582631">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free a,.mw-parser-output .citation .cs1-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-limited a,.mw-parser-output .id-lock-registration a,.mw-parser-output .citation .cs1-lock-limited a,.mw-parser-output .citation .cs1-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-subscription a,.mw-parser-output .citation .cs1-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:#d33}.mw-parser-output .cs1-visible-error{color:#d33}.mw-parser-output .cs1-maint{display:none;color:#3a3;margin-left:0.3em}.mw-parser-output .cs1-format{font-size:95%}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}</style><a href="/wiki/Special:BookSources/978-4-04-066697-6" title="Special:BookSources/978-4-04-066697-6">978-4-04-066697-6</a></td><td>August 30, 2016<sup id="cite_ref-Eng_ver1_16-1" class="reference"><a href="#cite_note-Eng_ver1-16">[9]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-0-31-631527-2" title="Special:BookSources/978-0-31-631527-2">978-0-31-631527-2</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol2" style="text-align: center; font-weight: normal; background-color: transparent;">2</th><td> January 25, 2013<sup id="cite_ref-18" class="reference"><a href="#cite_note-18">[11]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-066698-3" title="Special:BookSources/978-4-04-066698-3">978-4-04-066698-3</a></td><td>December 20, 2016<sup id="cite_ref-19" class="reference"><a href="#cite_note-19">[12]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-0-31-639858-9" title="Special:BookSources/978-0-31-639858-9">978-0-31-639858-9</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol3" style="text-align: center; font-weight: normal; background-color: transparent;">3</th><td> May 24, 2013<sup id="cite_ref-20" class="reference"><a href="#cite_note-20">[13]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-067093-5" title="Special:BookSources/978-4-04-067093-5">978-4-04-067093-5</a></td><td>April 18, 2017<sup id="cite_ref-21" class="reference"><a href="#cite_note-21">[14]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-0-31-639860-2" title="Special:BookSources/978-0-31-639860-2">978-0-31-639860-2</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol4" style="text-align: center; font-weight: normal; background-color: transparent;">4</th><td> September 25, 2013<sup id="cite_ref-22" class="reference"><a href="#cite_note-22">[15]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-84-015417-8" title="Special:BookSources/978-4-84-015417-8">978-4-84-015417-8</a></td><td>August 22, 2017<sup id="cite_ref-23" class="reference"><a href="#cite_note-23">[16]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-0-31-639862-6" title="Special:BookSources/978-0-31-639862-6">978-0-31-639862-6</a></td></tr>
</tbody></table>`
	wrongWayToHealTable = `<table class="wikitable" width="100%" style="">


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
</th></tr>


<tr style="text-align: center;"><th scope="row" id="vol1" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">1</th><td> March 25, 2016<sup id="cite_ref-8" class="reference"><a href="#cite_note-8">[8]</a></sup></td><td><style data-mw-deduplicate="TemplateStyles:r1133582631">.mw-parser-output cite.citation{font-style:inherit;word-wrap:break-word}.mw-parser-output .citation q{quotes:"\"""\"""'""'"}.mw-parser-output .citation:target{background-color:rgba(0,127,255,0.133)}.mw-parser-output .id-lock-free a,.mw-parser-output .citation .cs1-lock-free a{background:url("//upload.wikimedia.org/wikipedia/commons/6/65/Lock-green.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-limited a,.mw-parser-output .id-lock-registration a,.mw-parser-output .citation .cs1-lock-limited a,.mw-parser-output .citation .cs1-lock-registration a{background:url("//upload.wikimedia.org/wikipedia/commons/d/d6/Lock-gray-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .id-lock-subscription a,.mw-parser-output .citation .cs1-lock-subscription a{background:url("//upload.wikimedia.org/wikipedia/commons/a/aa/Lock-red-alt-2.svg")right 0.1em center/9px no-repeat}.mw-parser-output .cs1-ws-icon a{background:url("//upload.wikimedia.org/wikipedia/commons/4/4c/Wikisource-logo.svg")right 0.1em center/12px no-repeat}.mw-parser-output .cs1-code{color:inherit;background:inherit;border:none;padding:inherit}.mw-parser-output .cs1-hidden-error{display:none;color:#d33}.mw-parser-output .cs1-visible-error{color:#d33}.mw-parser-output .cs1-maint{display:none;color:#3a3;margin-left:0.3em}.mw-parser-output .cs1-format{font-size:95%}.mw-parser-output .cs1-kern-left{padding-left:0.2em}.mw-parser-output .cs1-kern-right{padding-right:0.2em}.mw-parser-output .citation .mw-selflink{font-weight:inherit}</style><style class="darkreader darkreader--sync" media="screen"></style><a href="/wiki/Special:BookSources/978-4-04-068185-6" title="Special:BookSources/978-4-04-068185-6">978-4-04-068185-6</a></td><td>August 23, 2022<sup id="cite_ref-9" class="reference"><a href="#cite_note-9">[9]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64273-200-9" title="Special:BookSources/978-1-64273-200-9">978-1-64273-200-9</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol2" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">2</th><td> June 24, 2016<sup id="cite_ref-10" class="reference"><a href="#cite_note-10">[10]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-068427-7" title="Special:BookSources/978-4-04-068427-7">978-4-04-068427-7</a></td><td>May 15, 2023<sup id="cite_ref-11" class="reference"><a href="#cite_note-11">[11]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64273-232-0" title="Special:BookSources/978-1-64273-232-0">978-1-64273-232-0</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol3" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">3</th><td> September 23, 2016<sup id="cite_ref-12" class="reference"><a href="#cite_note-12">[12]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-068636-3" title="Special:BookSources/978-4-04-068636-3">978-4-04-068636-3</a></td><td>August 22, 2023</td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-1-64273-286-3" title="Special:BookSources/978-1-64273-286-3">978-1-64273-286-3</a></td></tr>
<tr style="text-align: center;"><th scope="row" id="vol4" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">4</th><td> January 25, 2017<sup id="cite_ref-13" class="reference"><a href="#cite_note-13">[13]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-069054-4" title="Special:BookSources/978-4-04-069054-4">978-4-04-069054-4</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol5" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">5</th><td> April 25, 2017<sup id="cite_ref-14" class="reference"><a href="#cite_note-14">[14]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-069191-6" title="Special:BookSources/978-4-04-069191-6">978-4-04-069191-6</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol6" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">6</th><td> September 25, 2017<sup id="cite_ref-15" class="reference"><a href="#cite_note-15">[15]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-069498-6" title="Special:BookSources/978-4-04-069498-6">978-4-04-069498-6</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol7" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">7</th><td> February 24, 2018<sup id="cite_ref-16" class="reference"><a href="#cite_note-16">[16]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-069727-7" title="Special:BookSources/978-4-04-069727-7">978-4-04-069727-7</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol8" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">8</th><td> July 25, 2018<sup id="cite_ref-17" class="reference"><a href="#cite_note-17">[17]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-065023-4" title="Special:BookSources/978-4-04-065023-4">978-4-04-065023-4</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol9" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">9</th><td> November 24, 2018<sup id="cite_ref-18" class="reference"><a href="#cite_note-18">[18]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-065306-8" title="Special:BookSources/978-4-04-065306-8">978-4-04-065306-8</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol10" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">10</th><td> April 25, 2019<sup id="cite_ref-19" class="reference"><a href="#cite_note-19">[19]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-065682-3" title="Special:BookSources/978-4-04-065682-3">978-4-04-065682-3</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol11" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">11</th><td> October 25, 2019<sup id="cite_ref-20" class="reference"><a href="#cite_note-20">[20]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-064060-0" title="Special:BookSources/978-4-04-064060-0">978-4-04-064060-0</a></td><td>—</td><td>—</td></tr>
<tr style="text-align: center;"><th scope="row" id="vol12" style="text-align: center; font-weight: normal; background-color: transparent; --darkreader-inline-bgcolor: transparent;" data-darkreader-inline-bgcolor="">12</th><td> March 25, 2020<sup id="cite_ref-21" class="reference"><a href="#cite_note-21">[21]</a></sup></td><td><link rel="mw-deduplicated-inline-style" href="mw-data:TemplateStyles:r1133582631"><a href="/wiki/Special:BookSources/978-4-04-064538-4" title="Special:BookSources/978-4-04-064538-4">978-4-04-064538-4</a></td><td>—</td><td>—</td></tr>
</tbody></table>`
)

var ParseWikipediaTableToVolumeInfoTestCases = map[string]ParseWikipediaTableToVolumeInfoTestCase{
	"a simple file with 6 columns and an unreleased volume with no announced date is handled correctly": {
		InputTableHtml:  mushokuTensieTable,
		InputNamePrefix: "Mushoku Tensei",
		ExpectedVolumeInfo: []wikipedia.VolumeInfo{
			{
				Name:        "Mushoku Tensei Vol. 1",
				ReleaseDate: getDatePointer(2019, 4, time.April),
			},
			{
				Name:        "Mushoku Tensei Vol. 2",
				ReleaseDate: getDatePointer(2019, 23, time.May),
			},
			{
				Name:        "Mushoku Tensei Vol. 26",
				ReleaseDate: nil,
			},
		},
	},
	"a simple table with 5 rows gets properly handled": {
		InputTableHtml:  skeletonKnightTable,
		InputNamePrefix: "Skeleton Knight",
		ExpectedVolumeInfo: []wikipedia.VolumeInfo{
			{
				Name:        "Skeleton Knight Vol. 1",
				ReleaseDate: getDatePointer(2019, 28, time.March),
			},
			{
				Name:        "Skeleton Knight Vol. 2",
				ReleaseDate: getDatePointer(2019, 20, time.June),
			},
			{
				Name:        "Skeleton Knight Vol. 3",
				ReleaseDate: getDatePointer(2019, 8, time.August),
			},
			{
				Name:        "Skeleton Knight Vol. 4",
				ReleaseDate: getDatePointer(2019, 14, time.November),
			},
			{
				Name:        "Skeleton Knight Vol. 5",
				ReleaseDate: getDatePointer(2020, 27, time.February),
			},
			{
				Name:        "Skeleton Knight Vol. 6",
				ReleaseDate: getDatePointer(2020, 4, time.June),
			},
			{
				Name:        "Skeleton Knight Vol. 7",
				ReleaseDate: getDatePointer(2020, 20, time.August),
			},
			{
				Name:        "Skeleton Knight Vol. 8",
				ReleaseDate: getDatePointer(2020, 5, time.November),
			},
			{
				Name:        "Skeleton Knight Vol. 9",
				ReleaseDate: getDatePointer(2021, 27, time.May),
			},
			{
				Name:        "Skeleton Knight Vol. 10",
				ReleaseDate: getDatePointer(2023, 9, time.February),
			},
			{
				Name:        "Skeleton Knight Vol. 11",
				ReleaseDate: nil,
			},
		},
	},
	"a simple table with 4 rows gets properly handled": {
		InputTableHtml:  asteriskWarTable,
		InputNamePrefix: "Asterisk War",
		ExpectedVolumeInfo: []wikipedia.VolumeInfo{
			{
				Name:        "Asterisk War Vol. 1",
				ReleaseDate: getDatePointer(2016, 30, time.August),
			},
			{
				Name:        "Asterisk War Vol. 2",
				ReleaseDate: getDatePointer(2016, 20, time.December),
			},
			{
				Name:        "Asterisk War Vol. 3",
				ReleaseDate: getDatePointer(2017, 18, time.April),
			},
			{
				Name:        "Asterisk War Vol. 4",
				ReleaseDate: getDatePointer(2017, 22, time.August),
			},
		},
	},
	// "wrong way to heal should be properly parsed...": {
	// 	InputTableHtml:  wrongWayToHealTable,
	// 	InputNamePrefix: "The Wrong Way to Use Healing Magic",
	// 	ExpectedVolumeInfo: []wikipedia.VolumeInfo{
	// 		{
	// 			Name:        "The Wrong Way to Use Healing Magic Vol. 1",
	// 			ReleaseDate: getDatePointer(2022, 23, time.August),
	// 		},
	// 		{
	// 			Name:        "The Wrong Way to Use Healing Magic Vol. 2",
	// 			ReleaseDate: getDatePointer(2023, 15, time.May),
	// 		},
	// 		{
	// 			Name:        "The Wrong Way to Use Healing Magic Vol. 3",
	// 			ReleaseDate: getDatePointer(2023, 22, time.August),
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 4",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 5",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 6",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 7",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 8",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 9",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 10",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 11",
	// 		},
	// 		{
	// 			Name: "The Wrong Way to Use Healing Magic Vol. 12",
	// 		},
	// 	},
	// },
}

func TestParseWikipediaTableToVolumeInfo(t *testing.T) {
	for name, args := range ParseWikipediaTableToVolumeInfoTestCases {
		t.Run(name, func(t *testing.T) {
			actualVolumeInfo := wikipedia.ParseWikipediaTableToVolumeInfo(args.InputNamePrefix, args.InputTableHtml)
			assert.Equal(t, len(args.ExpectedVolumeInfo), len(actualVolumeInfo))

			for i, volume := range args.ExpectedVolumeInfo {
				assert.Equal(t, volume.Name, actualVolumeInfo[i].Name)
				assert.Equal(t, volume.ReleaseDate, actualVolumeInfo[i].ReleaseDate)
			}
		})
	}
}

func getDatePointer(year, day int, month time.Month) *time.Time {
	date := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return &date
}
