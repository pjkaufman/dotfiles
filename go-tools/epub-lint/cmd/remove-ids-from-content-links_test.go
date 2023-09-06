//go:build unit

package cmd_test

type RemoveIdsFromContentLinksTestCase struct {
	InputText       string
	ExpectedText    string
	ExpectedFileMap map[string]int
}

/* TOC sample 1
<h1 class="toc-title">Contents</h1>
<p class="toc-front" id="cover"><a href="../Text/cover.xhtml">Cover</a></p>
<p class="toc-front" id="insert001"><a href="../Text/insert001.xhtml">Insert</a></p>
<p class="toc-front" id="titlepage"><a href="../Text/titlepage.xhtml">Title Page</a></p>
<p class="toc-front" id="toc-copyright"><a href="../Text/copyright.xhtml">Copyright</a></p>
<p class="toc-chapter" id="toc-preface001"><a href="../Text/preface001.xhtml">Epigraph</a></p>
<p class="toc-chapter1" id="toc-chapter001"><a href="../Text/chapter001.xhtml">Prologue: The King of Corpses</a></p>
<p class="toc-chapter" id="toc-chapter004"><a href="../Text/chapter004.xhtml">Chapter 1: Melancholy of Monsters</a></p>
<p class="toc-chapter" id="toc-chapter009"><a href="../Text/chapter009.xhtml">Chapter 2: Citadel of the Swans</a></p>
<p class="toc-chapter" id="toc-chapter014"><a href="../Text/chapter014.xhtml">Chapter 3: Deaf to the Songbirds’ Lament</a></p>
<p class="toc-chapter" id="toc-chapter021"><a href="../Text/chapter021.xhtml">Chapter 4: Ex Machina</a></p>
<p class="toc-chapter" id="toc-chapter028"><a href="../Text/chapter028.xhtml">Epilogue: Flowers Bloom Not on Snowy Fields</a></p>
<p class="toc-chapter1" id="toc-appendix001"><a href="../Text/appendix001.xhtml">Afterword</a></p>
<p class="toc-chapter" id="newsletter1"><a href="../Text/newsletterSignup.xhtml">Yen Newsletter</a></p>
*/

/* TOC sample 2
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
</navMap>
*/

// var RemoveIdsFromContentLinksTestCases = map[string]RemoveIdsFromContentLinksTestCase{
// 	"make sure that html comments are left alone": {
// 		Input:    "<!--this is a comment. comments are not displayed in the browser-->",
// 		Expected: "<!--this is a comment. comments are not displayed in the browser-->",
// 	},
// 	"make sure that two en dashes are replaced with an em dash": {
// 		Input:    "-- test --",
// 		Expected: "— test —",
// 	},
// 	"make sure that three periods with a 0 or 1 spaces between them get cut down to proper ellipsis": {
// 		Input: `
// 		  ...
// 		  . . .
// 		  . ..
// 		  .. .
// 		  .  . .
// 		`,
// 		Expected: `
// 		  …
// 		  …
// 		  …
// 		  …
// 		  .  . .
// 		`,
// 	},
// 	"make sure that a lowercase 'by the by' results in a lowercase 'by the way'": {
// 		Input:    "by the by",
// 		Expected: "by the way",
// 	},
// 	"make sure that an uppercase 'By the by' results in an uppercase 'By the way'": {
// 		Input:    "By the by",
// 		Expected: "By the way",
// 	},
// 	"make sure that an uppercase 'Sneaked' results in an uppercase 'Snuck'": {
// 		Input:    "Sneaked",
// 		Expected: "Snuck",
// 	},
// 	"make sure that a lowercase 'snuck' results in a lowercase 'snuck'": {
// 		Input:    "On his way he sneaked out the door",
// 		Expected: "On his way he snuck out the door",
// 	},
// 	"make sure that single tilde is converted to an exclamation mark": {
// 		Input:    "~wow isn't this a joy~",
// 		Expected: "!wow isn't this a joy!",
// 	},
// 	"make sure that multiple tildes in a row are not converted to an exclamation mark": {
// 		Input:    "~~ is completely ~~~ left alone",
// 		Expected: "~~ is completely ~~~ left alone",
// 	},
// 	"make sure that a lowercase 'a bolt out of the blue' is correctly converted to 'out of the blue": {
// 		Input:    "a bolt out of the blue",
// 		Expected: "out of the blue",
// 	},
// 	"make sure that an uppercase 'A bolt out of the blue' is correctly converted to 'Out of the blue": {
// 		Input:    "A bolt out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
// 		Expected: "Out of the blue, a loud bang went off waking up people who had been sleeping soundly up until that point.",
// 	},
// 	"make sure that a lowercase 'little wonder' is correctly converted to 'no wonder": {
// 		Input:    "little wonder your attempt failed",
// 		Expected: "no wonder your attempt failed",
// 	},
// 	"make sure that an uppercase 'Little wonder' is correctly converted to 'No wonder": {
// 		Input:    "Little wonder, you were outmatched from the start",
// 		Expected: "No wonder, you were outmatched from the start",
// 	},
// }

// func TestRemoveIdsFromContentLinks(t *testing.T) {
// 	for name, args := range RemoveIdsFromContentLinksTestCases {
// 		t.Run(name, func(t *testing.T) {
// 			actual := cmd.RemoveIdsFromContentLinks(args.Input)

// 			if actual != args.Expected {
// 				t.Errorf("output text doesn't match: expected %v, got %v", args.Expected, actual)
// 			}
// 		})
// 	}

// }
