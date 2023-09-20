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
	sampleNavFile2 = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="en" xml:lang="en">
<head>
  <meta content="text/html; charset=UTF-8" http-equiv="default-style"/>
  <title>Tearmoon Empire: Volume 9</title>
  <link href="../Styles/stylesheet.css" rel="stylesheet" type="text/css"/>
</head>

<body>
  <nav epub:type="toc" id="toc"><h1 class="toc-title">Table of Contents</h1>

  <ol class="none" epub:type="list">
    <li class="toc-front"><a href="../Text/cover.xhtml">Cover</a></li>

    <li class="toc-front"><a href="../Text/frontmatter1.xhtml">Color Illustrations</a></li>

    <li class="toc-front"><a href="../Text/characters1.xhtml">Characters</a></li>

    <li class="toc-front"><a href="../Text/map.xhtml">Map</a></li>

    <li style="margin-top: 16px;"><span style="font-weight:bold">Part 4: To the Moon-Led Morrow III</span> 

    <ol class="none" epub:type="list">
      <li class="toc-chapter" id="toc-prologue"><a href="../Text/prologue.xhtml">Prologue: The Doctrine of Empress Mia —Ludwig’s Important Duty— </a></li>

      <li class="toc-chapter" id="toc-chapter1"><a href="../Text/chapter1.xhtml">Chapter 1: The Secret Behind That Dress... </a></li>

      <li class="toc-chapter" id="toc-chapter2"><a href="../Text/chapter2.xhtml">Chapter 2: Princess Mia Thinks...and Thinks... </a></li>

      <li class="toc-chapter" id="toc-chapter3"><a href="../Text/chapter3.xhtml">Chapter 3: Mia...Isn’t Living a Lie!</a></li>

      <li class="toc-chapter" id="toc-chapter4"><a href="../Text/chapter4.xhtml">Chapter 4: Nod, Nod...Nod?</a></li>

      <li class="toc-chapter" id="toc-chapter5"><a href="../Text/chapter5.xhtml">Chapter 5: Fond Visions of Four-Eyes Past</a></li>

      <li class="toc-chapter" id="toc-chapter6"><a href="../Text/chapter6.xhtml">Chapter 6: Princess Mia...Speaks with a Bit of an Accent</a></li>

      <li class="toc-chapter" id="toc-chapter7"><a href="../Text/chapter7.xhtml">Chapter 7: Spark</a></li>

      <li class="toc-chapter" id="toc-chapter8"><a href="../Text/chapter8.xhtml">Chapter 8: A World of F.A.T.</a></li>

      <li class="toc-chapter" id="toc-chapter9"><a href="../Text/chapter9.xhtml">Chapter 9: Mia Channels Her Inner Seductress</a></li>

      <li class="toc-chapter" id="toc-chapter10"><a href="../Text/chapter10.xhtml">Chapter 10: The Dance of Her Dreams, and Then...</a></li>

      <li class="toc-chapter" id="toc-chapter11"><a href="../Text/chapter11.xhtml">Chapter 11: A Turn for the Worse</a></li>

      <li class="toc-chapter" id="toc-chapter12"><a href="../Text/chapter12.xhtml">Chapter 12: An Ancient Poison —Coincidental Encounter—</a></li>

      <li class="toc-chapter" id="toc-chapter13"><a href="../Text/chapter13.xhtml">Chapter 13: Two Noble Girls Who Dared Rage...and One Princess Who Definitely Didn’t Dare and Must Watch with Horror!</a></li>

      <li class="toc-chapter" id="toc-chapter14"><a href="../Text/chapter14.xhtml">Chapter 14: Citrina...Finally Sees the True Nature of the Great Sage of the Empire...or Does She?</a></li>

      <li class="toc-chapter" id="toc-chapter15"><a href="../Text/chapter15.xhtml">Chapter 15: Miabel...Goes on an Adven—Tour of the Castle!</a></li>

      <li class="toc-chapter" id="toc-chapter16"><a href="../Text/chapter16.xhtml">Chapter 16: Time of Judgment, Friend or Foe...</a></li>

      <li class="toc-chapter" id="toc-chapter17"><a href="../Text/chapter17.xhtml">Chapter 17: Princess Mia...Rides a Wave in a Breakthrough</a></li>

      <li class="toc-chapter" id="toc-chapter18"><a href="../Text/chapter18.xhtml">Chapter 18: Swoosh, Swoosh!</a></li>

      <li class="toc-chapter" id="toc-chapter19"><a href="../Text/chapter19.xhtml">Chapter 19: The Miabel Dialogue: What Is a King?</a></li>

      <li class="toc-chapter" id="toc-chapter20"><a href="../Text/chapter20.xhtml">Chapter 20: Princess Mia...Is Left Bereft!</a></li>

      <li class="toc-chapter" id="toc-chapter21"><a href="../Text/chapter21.xhtml">Chapter 21: A Delivery from a Loyal Subject</a></li>

      <li class="toc-chapter" id="toc-chapter22"><a href="../Text/chapter22.xhtml">Chapter 22: Princess Mia...Finally Decides on Revenge!</a></li>

      <li class="toc-chapter" id="toc-chapter23"><a href="../Text/chapter23.xhtml">Chapter 23: Sion’s Decision</a></li>

      <li class="toc-chapter" id="toc-chapter24"><a href="../Text/chapter24.xhtml">Chapter 24: Someone Perfect for You Is Surely...</a></li>

      <li class="toc-chapter" id="toc-chapter25"><a href="../Text/chapter25.xhtml">Chapter 25: As King and as Father</a></li>

      <li class="toc-chapter" id="toc-chapter26"><a href="../Text/chapter26.xhtml">Chapter 26: The Night When Love Was Lost</a></li>

      <li class="toc-chapter" id="toc-chapter27"><a href="../Text/chapter27.xhtml">Chapter 27: The Brothers Embark</a></li>

      <li class="toc-chapter" id="toc-side1"><a href="../Text/side1.xhtml">Side Chapter: Mialogy —Ludwig’s Unpalatably Delicious Wine—</a></li>

      <li class="toc-chapter" id="toc-side2"><a href="../Text/side2.xhtml">Side Chapter: Empress Mia’s Trick or Treat</a></li>

      <li class="toc-chapter" id="toc-side3"><a href="../Text/side3.xhtml">Side Chapter: The Seed That Did Not Sprout</a></li>

      <li class="toc-chapter" id="toc-chapter28"><a href="../Text/chapter28.xhtml">Chapter 28: A Delightful Gathering of Girls —Princess Mia Is Filled with a Sense of Duty—</a></li>

      <li class="toc-chapter" id="toc-chapter29"><a href="../Text/chapter29.xhtml">Chapter 29: Miabel’s Nearly Spine-chilling Waking of Suspense and Horror</a></li>

      <li class="toc-chapter" id="toc-chapter30"><a href="../Text/chapter30.xhtml">Chapter 30: Bel Asks the Hard Questions</a></li>

      <li class="toc-chapter" id="toc-chapter31"><a href="../Text/chapter31.xhtml">Chapter 31: Onward! To the Equestrian Kingdom!</a></li>

      <li class="toc-chapter" id="toc-chapter32"><a href="../Text/chapter32.xhtml">Chapter 32: The High Priestess of the Serpents Dances</a></li>
    </ol></li>

    <li class="toc-chapter" id="toc-extra" style="margin-top: 16px;"><a href="../Text/extra.xhtml">Chancellor Ludwig Loves His Wine</a></li>

    <li class="toc-chapter" id="toc-diary"><a href="../Text/diary.xhtml">Mia’s Diary of Deeper Dining —Three Exquisite Sunkland Dishes for Those Who Wish to Deepen Their Appreciation of Food—</a></li>

    <li class="toc-appendix" style="margin-top: 16px;" id="toc-afterword"><a href="../Text/afterword.xhtml">Afterword</a></li>

    <li class="toc-appendix" id="toc-bonus"><a href="../Text/bonus.xhtml">Bonus Short Story</a></li>

    <li class="toc-appendix" id="toc-signup"><a href="../Text/signup.xhtml">About J-Novel Club</a></li>

    <li class="toc-appendix" id="toc-copyright"><a href="../Text/copyright.xhtml">Copyright</a></li>
  </ol></nav>

  <nav epub:type="landmarks" id="landmarks" hidden=""><h1>Landmarks</h1>

  <ol>
    <li><a epub:type="toc" href="../Text/frontmatter1.xhtml">Color Illustrations</a></li>

    <li><a epub:type="toc" href="../Text/toc.xhtml">Table of Contents</a></li>
  </ol></nav>
</body>
</html>`
	expectedSampleNavFile2 = `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops" lang="en" xml:lang="en">
<head>
  <meta content="text/html; charset=UTF-8" http-equiv="default-style"/>
  <title>Tearmoon Empire: Volume 9</title>
  <link href="../Styles/stylesheet.css" rel="stylesheet" type="text/css"/>
</head>

<body>
  <nav epub:type="toc" id="toc"><h1 class="toc-title">Table of Contents</h1>

  <ol class="none" epub:type="list">
    <li class="toc-front"><a href="../Text/cover.xhtml">Cover</a></li>

    <li class="toc-front"><a href="../Text/frontmatter1.xhtml">Color Illustrations</a></li>

    <li class="toc-front"><a href="../Text/characters1.xhtml">Characters</a></li>

    <li class="toc-front"><a href="../Text/map.xhtml">Map</a></li>

    <li style="margin-top: 16px;"><span style="font-weight:bold">Part 4: To the Moon-Led Morrow III</span> 

    <ol class="none" epub:type="list">
      <li class="toc-chapter" id="toc-prologue"><a href="../Text/prologue.xhtml">Prologue: The Doctrine of Empress Mia —Ludwig’s Important Duty— </a></li>

      <li class="toc-chapter" id="toc-chapter1"><a href="../Text/chapter1.xhtml">Chapter 1: The Secret Behind That Dress... </a></li>

      <li class="toc-chapter" id="toc-chapter2"><a href="../Text/chapter2.xhtml">Chapter 2: Princess Mia Thinks...and Thinks... </a></li>

      <li class="toc-chapter" id="toc-chapter3"><a href="../Text/chapter3.xhtml">Chapter 3: Mia...Isn’t Living a Lie!</a></li>

      <li class="toc-chapter" id="toc-chapter4"><a href="../Text/chapter4.xhtml">Chapter 4: Nod, Nod...Nod?</a></li>

      <li class="toc-chapter" id="toc-chapter5"><a href="../Text/chapter5.xhtml">Chapter 5: Fond Visions of Four-Eyes Past</a></li>

      <li class="toc-chapter" id="toc-chapter6"><a href="../Text/chapter6.xhtml">Chapter 6: Princess Mia...Speaks with a Bit of an Accent</a></li>

      <li class="toc-chapter" id="toc-chapter7"><a href="../Text/chapter7.xhtml">Chapter 7: Spark</a></li>

      <li class="toc-chapter" id="toc-chapter8"><a href="../Text/chapter8.xhtml">Chapter 8: A World of F.A.T.</a></li>

      <li class="toc-chapter" id="toc-chapter9"><a href="../Text/chapter9.xhtml">Chapter 9: Mia Channels Her Inner Seductress</a></li>

      <li class="toc-chapter" id="toc-chapter10"><a href="../Text/chapter10.xhtml">Chapter 10: The Dance of Her Dreams, and Then...</a></li>

      <li class="toc-chapter" id="toc-chapter11"><a href="../Text/chapter11.xhtml">Chapter 11: A Turn for the Worse</a></li>

      <li class="toc-chapter" id="toc-chapter12"><a href="../Text/chapter12.xhtml">Chapter 12: An Ancient Poison —Coincidental Encounter—</a></li>

      <li class="toc-chapter" id="toc-chapter13"><a href="../Text/chapter13.xhtml">Chapter 13: Two Noble Girls Who Dared Rage...and One Princess Who Definitely Didn’t Dare and Must Watch with Horror!</a></li>

      <li class="toc-chapter" id="toc-chapter14"><a href="../Text/chapter14.xhtml">Chapter 14: Citrina...Finally Sees the True Nature of the Great Sage of the Empire...or Does She?</a></li>

      <li class="toc-chapter" id="toc-chapter15"><a href="../Text/chapter15.xhtml">Chapter 15: Miabel...Goes on an Adven—Tour of the Castle!</a></li>

      <li class="toc-chapter" id="toc-chapter16"><a href="../Text/chapter16.xhtml">Chapter 16: Time of Judgment, Friend or Foe...</a></li>

      <li class="toc-chapter" id="toc-chapter17"><a href="../Text/chapter17.xhtml">Chapter 17: Princess Mia...Rides a Wave in a Breakthrough</a></li>

      <li class="toc-chapter" id="toc-chapter18"><a href="../Text/chapter18.xhtml">Chapter 18: Swoosh, Swoosh!</a></li>

      <li class="toc-chapter" id="toc-chapter19"><a href="../Text/chapter19.xhtml">Chapter 19: The Miabel Dialogue: What Is a King?</a></li>

      <li class="toc-chapter" id="toc-chapter20"><a href="../Text/chapter20.xhtml">Chapter 20: Princess Mia...Is Left Bereft!</a></li>

      <li class="toc-chapter" id="toc-chapter21"><a href="../Text/chapter21.xhtml">Chapter 21: A Delivery from a Loyal Subject</a></li>

      <li class="toc-chapter" id="toc-chapter22"><a href="../Text/chapter22.xhtml">Chapter 22: Princess Mia...Finally Decides on Revenge!</a></li>

      <li class="toc-chapter" id="toc-chapter23"><a href="../Text/chapter23.xhtml">Chapter 23: Sion’s Decision</a></li>

      <li class="toc-chapter" id="toc-chapter24"><a href="../Text/chapter24.xhtml">Chapter 24: Someone Perfect for You Is Surely...</a></li>

      <li class="toc-chapter" id="toc-chapter25"><a href="../Text/chapter25.xhtml">Chapter 25: As King and as Father</a></li>

      <li class="toc-chapter" id="toc-chapter26"><a href="../Text/chapter26.xhtml">Chapter 26: The Night When Love Was Lost</a></li>

      <li class="toc-chapter" id="toc-chapter27"><a href="../Text/chapter27.xhtml">Chapter 27: The Brothers Embark</a></li>

      <li class="toc-chapter" id="toc-side1"><a href="../Text/side1.xhtml">Side Chapter: Mialogy —Ludwig’s Unpalatably Delicious Wine—</a></li>

      <li class="toc-chapter" id="toc-side2"><a href="../Text/side2.xhtml">Side Chapter: Empress Mia’s Trick or Treat</a></li>

      <li class="toc-chapter" id="toc-side3"><a href="../Text/side3.xhtml">Side Chapter: The Seed That Did Not Sprout</a></li>

      <li class="toc-chapter" id="toc-chapter28"><a href="../Text/chapter28.xhtml">Chapter 28: A Delightful Gathering of Girls —Princess Mia Is Filled with a Sense of Duty—</a></li>

      <li class="toc-chapter" id="toc-chapter29"><a href="../Text/chapter29.xhtml">Chapter 29: Miabel’s Nearly Spine-chilling Waking of Suspense and Horror</a></li>

      <li class="toc-chapter" id="toc-chapter30"><a href="../Text/chapter30.xhtml">Chapter 30: Bel Asks the Hard Questions</a></li>

      <li class="toc-chapter" id="toc-chapter31"><a href="../Text/chapter31.xhtml">Chapter 31: Onward! To the Equestrian Kingdom!</a></li>

      <li class="toc-chapter" id="toc-chapter32"><a href="../Text/chapter32.xhtml">Chapter 32: The High Priestess of the Serpents Dances</a></li>
    </ol></li>

    <li class="toc-chapter" id="toc-extra" style="margin-top: 16px;"><a href="../Text/extra.xhtml">Chancellor Ludwig Loves His Wine</a></li>

    <li class="toc-chapter" id="toc-diary"><a href="../Text/diary.xhtml">Mia’s Diary of Deeper Dining —Three Exquisite Sunkland Dishes for Those Who Wish to Deepen Their Appreciation of Food—</a></li>

    <li class="toc-appendix" style="margin-top: 16px;" id="toc-afterword"><a href="../Text/afterword.xhtml">Afterword</a></li>

    <li class="toc-appendix" id="toc-bonus"><a href="../Text/bonus.xhtml">Bonus Short Story</a></li>

    <li class="toc-appendix" id="toc-signup"><a href="../Text/signup.xhtml">About J-Novel Club</a></li>

    <li class="toc-appendix" id="toc-copyright"><a href="../Text/copyright.xhtml">Copyright</a></li>
  </ol></nav>

  <nav epub:type="landmarks" id="landmarks" hidden=""><h1>Landmarks</h1>

  <ol>
    <li><a epub:type="toc" href="../Text/frontmatter1.xhtml">Color Illustrations</a></li>

    <li><a epub:type="toc" href="../Text/toc.xhtml">Table of Contents</a></li>
  </ol></nav>
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
	"make sure that empty list items are ignored": {
		InputText:    sampleNavFile2,
		ExpectedText: expectedSampleNavFile2,
		expectedErr:  nil,
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
