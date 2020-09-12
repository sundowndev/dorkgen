package googlesearch

import (
	"net/url"
	"strings"
)

const (
	searchURL    = "https://www.google.com/search"
	siteTag      = "site:"
	urlTag       = "inurl:"
	filetypeTag  = "filetype:"
	cacheTag     = "cache:"
	relatedTag   = "related:"
	extTag       = "ext:"
	excludeTag   = "-"
	intitleTag   = "intitle:"
	intextTag    = "intext:"
	operatorOr   = "|"
	operatorAnd  = "+"
	bookTag      = "book:"
	ipTag        = "ip:"
	mapsTag      = "maps:"
	allintextTag = "allintext:"
	infoTag      = "info:"
	inanchorTag  = "inanchor:"
)

// GoogleSearch is the Google search implementation for Dorkgen
type GoogleSearch struct {
	tags []string
}

// New creates a new instance of GoogleSearch
func New() *GoogleSearch {
	return &GoogleSearch{}
}

func (e *GoogleSearch) join(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
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
	e.tags = append(e.tags, e.join(siteTag, site, false))
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

// InText searches for the occurrences of keywords all at once or one at a time.
func (e *GoogleSearch) InText(text string) *GoogleSearch {
	e.tags = append(e.tags, e.join(intextTag, text, true))
	return e
}

// InURL searches for a URL matching one of the keywords.
func (e *GoogleSearch) InURL(url string) *GoogleSearch {
	e.tags = append(e.tags, e.join(urlTag, url, true))
	return e
}

// FileType searches for a particular filetype mentioned in the query.
func (e *GoogleSearch) FileType(filetype string) *GoogleSearch {
	e.tags = append(e.tags, e.join(filetypeTag, filetype, true))
	return e
}

// Cache shows the version of the web page that Google has in its cache.
func (e *GoogleSearch) Cache(url string) *GoogleSearch {
	e.tags = append(e.tags, e.join(cacheTag, url, true))
	return e
}

// Related list web pages that are “similar” to a specified web page.
func (e *GoogleSearch) Related(url string) *GoogleSearch {
	e.tags = append(e.tags, e.join(relatedTag, url, true))
	return e
}

// Ext searches for a particular file extension mentioned in the query.
func (e *GoogleSearch) Ext(ext string) *GoogleSearch {
	e.tags = append(e.tags, e.join(extTag, ext, false))
	return e
}

// Exclude excludes some results.
func (e *GoogleSearch) Exclude(tags *GoogleSearch) *GoogleSearch {
	e.tags = append(e.tags, e.join(excludeTag, tags.String(), false))
	return e
}

// Group isolate tags between parentheses
func (e *GoogleSearch) Group(tags *GoogleSearch) *GoogleSearch {
	e.tags = append(e.tags, "("+tags.String()+")")
	return e
}

// InTitle searches for occurrences of keywords in title all or one.
func (e *GoogleSearch) InTitle(value string) *GoogleSearch {
	e.tags = append(e.tags, e.join(intitleTag, value, true))
	return e
}

// Plain allows you to add additional values as string without any kind of formatting.
func (e *GoogleSearch) Plain(value string) *GoogleSearch {
	e.tags = append(e.tags, value)
	return e
}

// Book searches for book titles related to keywords.
func (e *GoogleSearch) Book(keyword string) *GoogleSearch {
	e.tags = append(e.tags, e.join(bookTag, keyword, true))
	return e
}

// Maps searches for maps related to keywords.
func (e *GoogleSearch) Maps(location string) *GoogleSearch {
	e.tags = append(e.tags, e.join(mapsTag, location, false))
	return e
}

// AllInText searches text of page.
func (e *GoogleSearch) AllInText(text string) *GoogleSearch {
	e.tags = append(e.tags, e.join(allintextTag, text, true))
	return e
}

// Info presents some information that Google has about a web page, including similar pages, the cached version of the page, and sites linking to the page.
func (e *GoogleSearch) Info(url string) *GoogleSearch {
	e.tags = append(e.tags, e.join(infoTag, url, true))
	return e
}

// InAnchor search link anchor text.
func (e *GoogleSearch) InAnchor(text string) *GoogleSearch {
	e.tags = append(e.tags, e.join(inanchorTag, text, true))
	return e
}
