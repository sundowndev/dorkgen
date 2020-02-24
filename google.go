package dorkgen

import (
	"net/url"
	"strings"
)

const (
	searchUrl = "https://www.google.com/search"
	siteTag     = "site:"
	urlTag      = "inurl:"
	filetypeTag = "filetype:"
	cacheTag    = "cache:"
	relatedTag  = "related:"
	extTag      = "ext:"
	excludeTag  = "-"
)

// GoogleSearch ...
type GoogleSearch struct {
	EngineFactory
	tags []string
}

func concat(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
}

// ToString ...
func (e *GoogleSearch) ToString() string {
	return strings.Join(e.tags, " ")
}

// ToURL ...
func (e *GoogleSearch) ToURL() string {
	baseURL, _ := url.Parse(searchUrl)

	tags := strings.Join(e.tags, " ")

	params := url.Values{}
	params.Add("q", tags)

	baseURL.RawQuery = params.Encode()

	return baseURL.String()
}

// Site ...
func (e *GoogleSearch) Site(site string) *GoogleSearch {
	e.tags = append(e.tags, concat(siteTag, site, false))

	return e
}

// Or ...
func (e *GoogleSearch) Or() *GoogleSearch {
	e.tags = append(e.tags, "OR")
	return e
}

// Intext ...
func (e *GoogleSearch) Intext(text string) *GoogleSearch {
	e.tags = append(e.tags, concat("", text, true))
	return e
}

// Inurl ...
func (e *GoogleSearch) Inurl(url string) *GoogleSearch {
	e.tags = append(e.tags, concat(urlTag, url, true))
	return e
}

// Filetype ...
func (e *GoogleSearch) Filetype(filetype string) *GoogleSearch {
	e.tags = append(e.tags, concat(filetypeTag, filetype, true))
	return e
}

// Cache ...
func (e *GoogleSearch) Cache(url string) *GoogleSearch {
	e.tags = append(e.tags, concat(cacheTag, url, true))
	return e
}

// Related ...
func (e *GoogleSearch) Related(url string) *GoogleSearch {
	e.tags = append(e.tags, concat(relatedTag, url, true))
	return e
}

// Ext ...
func (e *GoogleSearch) Ext(ext string) *GoogleSearch {
	e.tags = append(e.tags, concat(extTag, ext, false))
	return e
}

// Exclude ...
func (e *GoogleSearch) Exclude(value string) *GoogleSearch {
	e.tags = append(e.tags, concat(excludeTag, value, false))
	return e
}
