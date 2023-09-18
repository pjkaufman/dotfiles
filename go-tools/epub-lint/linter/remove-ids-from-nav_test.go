//go:build unit

package linter_test

import (
	"fmt"
	"testing"

	"github.com/pjkaufman/dotfiles/go-tools/epub-lint/linter"
	"github.com/stretchr/testify/assert"
)

const (
	sampleNavFile1 = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="en" xml:lang="en">
<head>
<meta content="text/html; charset=utf-8" http-equiv="default-style"/>
<title>The Asterisk War, Vol 1: The Academy City on the Water</title>
<link rel="stylesheet" href="css/stylesheet.css" type="text/css"/>

<!-- kobo-style -->
<script xmlns="http://www.w3.org/1999/xhtml" type="text/javascript" src="../js/kobo.js"></script>

</head>
<body>
<nav id="toc" epub:type="toc">
<h1 class="toc-title">Contents</h1>
<ol class="none" epub:type="list">
<li class="toc-front" id="cover"><a href="cover.xhtml">Cover</a></li>
<li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
<li class="toc-appendix" id="toc-photo-insert"><a href="photo-insert.xhtml">Insert</a></li>
<li class="toc-front" id="titlepage"><a href="titlepage.xhtml">Title Page</a></li>
<li class="toc-front" id="preface001"><a href="preface001.xhtml">Map</a></li>
<li class="toc-chapter1" id="id_1_5"><a id="Ref_8204a" href="chapter001.xhtml#Ref_8204">Chapter 1: Glühen Rose</a></li>
<li class="toc-chapter" id="id_1_7"><a id="Ref_8207a" href="chapter004.xhtml#Ref_8207">Chapter 2: Asterisk, The City of Academic Warfare</a></li>
<li class="toc-chapter" id="id_1_11"><a id="Ref_8210a" href="chapter009.xhtml#Ref_8210">Chapter 3: Her Noble Eyes</a></li>
<li class="toc-chapter" id="id_1_13"><a id="Ref_8213a" href="chapter012.xhtml#Ref_8213">Chapter 4: Reminiscence and Reunion</a></li>
<li class="toc-chapter" id="id_1_16"><a id="Ref_8216a" href="chapter016.xhtml#Ref_8216">Chapter 5: The Ser Veresta</a></li>
<li class="toc-chapter" id="id_1_18"><a id="Ref_8219a" href="chapter019.xhtml#Ref_8219">Chapter 6: A Holiday For Two</a></li>
<li class="toc-chapter" id="id_1_20"><a id="Ref_8222a" href="chapter022.xhtml#Ref_8222">Chapter 7: Unchained</a></li>
<li class="toc-chapter" id="id_1_23"><a id="Ref_8228a" href="chapter027.xhtml#Ref_8228">Epilogue</a></li>
<li class="toc-appendix1" id="id_1_24"><a id="Ref_8225a" href="chapter028.xhtml#Ref_8225">Afterword</a></li>
<li class="toc-appendix" id="id_1_25"><a id="Ref_8225aa" href="chapter047.xhtml">Manga Preview</a></li>
<li class="toc-appendix" id="Newsletters"><a href="newsletterSignup.xhtml">Yen Newsletter</a></li>
<li class="toc-appendix" id="toc-copyright"><a href="copyright.xhtml">Copyright</a></li>
</ol>
</nav>
<nav epub:type="landmarks" class="hidden-tag" hidden="hidden">
<h1>Navigation</h1>
<ol class="none" epub:type="list">
<li><a epub:type="bodymatter" href="welcome.xhtml">Begin Reading</a></li>
<li><a epub:type="toc" href="toc.xhtml">Table of Contents</a></li>
</ol>
</nav>
</body>
</html>`
	expectedSampleNavFile1 = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="en" xml:lang="en">
<head>
<meta content="text/html; charset=utf-8" http-equiv="default-style"/>
<title>The Asterisk War, Vol 1: The Academy City on the Water</title>
<link rel="stylesheet" href="css/stylesheet.css" type="text/css"/>

<!-- kobo-style -->
<script xmlns="http://www.w3.org/1999/xhtml" type="text/javascript" src="../js/kobo.js"></script>

</head>
<body>
<nav id="toc" epub:type="toc">
<h1 class="toc-title">Contents</h1>
<ol class="none" epub:type="list">
<li class="toc-front" id="cover"><a href="cover.xhtml">Cover</a></li>
<li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
<li class="toc-appendix" id="toc-photo-insert"><a href="photo-insert.xhtml">Insert</a></li>
<li class="toc-front" id="titlepage"><a href="titlepage.xhtml">Title Page</a></li>
<li class="toc-front" id="preface001"><a href="preface001.xhtml">Map</a></li>
<li class="toc-chapter1" id="id_1_5"><a id="Ref_8204a" href="chapter001.xhtml">Chapter 1: Glühen Rose</a></li>
<li class="toc-chapter" id="id_1_7"><a id="Ref_8207a" href="chapter004.xhtml">Chapter 2: Asterisk, The City of Academic Warfare</a></li>
<li class="toc-chapter" id="id_1_11"><a id="Ref_8210a" href="chapter009.xhtml">Chapter 3: Her Noble Eyes</a></li>
<li class="toc-chapter" id="id_1_13"><a id="Ref_8213a" href="chapter012.xhtml">Chapter 4: Reminiscence and Reunion</a></li>
<li class="toc-chapter" id="id_1_16"><a id="Ref_8216a" href="chapter016.xhtml">Chapter 5: The Ser Veresta</a></li>
<li class="toc-chapter" id="id_1_18"><a id="Ref_8219a" href="chapter019.xhtml">Chapter 6: A Holiday For Two</a></li>
<li class="toc-chapter" id="id_1_20"><a id="Ref_8222a" href="chapter022.xhtml">Chapter 7: Unchained</a></li>
<li class="toc-chapter" id="id_1_23"><a id="Ref_8228a" href="chapter027.xhtml">Epilogue</a></li>
<li class="toc-appendix1" id="id_1_24"><a id="Ref_8225a" href="chapter028.xhtml">Afterword</a></li>
<li class="toc-appendix" id="id_1_25"><a id="Ref_8225aa" href="chapter047.xhtml">Manga Preview</a></li>
<li class="toc-appendix" id="Newsletters"><a href="newsletterSignup.xhtml">Yen Newsletter</a></li>
<li class="toc-appendix" id="toc-copyright"><a href="copyright.xhtml">Copyright</a></li>
</ol>
</nav>
<nav epub:type="landmarks" class="hidden-tag" hidden="hidden">
<h1>Navigation</h1>
<ol class="none" epub:type="list">
<li><a epub:type="bodymatter" href="welcome.xhtml">Begin Reading</a></li>
<li><a epub:type="toc" href="toc.xhtml">Table of Contents</a></li>
</ol>
</nav>
</body>
</html>`
)

type RemoveIdsFromNavTestCase struct {
	InputText    string
	ExpectedText string
	expectedErr  error
}

var RemoveIdsFromNavTestCases = map[string]RemoveIdsFromNavTestCase{
	"make sure that list items with the ids removed": {
		InputText:    sampleNavFile1,
		ExpectedText: expectedSampleNavFile1,
		expectedErr:  nil,
	},
	"make sure that a nav with list item anchor tag hrefs pointing to the same file throws an error": {
		InputText: `<nav epub:type="toc">
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a href="welcome.xhtml">Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		ExpectedText: `<nav epub:type="toc">
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a href="welcome.xhtml">Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		expectedErr: fmt.Errorf(`there is more than one reference to "%s". Consider restructuring the epub to fix up the nav`, "welcome.xhtml"),
	},
	"make sure no nav results in no error being thrown": {
		InputText:    "",
		ExpectedText: "",
		expectedErr:  nil,
	},
	"make sure no nav ending element results in an error being thrown": {
		InputText:    "<nav epub:type=\"toc\">",
		ExpectedText: "<nav epub:type=\"toc\">",
		expectedErr:  linter.ErrNoEndOfNav,
	},
	"make sure that a nav with a list item anchor without an href for the anchor element throws an error": {
		InputText: `<nav epub:type="toc">
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a>Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		ExpectedText: `<nav epub:type="toc">
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a>Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		expectedErr: fmt.Errorf("possible problem with list anchor tag href: []"),
	},
	"make sure that a nav without the epub type of toc is ignored": {
		InputText: `<nav>
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a>Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		ExpectedText: `<nav>
  <ol class="none" epub:type="list">
    <li class="toc-front" id="cover"><a>Cover</a></li>
    <li class="toc-front" id="toc-welcome"><a href="welcome.xhtml">Welcome</a></li>
  </ol>
</nav>`,
		expectedErr: nil,
	},
}

func TestRemoveIdsFromNav(t *testing.T) {
	for name, args := range RemoveIdsFromNavTestCases {
		t.Run(name, func(t *testing.T) {
			actual, err := linter.RemoveIdsFromNav(args.InputText)

			assert.Equal(t, args.ExpectedText, actual)
			assert.Equal(t, args.expectedErr, err)
		})
	}
}
