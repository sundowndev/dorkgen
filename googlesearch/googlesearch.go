package googlesearch

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
	tags []string
}

// New creates a new instance of GoogleSearch
func New() *GoogleSearch {
	return &GoogleSearch{}
}

func (_ *GoogleSearch) concat(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
}

// String converts all tags to a single request
func (g *GoogleSearch) String() string {
	return strings.Join(g.tags, " ")
}

// QueryValues returns search request as URL values
func (g *GoogleSearch) QueryValues() url.Values {
	tags := strings.Join(g.tags, " ")

	params := url.Values{}
	params.Add("q", tags)

	return params
}

// URL converts tags to an encoded Google Search URL
func (g *GoogleSearch) URL() string {
	baseURL, _ := url.Parse(searchURL)

	baseURL.RawQuery = g.QueryValues().Encode()

	return baseURL.String()
}

// Site specifically searches that particular site and lists all the results for that site.
func (g *GoogleSearch) Site(site string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(siteTag, site, false))
	return g
}

// Or puts an OR operator in the request
func (g *GoogleSearch) Or() *GoogleSearch {
	g.tags = append(g.tags, operatorOr)
	return g
}

// And puts an AND operator in the request
func (g *GoogleSearch) And() *GoogleSearch {
	g.tags = append(g.tags, operatorAnd)
	return g
}

// Intext searches for the occurrences of keywords all at once or one at a time.
func (g *GoogleSearch) Intext(text string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(intextTag, text, true))
	return g
}

// Inurl searches for a URL matching one of the keywords.
func (g *GoogleSearch) Inurl(url string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(urlTag, url, true))
	return g
}

// Filetype searches for a particular filetype mentioned in the query.
func (g *GoogleSearch) Filetype(filetype string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(filetypeTag, filetype, true))
	return g
}

// Cache shows the version of the web page that Google has in its cache.
func (g *GoogleSearch) Cache(url string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(cacheTag, url, true))
	return g
}

// Related list web pages that are “similar” to a specified web page.
func (g *GoogleSearch) Related(url string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(relatedTag, url, true))
	return g
}

// Ext searches for a particular file extension mentioned in the query.
func (g *GoogleSearch) Ext(ext string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(extTag, ext, false))
	return g
}

// Exclude excludes some results.
func (g *GoogleSearch) Exclude(tags *GoogleSearch) *GoogleSearch {
	g.tags = append(g.tags, g.concat(excludeTag, tags.String(), false))
	return g
}

// Group isolate tags between parentheses
func (g *GoogleSearch) Group(tags *GoogleSearch) *GoogleSearch {
	g.tags = append(g.tags, "("+tags.String()+")")
	return g
}

// Intitle searches for occurrences of keywords in title all or one.
func (g *GoogleSearch) Intitle(value string) *GoogleSearch {
	g.tags = append(g.tags, g.concat(intitleTag, value, true))
	return g
}

// Plain allows you to add additional values as string without any kind of formatting.
func (g *GoogleSearch) Plain(value string) *GoogleSearch {
	g.tags = append(g.tags, value)
	return g
}
