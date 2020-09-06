package duckduckgo

import (
	"net/url"
	"strings"
)

const (
	searchURL     = "https://www.google.com/search"
	siteTag       = "site:"
	urlTag        = "inurl:"
	filetypeTag   = "filetype:"
	extTag        = "ext:"
	excludeTag    = "-"
	intitleTag    = "intitle:"
	intextTag     = "intext:"
	operatorOr    = "|"
	operatorAnd   = "+"
	allInURLTag   = "allinurl:"
	locationTag   = "region:"
	feedTag       = "feed:"
	hasfeedTag    = "hasfeed:"
	languageTag   = "language:"
	allintitleTag = "allintitle:"
)

// DuckDuckGo is the Google search implementation for Dorkgen
type DuckDuckGo struct {
	tags []string
}

// New creates a new instance of DuckDuckGo
func New() *DuckDuckGo {
	return &DuckDuckGo{}
}

func (e *DuckDuckGo) join(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
}

// String converts all tags to a single request
func (e *DuckDuckGo) String() string {
	return strings.Join(e.tags, " ")
}

// QueryValues returns search request as URL values
func (e *DuckDuckGo) QueryValues() url.Values {
	tags := strings.Join(e.tags, " ")

	params := url.Values{}
	params.Add("q", tags)

	return params
}

// URL converts tags to an encoded Google Search URL
func (e *DuckDuckGo) URL() string {
	baseURL, _ := url.Parse(searchURL)

	baseURL.RawQuery = e.QueryValues().Encode()

	return baseURL.String()
}

// Site specifically searches that particular site and lists all the results for that site.
func (e *DuckDuckGo) Site(site string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(siteTag, site, false))
	return e
}

// Or puts an OR operator in the request
func (e *DuckDuckGo) Or() *DuckDuckGo {
	e.tags = append(e.tags, operatorOr)
	return e
}

// And puts an AND operator in the request
func (e *DuckDuckGo) And() *DuckDuckGo {
	e.tags = append(e.tags, operatorAnd)
	return e
}

// InText searches for the occurrences of keywords all at once or one at a time.
func (e *DuckDuckGo) InText(text string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(intextTag, text, true))
	return e
}

// InURL searches for a URL matching one of the keywords.
func (e *DuckDuckGo) InURL(url string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(urlTag, url, true))
	return e
}

// FileType searches for a particular filetype mentioned in the query.
func (e *DuckDuckGo) FileType(filetype string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(filetypeTag, filetype, true))
	return e
}

// Ext searches for a particular file extension mentioned in the query.
func (e *DuckDuckGo) Ext(ext string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(extTag, ext, false))
	return e
}

// Exclude excludes some results.
func (e *DuckDuckGo) Exclude(tags *DuckDuckGo) *DuckDuckGo {
	e.tags = append(e.tags, e.join(excludeTag, tags.String(), false))
	return e
}

// Group isolate tags between parentheses
func (e *DuckDuckGo) Group(tags *DuckDuckGo) *DuckDuckGo {
	e.tags = append(e.tags, "("+tags.String()+")")
	return e
}

// InTitle searches for occurrences of keywords in title all or one.
func (e *DuckDuckGo) InTitle(value string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(intitleTag, value, true))
	return e
}

// Plain allows you to add additional values as string without any kind of formatting.
func (e *DuckDuckGo) Plain(value string) *DuckDuckGo {
	e.tags = append(e.tags, value)
	return e
}

// AllInURL finds pages that include a specific keyword as part of their indexed URLs.
func (e *DuckDuckGo) AllInURL(value string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(allInURLTag, value, true))
	return e
}

// Location searches for specific region.
// An iso location code is a short code for a country for example, Egypt is eg and USA is us.
// https://en.wikipedia.org/wiki/ISO_3166-1
func (e *DuckDuckGo) Location(isoCode string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(locationTag, isoCode, true))
	return e
}

// Feed finds RSS feed related to search term (i.e. rss).
func (e *DuckDuckGo) Feed(feed string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(feedTag, feed, false))
	return e
}

// HasFeed finds webpages that contain both the term or terms for which you are querying and one or more RSS or Atom feeds.
func (e *DuckDuckGo) HasFeed(url string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(hasfeedTag, url, true))
	return e
}

// Language returns websites that match the search term in a specified language.
// See https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes for a complete list of ISO 639-1 codes you can use.
func (e *DuckDuckGo) Language(lang string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(languageTag, lang, false))
	return e
}

// AllInTitle finds pages that include a specific keyword as part of the indexed title tag.
func (e *DuckDuckGo) AllInTitle(value string) *DuckDuckGo {
	e.tags = append(e.tags, e.join(allintitleTag, value, true))
	return e
}
