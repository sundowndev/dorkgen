package dorkgen

import (
	"net/url"
	"strings"
)

const (
	searchURL   = "https://www.google.com/search"
	siteTag     = "site:"
	urlTag      = "inurl:"
	filetypeTag = "filetype:"
	cacheTag    = "cache:"
	relatedTag  = "related:"
	extTag      = "ext:"
	excludeTag  = "-"
	intitleTag  = "intitle:"
	intextTag   = "intext:"
	operatorOr  = "|"
	operatorAnd = "+"
)

// GoogleSearch is the Google search implementation for Dorkgen
type GoogleSearch struct {
	EngineFactory
}

func NewGoogleSearch() *GoogleSearch {
	return &GoogleSearch{}
}

// String converts all tags to a single request
func (e *GoogleSearch) String() string {
	return strings.Join(e.tags, " ")
}

// QueryValues returns search request as URL values
func (e *GoogleSearch) QueryValues() url.Values {
	tags := strings.Join(e.tags, " ")

	params := url.Values{}
	params.Add("q", tags)

	return params
}

// URL converts tags to an encoded Google Search URL
func (e *GoogleSearch) URL() string {
	baseURL, _ := url.Parse(searchURL)

	baseURL.RawQuery = e.QueryValues().Encode()

	return baseURL.String()
}

// Site specifically searches that particular site and lists all the results for that site.
func (e *GoogleSearch) Site(site string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(siteTag, site, false))
	return e
}

// Or puts an OR operator in the request
func (e *GoogleSearch) Or() *GoogleSearch {
	e.tags = append(e.tags, operatorOr)
	return e
}

// And puts an AND operator in the request
func (e *GoogleSearch) And() *GoogleSearch {
	e.tags = append(e.tags, operatorAnd)
	return e
}

// Intext searches for the occurrences of keywords all at once or one at a time.
func (e *GoogleSearch) Intext(text string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(intextTag, text, true))
	return e
}

// Inurl searches for a URL matching one of the keywords.
func (e *GoogleSearch) Inurl(url string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(urlTag, url, true))
	return e
}

// Filetype searches for a particular filetype mentioned in the query.
func (e *GoogleSearch) Filetype(filetype string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(filetypeTag, filetype, true))
	return e
}

// Cache shows the version of the web page that Google has in its cache.
func (e *GoogleSearch) Cache(url string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(cacheTag, url, true))
	return e
}

// Related list web pages that are “similar” to a specified web page.
func (e *GoogleSearch) Related(url string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(relatedTag, url, true))
	return e
}

// Ext searches for a particular file extension mentioned in the query.
func (e *GoogleSearch) Ext(ext string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(extTag, ext, false))
	return e
}

// Exclude excludes some results.
func (e *GoogleSearch) Exclude(tags *GoogleSearch) *GoogleSearch {
	e.tags = append(e.tags, e.concat(excludeTag, tags.String(), false))
	return e
}

// Group isolate tags between parentheses
func (e *GoogleSearch) Group(tags *GoogleSearch) *GoogleSearch {
	e.tags = append(e.tags, "("+tags.String()+")")
	return e
}

// Intitle searches for occurrences of keywords in title all or one.
func (e *GoogleSearch) Intitle(value string) *GoogleSearch {
	e.tags = append(e.tags, e.concat(intitleTag, value, true))
	return e
}

// Plain allows you to add additional values as string without any kind of formatting.
func (e *GoogleSearch) Plain(value string) *GoogleSearch {
	e.tags = append(e.tags, value)
	return e
}
