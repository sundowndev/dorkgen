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

func (d *DuckDuckGo) concat(tag string, value string, quotes bool) string {
	if quotes {
		return tag + "\"" + value + "\""
	}

	return tag + value
}

// String converts all tags to a single request
func (d *DuckDuckGo) String() string {
	return strings.Join(d.tags, " ")
}

// QueryValues returns search request as URL values
func (d *DuckDuckGo) QueryValues() url.Values {
	tags := strings.Join(d.tags, " ")

	params := url.Values{}
	params.Add("q", tags)

	return params
}

// URL converts tags to an encoded Google Search URL
func (d *DuckDuckGo) URL() string {
	baseURL, _ := url.Parse(searchURL)

	baseURL.RawQuery = d.QueryValues().Encode()

	return baseURL.String()
}

// Site specifically searches that particular site and lists all the results for that site.
func (d *DuckDuckGo) Site(site string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(siteTag, site, false))
	return d
}

// Or puts an OR operator in the request
func (d *DuckDuckGo) Or() *DuckDuckGo {
	d.tags = append(d.tags, operatorOr)
	return d
}

// And puts an AND operator in the request
func (d *DuckDuckGo) And() *DuckDuckGo {
	d.tags = append(d.tags, operatorAnd)
	return d
}

// Intext searches for the occurrences of keywords all at once or one at a time.
func (d *DuckDuckGo) Intext(text string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(intextTag, text, true))
	return d
}

// Inurl searches for a URL matching one of the keywords.
func (d *DuckDuckGo) Inurl(url string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(urlTag, url, true))
	return d
}

// Filetype searches for a particular filetype mentioned in the query.
func (d *DuckDuckGo) Filetype(filetype string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(filetypeTag, filetype, true))
	return d
}

// Ext searches for a particular file extension mentioned in the query.
func (d *DuckDuckGo) Ext(ext string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(extTag, ext, false))
	return d
}

// Exclude excludes some results.
func (d *DuckDuckGo) Exclude(tags *DuckDuckGo) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(excludeTag, tags.String(), false))
	return d
}

// Group isolate tags between parentheses
func (d *DuckDuckGo) Group(tags *DuckDuckGo) *DuckDuckGo {
	d.tags = append(d.tags, "("+tags.String()+")")
	return d
}

// Intitle searches for occurrences of keywords in title all or one.
func (d *DuckDuckGo) Intitle(value string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(intitleTag, value, true))
	return d
}

// Plain allows you to add additional values as string without any kind of formatting.
func (d *DuckDuckGo) Plain(value string) *DuckDuckGo {
	d.tags = append(d.tags, value)
	return d
}

// AllInURL finds pages that include a specific keyword as part of their indexed URLs.
func (d *DuckDuckGo) AllInURL(value string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(allInURLTag, value, true))
	return d
}

// Location searches for specific region.
// An iso location code is a short code for a country for example, Egypt is eg and USA is us.
// https://en.wikipedia.org/wiki/ISO_3166-1
func (d *DuckDuckGo) Location(isoCode string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(locationTag, isoCode, true))
	return d
}

// Feed finds RSS feed related to search term (i.e. rss).
func (d *DuckDuckGo) Feed(feed string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(feedTag, feed, false))
	return d
}

// HasFeed finds webpages that contain both the term or terms for which you are querying and one or more RSS or Atom feeds.
func (d *DuckDuckGo) HasFeed(url string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(hasfeedTag, url, true))
	return d
}

// Language returns websites that match the search term in a specified language.
// See https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes for a complete list of ISO 639-1 codes you can use.
func (d *DuckDuckGo) Language(lang string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(languageTag, lang, false))
	return d
}

// AllInTitle finds pages that include a specific keyword as part of the indexed title tag.
func (d *DuckDuckGo) AllInTitle(value string) *DuckDuckGo {
	d.tags = append(d.tags, d.concat(allintitleTag, value, true))
	return d
}
