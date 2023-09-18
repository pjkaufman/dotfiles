package linter

import (
	"fmt"
	"regexp"
	"strings"
)

/**
<?xml version="1.0" encoding="utf-8"?>
<package version="3.0" unique-identifier="BookId" xmlns="http://www.idpf.org/2007/opf">
  <metadata xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:opf="http://www.idpf.org/2007/opf" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <dc:title id="title1">Mushoku Tensei: Jobless Reincarnation Vol. 24</dc:title>
    <dc:creator id="id">Rifujin na Magonote</dc:creator>
    <dc:identifier>calibre:17</dc:identifier>
    <dc:identifier>uuid:d26524a6-6710-4c70-a8f1-4b95864f2eed</dc:identifier>
    <dc:identifier id="BookId">9798888439722</dc:identifier>
    <dc:relation>http://sevenseasentertainment.com</dc:relation>
    <dc:identifier>urn:calibre:9798888439722</dc:identifier>
    <dc:language>en</dc:language>
    <dc:publisher>Seven Seas Entertainment</dc:publisher>
    <dc:subject>light novel</dc:subject>
    <meta refines="#title1" property="title-type">main</meta>
    <meta refines="#title1" property="file-as">Mushoku Tensei: Jobless Reincarnation Vol. 24</meta>
    <meta name="cover" content="cover" />
    <meta content="1.9.30" name="Sigil version" />
    <meta property="dcterms:modified">2023-09-02T17:43:45Z</meta>
    <meta refines="#id" property="role" scheme="marc:relators">aut</meta>
    <meta refines="#id" property="file-as">Magonote, Rifujin na</meta>
  </metadata>
  <manifest>
    <item id="CoverPage_html" href="Text/CoverPage.html" media-type="application/xhtml+xml"/>
    <item id="toc" href="Text/TableOfContents.html" media-type="application/xhtml+xml"/>
    <item id="jnovels.xhtml" href="Text/jnovels.xhtml" media-type="application/xhtml+xml"/>
    <item id="section-0001_html" href="Text/section-0001.html" media-type="application/xhtml+xml"/>
    <item id="section-0002_html" href="Text/section-0002.html" media-type="application/xhtml+xml"/>
    <item id="section-0003_html" href="Text/section-0003.html" media-type="application/xhtml+xml"/>
    <item id="section-0004_html" href="Text/section-0004.html" media-type="application/xhtml+xml"/>
    <item id="section-0005_html" href="Text/section-0005.html" media-type="application/xhtml+xml"/>
    <item id="section-0006_html" href="Text/section-0006.html" media-type="application/xhtml+xml"/>
    <item id="section-0007_html" href="Text/section-0007.html" media-type="application/xhtml+xml"/>
    <item id="section-0008_html" href="Text/section-0008.html" media-type="application/xhtml+xml"/>
    <item id="section-0009_html" href="Text/section-0009.html" media-type="application/xhtml+xml"/>
    <item id="section-0010_html" href="Text/section-0010.html" media-type="application/xhtml+xml"/>
    <item id="section-0011_html" href="Text/section-0011.html" media-type="application/xhtml+xml"/>
    <item id="section-0012_html" href="Text/section-0012.html" media-type="application/xhtml+xml"/>
    <item id="section-0013_html" href="Text/section-0013.html" media-type="application/xhtml+xml"/>
    <item id="section-0014_html" href="Text/section-0014.html" media-type="application/xhtml+xml"/>
    <item id="section-0015_html" href="Text/section-0015.html" media-type="application/xhtml+xml"/>
    <item id="section-0016_html" href="Text/section-0016.html" media-type="application/xhtml+xml"/>
    <item id="section-0017_html" href="Text/section-0017.html" media-type="application/xhtml+xml"/>
    <item id="section-0018_html" href="Text/section-0018.html" media-type="application/xhtml+xml"/>
    <item id="navid" href="nav.xhtml" media-type="application/xhtml+xml" properties="nav"/>
    <item id="ncx" href="toc.ncx" media-type="application/x-dtbncx+xml"/>
    <item id="styles_css" href="Styles/styles.css" media-type="text/css"/>
    <item id="x1.png" href="Images/1.png" media-type="image/png"/>
    <item id="COLORGALLERY__jpg" href="Images/COLORGALLERY_.jpg" media-type="image/jpeg"/>
    <item id="COLORGALLERY_1_jpg" href="Images/COLORGALLERY_1.jpg" media-type="image/jpeg"/>
    <item id="COLORGALLERY_2_jpg" href="Images/COLORGALLERY_2.jpg" media-type="image/jpeg"/>
    <item id="cover" href="Images/CoverDesign.jpg" media-type="image/jpeg" properties="cover-image"/>
    <item id="FRONTMATTER__jpg" href="Images/FRONTMATTER_.jpg" media-type="image/jpeg"/>
    <item id="FRONTMATTER_2_jpg" href="Images/FRONTMATTER_2.jpg" media-type="image/jpeg"/>
    <item id="FRONTMATTER_3_jpg" href="Images/FRONTMATTER_3.jpg" media-type="image/jpeg"/>
    <item id="FRONTMATTER_4_jpg" href="Images/FRONTMATTER_4.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES__jpg" href="Images/INTERIORIMAGES_.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_2_jpg" href="Images/INTERIORIMAGES_2.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_3_jpg" href="Images/INTERIORIMAGES_3.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_4_jpg" href="Images/INTERIORIMAGES_4.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_5_jpg" href="Images/INTERIORIMAGES_5.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_6_jpg" href="Images/INTERIORIMAGES_6.jpg" media-type="image/jpeg"/>
    <item id="INTERIORIMAGES_7_jpg" href="Images/INTERIORIMAGES_7.jpg" media-type="image/jpeg"/>
    <item id="sevenseaslogo_jpg" href="Images/sevenseaslogo.jpg" media-type="image/jpeg"/>
  </manifest>
  <spine toc="ncx">
    <itemref idref="CoverPage_html"/>
    <itemref idref="toc"/>
    <itemref idref="jnovels.xhtml"/>
    <itemref idref="section-0001_html"/>
    <itemref idref="section-0002_html"/>
    <itemref idref="section-0003_html"/>
    <itemref idref="section-0004_html"/>
    <itemref idref="section-0005_html"/>
    <itemref idref="section-0006_html"/>
    <itemref idref="section-0007_html"/>
    <itemref idref="section-0008_html"/>
    <itemref idref="section-0009_html"/>
    <itemref idref="section-0010_html"/>
    <itemref idref="section-0011_html"/>
    <itemref idref="section-0012_html"/>
    <itemref idref="section-0013_html"/>
    <itemref idref="section-0014_html"/>
    <itemref idref="section-0015_html"/>
    <itemref idref="section-0016_html"/>
    <itemref idref="section-0017_html"/>
    <itemref idref="section-0018_html"/>
    <itemref idref="navid" linear="no"/>
  </spine>
  <guide>
    <reference type="cover" title="Cover Page" href="Text/CoverPage.html"/>
    <reference type="toc" title="Table of Contents" href="Text/TableOfContents.html#tableofcontents"/>
  </guide>
</package>

*/

type EpubInfo struct {
	HtmlFiles   map[string]struct{}
	ImagesFiles map[string]struct{}
	OtherFiles  map[string]struct{}
	PageIds     []PageIdInfo
	NcxFile     string
	NavFile     string
	TocFile     string
	Version     int
}

const closingManifestEl = "</manifest>"

var openingManifestRegex = regexp.MustCompile("<manifest[^>\n]*>")
var packageVersionRegex = regexp.MustCompile(`<package[^>\n]*version=["']([^"']+)["'][^>\n]*>`)
var itemElRegex = regexp.MustCompile(`<item [^>\n]*>`)
var tocReferenceElRegex = regexp.MustCompile(`<reference [^>\n]*type=["']toc["'][^>\n]*>`)
var hrefAttributeRegex = regexp.MustCompile(`href=["']([^'"]*)(["'])`)
var mediaTypeAttributeRegex = regexp.MustCompile(`media-type=["']([^'"]*)(["'])`)
var isNavFileRegex = regexp.MustCompile(`properties=["']nav["']`)
var ErrNoPackageInfo = fmt.Errorf("no package info found for the epub - please verify that the opf has a version in it")
var ErrNoItemEls = fmt.Errorf("no manifest items found for the epub - please verify that the opf has a items in it")
var ErrNoManifest = fmt.Errorf("no manifest found")
var ErrNoEndOfManifest = fmt.Errorf("manifest is incorrectly formatted since it has no closing manifest element")

func ParseOpfFile(text string) (EpubInfo, error) {
	var epubInfo = EpubInfo{
		HtmlFiles:   make(map[string]struct{}),
		ImagesFiles: make(map[string]struct{}),
		OtherFiles:  make(map[string]struct{}),
		PageIds:     []PageIdInfo{},
	}
	var version, err = getVersion(text)
	if err != nil {
		return epubInfo, err
	}

	epubInfo.Version = version
	err = epubInfo.parseManifest(text)
	if err != nil {
		return epubInfo, err
	}

	err = epubInfo.getTocFile(text)
	if err != nil {
		return epubInfo, err
	}

	return epubInfo, nil
}

func getVersion(text string) (int, error) {
	var packageInfo = packageVersionRegex.FindAllStringSubmatch(text, 1)
	if len(packageInfo) == 0 {
		return 2, ErrNoPackageInfo
	}

	if strings.Contains(packageInfo[0][1], "3.") {
		return 3, nil
	}

	return 2, nil
}

func (ei *EpubInfo) parseManifest(text string) error {
	var startOfManifest = openingManifestRegex.FindStringIndex(text)
	if len(startOfManifest) == 0 {
		return ErrNoManifest
	}

	var endOfManifest = strings.Index(text, closingManifestEl)
	if endOfManifest < 0 {
		return ErrNoEndOfManifest
	}

	var manifest = text[startOfManifest[0]:endOfManifest]
	var itemEls = itemElRegex.FindAllStringIndex(manifest, -1)
	for _, locs := range itemEls {
		var itemEl = manifest[locs[0]:locs[1]]
		hrefInfo := hrefAttributeRegex.FindAllStringSubmatch(itemEl, 1)
		if len(hrefInfo) == 0 {
			return fmt.Errorf("failed to get the href attribute on the item \"%s\"", itemEl)
		}

		var filePath = hrefInfo[0][1]

		mediaTypeInfo := mediaTypeAttributeRegex.FindAllStringSubmatch(itemEl, 1)
		if len(hrefInfo) == 0 {
			return fmt.Errorf("failed to get the media-type attribute on the item \"%s\"", itemEl)
		}
		var mediaType = mediaTypeInfo[0][1]

		if ei.Version == 3 && isNavFileRegex.MatchString(itemEl) {
			ei.NavFile = filePath
		}

		if strings.Contains(mediaType, "xhtml") {
			ei.HtmlFiles[filePath] = struct{}{}
		} else if strings.Contains(mediaType, "image") {
			ei.ImagesFiles[filePath] = struct{}{}
		} else {
			if strings.HasSuffix(filePath, ".ncx") {
				ei.NcxFile = filePath
			}

			ei.OtherFiles[filePath] = struct{}{}
		}

	}
	if len(itemEls) == 0 {
		return ErrNoItemEls
	}

	return nil
}

func (ei *EpubInfo) getTocFile(text string) error {
	var tocInfo = tocReferenceElRegex.FindAllStringSubmatch(text, 1)
	if len(tocInfo) == 0 {
		return nil
	}

	var tocEl = tocInfo[0][0]
	hrefInfo := hrefAttributeRegex.FindAllStringSubmatch(tocEl, 1)
	if len(hrefInfo) == 0 {
		return fmt.Errorf("failed to get the href attribute on the toc reference \"%s\"", tocEl)
	}

	var href = hrefInfo[0][1]
	var hashTagIndex = strings.Index(href, "#")
	if hashTagIndex != -1 {
		href = href[0:hashTagIndex]
	}

	ei.TocFile = href

	return nil
}
