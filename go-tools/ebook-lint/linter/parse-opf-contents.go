package linter

import (
	"fmt"
	"regexp"
	"strings"
)

type EpubInfo struct {
	HtmlFiles   map[string]struct{}
	ImagesFiles map[string]struct{}
	CssFiles    map[string]struct{}
	OtherFiles  map[string]struct{}
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
		CssFiles:    make(map[string]struct{}),
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
		} else if strings.Contains(mediaType, "css") {
			ei.CssFiles[filePath] = struct{}{}
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
